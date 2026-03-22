package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/open-rta/genesis/tools/validator/internal/manifest"
	"github.com/open-rta/genesis/tools/validator/internal/report"
	"github.com/open-rta/genesis/tools/validator/internal/validator"
)

func main() {
	if len(os.Args) < 2 {
		printRootUsage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "validate":
		os.Exit(runValidate(os.Args[2:]))
	case "init-manifest":
		os.Exit(runInitManifest(os.Args[2:]))
	default:
		// Backward-compatible single-command validate mode.
		os.Exit(runValidate(os.Args[1:]))
	}
}

func runValidate(args []string) int {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	reportPath := fs.String("report", "", "Path to write machine-readable validation report JSON")
	manifestFlag := fs.String("manifest", "", "Path to manifest")
	if err := fs.Parse(args); err != nil {
		return 2
	}

	manifestPath := strings.TrimSpace(*manifestFlag)
	if manifestPath == "" && fs.NArg() > 0 {
		manifestPath = fs.Arg(0)
	}
	if manifestPath == "" {
		fmt.Fprintln(os.Stderr, "Usage: validate <manifest-path> [--report <output.json>] OR validate --manifest <manifest-path>")
		return 2
	}

	r, err := validator.Run(manifestPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		return 1
	}

	if *reportPath != "" {
		if err := report.WriteJSON(*reportPath, r); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return 1
		}
		fmt.Printf("Report written to %s\n", *reportPath)
	}

	if r.Passed {
		return 0
	}
	return 1
}

func runInitManifest(args []string) int {
	fs := flag.NewFlagSet("init-manifest", flag.ContinueOnError)
	output := fs.String("output", "./open-rta-manifest.json", "Output path")
	force := fs.Bool("force", false, "Overwrite existing output file")
	nonInteractive := fs.Bool("non-interactive", false, "Disable prompts; rely on flags")

	runtimeName := fs.String("runtime-name", "", "Runtime name")
	runtimeVersion := fs.String("runtime-version", "0.1.0", "Runtime version")
	openRTAVersion := fs.String("open-rta-version", "0.1.0", "Open RTA version")
	level := fs.String("level", "L0", "Claimed level: L0/L1/L2/L3/L4")
	description := fs.String("description", "", "Short runtime description")
	repository := fs.String("repository", "", "Repository URL")
	homepage := fs.String("homepage", "", "Homepage URL")
	declaredBy := fs.String("declared-by", "", "Declared by")
	contact := fs.String("contact", "", "Contact info")
	laws := fs.String("laws", "", "Comma-separated laws claimed")
	appliesTo := fs.String("applies-to", "", "Comma-separated applies-to scope")
	exclusions := fs.String("exclusions", "", "Comma-separated exclusions")
	objectiveRef := fs.String("objective-ref", "", "Objective evidence path or URL")
	authorityRef := fs.String("authority-ref", "", "Authority evidence path or URL")
	traceRef := fs.String("trace-ref", "", "Trace evidence path or URL")
	controlRef := fs.String("control-ref", "", "Control evidence path or URL")
	oversightRef := fs.String("oversight-ref", "", "Oversight evidence path or URL")
	replayRef := fs.String("replay-ref", "", "Replay evidence path or URL")
	conformanceRef := fs.String("conformance-ref", "", "Conformance report path or URL")
	limitations := fs.String("limitations", "", "Comma-separated limitations")
	autoTargets := fs.String("auto-targets", "", "Comma-separated automatic validation targets")
	manualTargets := fs.String("manual-targets", "", "Comma-separated manual review targets")

	if err := fs.Parse(args); err != nil {
		return 2
	}

	reader := bufio.NewReader(os.Stdin)
	opts := manifest.InitOptions{
		RuntimeName:    *runtimeName,
		RuntimeVersion: *runtimeVersion,
		OpenRTAVersion: *openRTAVersion,
		Level:          manifest.Level(strings.ToUpper(strings.TrimSpace(*level))),
		Description:    *description,
		Repository:     *repository,
		Homepage:       *homepage,
		DeclaredBy:     *declaredBy,
		Contact:        *contact,
		Laws:           manifest.ParseCSV(*laws),
		AppliesTo:      manifest.ParseCSV(*appliesTo),
		Exclusions:     manifest.ParseCSV(*exclusions),
		ObjectiveRef:   *objectiveRef,
		AuthorityRef:   *authorityRef,
		TraceRef:       *traceRef,
		ControlRef:     *controlRef,
		OversightRef:   *oversightRef,
		ReplayRef:      *replayRef,
		ConformanceRef: *conformanceRef,
		Limitations:    manifest.ParseCSV(*limitations),
		AutoTargets:    manifest.ParseCSV(*autoTargets),
		ManualTargets:  manifest.ParseCSV(*manualTargets),
		OutputPath:     *output,
		Force:          *force,
		NonInteractive: *nonInteractive,
	}

	if !opts.NonInteractive {
		promptInitManifest(reader, &opts)
	}

	m, err := manifest.GenerateFromOptions(opts)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not generate manifest: %v\n", err)
		return 1
	}

	showSummary(m, opts.OutputPath)
	if !opts.NonInteractive && !manifest.PromptYesNo(reader, "Generate manifest with these values? [Y/n]: ", true) {
		fmt.Println("Aborted.")
		return 1
	}

	writtenPath, err := manifest.WriteManifest(opts.OutputPath, m, opts.Force, !opts.NonInteractive, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not write manifest: %v\n", err)
		return 1
	}

	fmt.Printf("Manifest written to %s\n", writtenPath)
	fmt.Printf("Next step: validate this manifest with the Open RTA validator.\n")
	fmt.Printf("Example: go run . validate --manifest %s\n", writtenPath)
	return 0
}

func promptInitManifest(reader *bufio.Reader, opts *manifest.InitOptions) {
	fmt.Println("SECTION A — BASIC RUNTIME IDENTITY")
	opts.RuntimeName = manifest.PromptValue(reader, fmt.Sprintf("Runtime name%s: ", requiredMark(opts.RuntimeName)), opts.RuntimeName)
	opts.RuntimeVersion = manifest.PromptValue(reader, fmt.Sprintf("Runtime version [%s]: ", defaultString(opts.RuntimeVersion, "0.1.0")), defaultString(opts.RuntimeVersion, "0.1.0"))
	opts.OpenRTAVersion = manifest.PromptValue(reader, fmt.Sprintf("Open RTA version [%s]: ", defaultString(opts.OpenRTAVersion, "0.1.0")), defaultString(opts.OpenRTAVersion, "0.1.0"))
	opts.Description = manifest.PromptValue(reader, "Short description (optional): ", opts.Description)
	opts.Repository = manifest.PromptValue(reader, "Repository URL (optional): ", opts.Repository)
	opts.Homepage = manifest.PromptValue(reader, "Homepage URL (optional): ", opts.Homepage)

	fmt.Println("\nSECTION B — COMPLIANCE CLAIM")
	level := manifest.PromptValue(reader, fmt.Sprintf("Claimed compliance level [%s]: ", defaultString(string(opts.Level), "L0")), defaultString(string(opts.Level), "L0"))
	opts.Level = manifest.Level(strings.ToUpper(strings.TrimSpace(level)))

	laws := []string{}
	if manifest.PromptYesNo(reader, "Claim no-rogue-autonomy? [Y/n]: ", true) {
		laws = append(laws, "no-rogue-autonomy")
	}
	if manifest.PromptYesNo(reader, "Claim observable-execution? [Y/n]: ", true) {
		laws = append(laws, "observable-execution")
	}
	if manifest.PromptYesNo(reader, "Claim interruptible-control? [y/N]: ", false) {
		laws = append(laws, "interruptible-control")
	}
	if manifest.PromptYesNo(reader, "Claim replayable-behavior? [y/N]: ", false) {
		laws = append(laws, "replayable-behavior")
	}
	opts.Laws = laws
	if opts.Level == manifest.LevelL4 {
		fmt.Println("Note: L4 still requires manual/foundation review.")
	}
	opts.AppliesTo = manifest.ParseCSV(manifest.PromptValue(reader, "Applies to (comma-separated, optional): ", strings.Join(opts.AppliesTo, ",")))
	opts.Exclusions = manifest.ParseCSV(manifest.PromptValue(reader, "Exclusions or unsupported areas (comma-separated, optional): ", strings.Join(opts.Exclusions, ",")))

	fmt.Println("\nSECTION C — EVIDENCE REFERENCES")
	fmt.Println("Provide references to evidence artifacts if available now. You can leave fields blank and fill them later.")
	opts.ObjectiveRef = manifest.PromptValue(reader, "Objective evidence path or URL (optional): ", opts.ObjectiveRef)
	opts.AuthorityRef = manifest.PromptValue(reader, "Authority evidence path or URL (optional): ", opts.AuthorityRef)
	opts.TraceRef = manifest.PromptValue(reader, "Trace evidence path or URL (optional): ", opts.TraceRef)
	opts.ControlRef = manifest.PromptValue(reader, "Control evidence path or URL (optional): ", opts.ControlRef)
	opts.OversightRef = manifest.PromptValue(reader, "Oversight evidence path or URL (optional): ", opts.OversightRef)
	opts.ReplayRef = manifest.PromptValue(reader, "Replay evidence path or URL (optional): ", opts.ReplayRef)
	opts.ConformanceRef = manifest.PromptValue(reader, "Conformance report path or URL (optional): ", opts.ConformanceRef)

	fmt.Println("\nSECTION D — TEST TARGETS")
	useDefaultAuto := manifest.PromptYesNo(reader, "Use default automatic validation targets based on provided evidence? [Y/n]: ", true)
	if !useDefaultAuto {
		opts.AutoTargets = manifest.ParseCSV(manifest.PromptValue(reader, "Automatic validation targets (comma-separated): ", strings.Join(opts.AutoTargets, ",")))
	}
	opts.ManualTargets = manifest.ParseCSV(manifest.PromptValue(reader, "Add manual review targets (comma-separated, optional): ", strings.Join(opts.ManualTargets, ",")))

	fmt.Println("\nSECTION E — ATTESTATION")
	opts.DeclaredBy = manifest.PromptValue(reader, fmt.Sprintf("Declared by%s: ", requiredMark(opts.DeclaredBy)), opts.DeclaredBy)
	opts.Contact = manifest.PromptValue(reader, "Contact info (optional): ", opts.Contact)

	fmt.Println("\nSECTION F — LIMITATIONS")
	opts.Limitations = manifest.ParseCSV(manifest.PromptValue(reader, "Limitations or known gaps (comma-separated, optional): ", strings.Join(opts.Limitations, ",")))

	fmt.Println("\nSECTION G — OUTPUT")
	opts.OutputPath = manifest.PromptValue(reader, fmt.Sprintf("Output path [%s]: ", defaultString(opts.OutputPath, "./open-rta-manifest.json")), defaultString(opts.OutputPath, "./open-rta-manifest.json"))
}

func showSummary(m manifest.Manifest, outputPath string) {
	fmt.Println("\nSECTION H — FINAL SUMMARY")
	fmt.Printf("Runtime: %s @ %s\n", m.Runtime.Name, m.Runtime.Version)
	fmt.Printf("Claimed level: %s\n", m.Compliance.LevelClaimed)
	fmt.Printf("Laws claimed: %s\n", strings.Join(m.Compliance.LawsClaimed, ", "))
	keys := make([]string, 0, len(m.Evidence))
	for k := range m.Evidence {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("Evidence keys provided: %s\n", strings.Join(keys, ", "))
	fmt.Printf("Output path: %s\n", outputPath)
}

func requiredMark(v string) string {
	if strings.TrimSpace(v) == "" {
		return ""
	}
	return " [currently set]"
}

func defaultString(v, fallback string) string {
	if strings.TrimSpace(v) == "" {
		return fallback
	}
	return v
}

func printRootUsage() {
	fmt.Println("Open RTA Validator")
	fmt.Println("Usage:")
	fmt.Println("  validate <manifest-path> [--report <output.json>]")
	fmt.Println("  init-manifest [flags]")
}
