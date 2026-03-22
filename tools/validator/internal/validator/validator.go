package validator

func Validate(manifestPath string) (ValidationReport, error) {
	m, ctx, err := LoadManifest(manifestPath)
	if err != nil {
		return ValidationReport{}, err
	}

	report := ValidationReport{
		ManifestPath:    ctx.ManifestPath,
		RuntimeName:     m.Runtime.Name,
		RuntimeVersion:  m.Runtime.Version,
		ClaimedLevel:    m.Compliance.LevelClaimed,
		Errors:          []string{},
		Warnings:        []string{},
		EvidenceResults: map[string]EvidenceResult{},
		Notes:           []string{},
	}

	ok, schemaErrs := validateManifestSchema(m, ctx.RepoRoot)
	if !ok {
		for _, e := range schemaErrs {
			report.Errors = append(report.Errors, "manifest schema: "+e)
		}
	}

	evidenceResults, evidenceErrs, evidenceWarns := resolveEvidence(m, ctx)
	report.EvidenceResults = evidenceResults
	report.Errors = append(report.Errors, evidenceErrs...)
	report.Warnings = append(report.Warnings, evidenceWarns...)

	report.Errors = append(report.Errors, validateClaims(m)...)

	validatedLevel, levelErrs, notes := validateLevel(m)
	report.ValidatedLevel = validatedLevel
	report.Errors = append(report.Errors, levelErrs...)
	report.Notes = append(report.Notes, notes...)

	report.Passed = len(report.Errors) == 0
	for _, res := range report.EvidenceResults {
		if res.SchemaValid != nil && *res.SchemaValid {
			report.Summary.SchemaValidatedCount++
		}
	}
	report.Summary.CheckedEvidenceCount = len(report.EvidenceResults)
	report.Summary.HardErrorCount = len(report.Errors)
	report.Summary.WarningCount = len(report.Warnings)

	if report.ValidatedLevel == "" {
		report.ValidatedLevel = LevelL0
	}

	printSummary(report)
	return report, nil
}
