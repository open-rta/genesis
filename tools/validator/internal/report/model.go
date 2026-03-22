package report

import "github.com/open-rta/genesis/tools/validator/internal/manifest"

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
	ClaimedLevel    manifest.Level            `json:"claimed_level"`
	ValidatedLevel  manifest.Level            `json:"validated_level"`
	Passed          bool                      `json:"passed"`
	Errors          []string                  `json:"errors"`
	Warnings        []string                  `json:"warnings"`
	EvidenceResults map[string]EvidenceResult `json:"evidence_results"`
	Notes           []string                  `json:"notes"`
	Summary         Summary                   `json:"summary"`
}
