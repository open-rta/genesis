package manifest

type Level string

const (
	LevelL0 Level = "L0"
	LevelL1 Level = "L1"
	LevelL2 Level = "L2"
	LevelL3 Level = "L3"
	LevelL4 Level = "L4"
)

type EvidenceReference struct {
	Type        string `json:"type"`
	Ref         string `json:"ref"`
	Schema      string `json:"schema,omitempty"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description,omitempty"`
	Required    *bool  `json:"required,omitempty"`
}

type TestTarget struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Manifest struct {
	ManifestVersion string `json:"manifest_version"`
	OpenRTAVersion  string `json:"open_rta_version"`
	Runtime         struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description,omitempty"`
		Repository  string `json:"repository,omitempty"`
		Homepage    string `json:"homepage,omitempty"`
		Vendor      string `json:"vendor,omitempty"`
	} `json:"runtime"`
	Compliance struct {
		LawsClaimed  []string `json:"laws_claimed"`
		LevelClaimed Level    `json:"level_claimed"`
		AppliesTo    []string `json:"applies_to,omitempty"`
		Exclusions   []string `json:"exclusions,omitempty"`
	} `json:"compliance"`
	Evidence    map[string]EvidenceReference `json:"evidence"`
	TestTargets []TestTarget                 `json:"test_targets,omitempty"`
	Attestation struct {
		By        string `json:"by"`
		Contact   string `json:"contact,omitempty"`
		Timestamp string `json:"timestamp"`
		Statement string `json:"statement"`
		Signature string `json:"signature,omitempty"`
	} `json:"attestation"`
	Limitations []string `json:"limitations,omitempty"`
}

type Context struct {
	RepoRoot     string
	ManifestPath string
	ManifestDir  string
}
