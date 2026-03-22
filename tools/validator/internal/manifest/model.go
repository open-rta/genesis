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
	Description string `json:"description,omitempty"`
	Required    *bool  `json:"required,omitempty"`
}

type Manifest struct {
	ManifestVersion string `json:"manifest_version"`
	OpenRTAVersion  string `json:"open_rta_version"`
	Runtime         struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Vendor  string `json:"vendor,omitempty"`
	} `json:"runtime"`
	Compliance struct {
		LawsClaimed  []string `json:"laws_claimed"`
		LevelClaimed Level    `json:"level_claimed"`
	} `json:"compliance"`
	Evidence    map[string]EvidenceReference `json:"evidence"`
	Attestation struct {
		By        string `json:"by"`
		Timestamp string `json:"timestamp"`
		Statement string `json:"statement"`
	} `json:"attestation"`
}

type Context struct {
	RepoRoot     string
	ManifestPath string
	ManifestDir  string
}
