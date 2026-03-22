package validator

import (
	"github.com/open-rta/genesis/tools/validator/internal/evidence"
	"github.com/open-rta/genesis/tools/validator/internal/levels"
	"github.com/open-rta/genesis/tools/validator/internal/manifest"
	"github.com/open-rta/genesis/tools/validator/internal/report"
)

func Run(manifestPath string) (report.ValidationReport, error) {
	m, ctx, err := manifest.Load(manifestPath)
	if err != nil {
		return report.ValidationReport{}, err
	}

	r := report.ValidationReport{
		ManifestPath:    ctx.ManifestPath,
		RuntimeName:     m.Runtime.Name,
		RuntimeVersion:  m.Runtime.Version,
		ClaimedLevel:    m.Compliance.LevelClaimed,
		Errors:          []string{},
		Warnings:        []string{},
		EvidenceResults: map[string]report.EvidenceResult{},
		Notes:           []string{},
	}

	ok, schemaErrs := manifest.ValidateSchema(m, ctx.RepoRoot)
	if !ok {
		for _, e := range schemaErrs {
			r.Errors = append(r.Errors, "manifest schema: "+e)
		}
	}

	evidenceResults, evidenceErrs, evidenceWarns := evidence.ResolveAndValidate(m, ctx)
	r.EvidenceResults = evidenceResults
	r.Errors = append(r.Errors, evidenceErrs...)
	r.Warnings = append(r.Warnings, evidenceWarns...)
	r.Errors = append(r.Errors, validateClaims(m)...)

	validatedLevel, levelErrs, notes := levels.ValidateClaimedLevel(m)
	r.ValidatedLevel = validatedLevel
	r.Errors = append(r.Errors, levelErrs...)
	r.Notes = append(r.Notes, notes...)

	r.Passed = len(r.Errors) == 0
	for _, res := range r.EvidenceResults {
		if res.SchemaValid != nil && *res.SchemaValid {
			r.Summary.SchemaValidatedCount++
		}
	}
	r.Summary.CheckedEvidenceCount = len(r.EvidenceResults)
	r.Summary.HardErrorCount = len(r.Errors)
	r.Summary.WarningCount = len(r.Warnings)
	if r.ValidatedLevel == "" {
		r.ValidatedLevel = manifest.LevelL0
	}

	report.PrintSummary(r)
	return r, nil
}
