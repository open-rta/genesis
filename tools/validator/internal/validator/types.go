package validator

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

type EvidenceResult struct {
	Found       bool     `json:"found"`
	Path        string   `json:"path,omitempty"`
	SchemaValid *bool    `json:"schema_valid,omitempty"`
	Errors      []string `json:"errors,omitempty"`
	Warnings    []string `json:"warnings,omitempty"`
}

type Summary struct {
	CheckedEvidenceCount int `json:"checked_evidence_count"`
	SchemaValidatedCount int `json:"schema_validated_count"`
	HardErrorCount       int `json:"hard_error_count"`
	WarningCount         int `json:"warning_count"`
}

type ValidationReport struct {
	ManifestPath    string                    `json:"manifest_path"`
	RuntimeName     string                    `json:"runtime_name"`
	RuntimeVersion  string                    `json:"runtime_version"`
	ClaimedLevel    Level                     `json:"claimed_level"`
	ValidatedLevel  Level                     `json:"validated_level"`
	Passed          bool                      `json:"passed"`
	Errors          []string                  `json:"errors"`
	Warnings        []string                  `json:"warnings"`
	EvidenceResults map[string]EvidenceResult `json:"evidence_results"`
	Notes           []string                  `json:"notes"`
	Summary         Summary                   `json:"summary"`
}
