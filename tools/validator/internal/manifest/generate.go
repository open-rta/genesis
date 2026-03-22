package manifest

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type InitOptions struct {
	AutoTargets    []string
	ManualTargets  []string
	RuntimeName    string
	RuntimeVersion string
	OpenRTAVersion string
	Level          Level
	Description    string
	Repository     string
	Homepage       string
	DeclaredBy     string
	Contact        string
	Laws           []string
	AppliesTo      []string
	Exclusions     []string
	ObjectiveRef   string
	AuthorityRef   string
	TraceRef       string
	ControlRef     string
	OversightRef   string
	ReplayRef      string
	ConformanceRef string
	Limitations    []string
	OutputPath     string
	Force          bool
	NonInteractive bool
}

func GenerateFromOptions(opts InitOptions) (Manifest, error) {
	if strings.TrimSpace(opts.RuntimeName) == "" {
		return Manifest{}, fmt.Errorf("runtime name is required")
	}
	if strings.TrimSpace(opts.DeclaredBy) == "" {
		return Manifest{}, fmt.Errorf("declared-by is required")
	}

	m := Manifest{
		ManifestVersion: "1.0.0",
		OpenRTAVersion:  defaultString(opts.OpenRTAVersion, "0.1.0"),
		Evidence:        map[string]EvidenceReference{},
		Limitations:     filterEmpty(opts.Limitations),
	}
	m.Runtime.Name = strings.TrimSpace(opts.RuntimeName)
	m.Runtime.Version = defaultString(opts.RuntimeVersion, "0.1.0")
	m.Runtime.Description = strings.TrimSpace(opts.Description)
	m.Runtime.Repository = strings.TrimSpace(opts.Repository)
	m.Runtime.Homepage = strings.TrimSpace(opts.Homepage)
	m.Compliance.LevelClaimed = defaultLevel(opts.Level)
	m.Compliance.LawsClaimed = inferLaws(m.Compliance.LevelClaimed, filterEmpty(opts.Laws))
	m.Compliance.AppliesTo = filterEmpty(opts.AppliesTo)
	m.Compliance.Exclusions = filterEmpty(opts.Exclusions)
	m.Attestation.By = strings.TrimSpace(opts.DeclaredBy)
	m.Attestation.Contact = strings.TrimSpace(opts.Contact)
	m.Attestation.Timestamp = time.Now().UTC().Format(time.RFC3339)
	m.Attestation.Statement = "To the best of our knowledge, this manifest accurately references current conformance evidence."

	addEvidence(&m, "objective", opts.ObjectiveRef)
	addEvidence(&m, "authority", opts.AuthorityRef)
	addEvidence(&m, "trace", opts.TraceRef)
	addEvidence(&m, "control", opts.ControlRef)
	addEvidence(&m, "oversight", opts.OversightRef)
	addEvidence(&m, "replay", opts.ReplayRef)
	addEvidence(&m, "conformance_report", opts.ConformanceRef)

	if len(m.Evidence) == 0 {
		m.Evidence = map[string]EvidenceReference{}
	}

	buildTestTargets(&m, opts)

	return m, nil
}

func addEvidence(m *Manifest, key, ref string) {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return
	}
	eType := "file"
	if strings.HasPrefix(strings.ToLower(ref), "http://") || strings.HasPrefix(strings.ToLower(ref), "https://") {
		eType = "url"
	}
	m.Evidence[key] = EvidenceReference{Type: eType, Ref: ref, Format: inferFormat(ref), Schema: key}
}

func inferFormat(ref string) string {
	r := strings.ToLower(ref)
	switch {
	case strings.HasSuffix(r, ".json"):
		return "json"
	case strings.HasSuffix(r, ".yaml"), strings.HasSuffix(r, ".yml"):
		return "yaml"
	case strings.HasSuffix(r, ".md"):
		return "markdown"
	default:
		return "text"
	}
}

func defaultString(v, fallback string) string {
	if strings.TrimSpace(v) == "" {
		return fallback
	}
	return strings.TrimSpace(v)
}

func defaultLevel(l Level) Level {
	if l == "" {
		return LevelL0
	}
	return l
}

func inferLaws(level Level, laws []string) []string {
	if len(laws) > 0 {
		return laws
	}
	if level == LevelL1 {
		return []string{"no-rogue-autonomy", "observable-execution"}
	}
	if level == LevelL2 {
		return []string{"no-rogue-autonomy", "observable-execution", "interruptible-control"}
	}
	if level == LevelL3 || level == LevelL4 {
		return []string{"no-rogue-autonomy", "observable-execution", "interruptible-control", "replayable-behavior"}
	}
	return []string{"no-rogue-autonomy"}
}

func filterEmpty(items []string) []string {
	out := []string{}
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func buildTestTargets(m *Manifest, opts InitOptions) {
	autoTargets := filterEmpty(opts.AutoTargets)
	if len(autoTargets) == 0 {
		for key := range m.Evidence {
			autoTargets = append(autoTargets, key)
		}
	}
	for _, t := range autoTargets {
		m.TestTargets = append(m.TestTargets, TestTarget{ID: "auto-" + t, Description: "automatic validation target: " + t})
	}
	for _, t := range filterEmpty(opts.ManualTargets) {
		m.TestTargets = append(m.TestTargets, TestTarget{ID: "manual-" + t, Description: "manual review target: " + t})
	}
}

func ParseCSV(input string) []string {
	if strings.TrimSpace(input) == "" {
		return nil
	}
	parts := strings.Split(input, ",")
	return filterEmpty(parts)
}

func PromptYesNo(reader *bufio.Reader, prompt string, defaultYes bool) bool {
	fmt.Print(prompt)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(strings.ToLower(val))
	if val == "" {
		return defaultYes
	}
	return val == "y" || val == "yes"
}

func PromptValue(reader *bufio.Reader, prompt, fallback string) string {
	fmt.Print(prompt)
	val, _ := reader.ReadString('\n')
	val = strings.TrimSpace(val)
	if val == "" {
		return fallback
	}
	return val
}

func WriteManifest(outputPath string, manifest Manifest, force bool, interactive bool, reader *bufio.Reader) (string, error) {
	if outputPath == "" {
		outputPath = "./open-rta-manifest.json"
	}
	absPath, err := filepath.Abs(outputPath)
	if err != nil {
		return "", err
	}
	if _, err := os.Stat(absPath); err == nil && !force {
		if !interactive {
			return "", fmt.Errorf("output file already exists: %s (use --force to overwrite)", absPath)
		}
		if !PromptYesNo(reader, "File already exists. Overwrite? [y/N]: ", false) {
			return "", fmt.Errorf("aborted: output file exists")
		}
	}

	bytes, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(absPath, append(bytes, '\n'), 0o644); err != nil {
		return "", err
	}
	return absPath, nil
}
