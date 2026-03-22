package evidence

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/open-rta/genesis/tools/validator/internal/manifest"
	"github.com/open-rta/genesis/tools/validator/internal/report"
	"github.com/open-rta/genesis/tools/validator/internal/schema"
)

func ResolveAndValidate(m manifest.Manifest, ctx manifest.Context) (map[string]report.EvidenceResult, []string, []string) {
	results := map[string]report.EvidenceResult{}
	errs := []string{}
	warns := []string{}

	for key, ref := range m.Evidence {
		r := report.EvidenceResult{Found: false}
		if ref.Type == "url" || isURL(ref.Ref) {
			r.Warnings = append(r.Warnings, "URL evidence is not fetched automatically in v0.2")
			warns = append(warns, fmt.Sprintf("evidence '%s' uses URL and was not auto-validated", key))
			results[key] = r
			continue
		}

		resolved := resolveLocalPath(ctx.ManifestDir, ref.Ref)
		r.Path = resolved
		if !fileExists(resolved) {
			r.Errors = append(r.Errors, "referenced evidence file does not exist")
			errs = append(errs, fmt.Sprintf("evidence '%s' not found at %s", key, resolved))
			results[key] = r
			continue
		}
		r.Found = true

		schemaPath, known := schema.ArtifactSchemaPath(ctx.RepoRoot, key)
		if !known {
			r.Warnings = append(r.Warnings, "no known schema mapping for this evidence key")
			results[key] = r
			continue
		}

		bytes, err := os.ReadFile(resolved)
		if err != nil {
			r.Errors = append(r.Errors, err.Error())
			errs = append(errs, fmt.Sprintf("evidence '%s' could not be read", key))
			results[key] = r
			continue
		}

		if key == "trace" {
			var events []any
			if err := json.Unmarshal(bytes, &events); err != nil {
				f := false
				r.SchemaValid = &f
				r.Errors = append(r.Errors, "trace must be a JSON array")
				errs = append(errs, "trace evidence must be a JSON array")
				results[key] = r
				continue
			}
			traceErrs := []string{}
			for idx, ev := range events {
				ok, schemaErrs := schema.ValidateAgainstSchema(schemaPath, ev)
				if !ok {
					for _, e := range schemaErrs {
						traceErrs = append(traceErrs, fmt.Sprintf("[index %d] %s", idx, e))
					}
				}
			}
			if len(traceErrs) > 0 {
				f := false
				r.SchemaValid = &f
				r.Errors = append(r.Errors, traceErrs...)
				errs = append(errs, fmt.Sprintf("evidence '%s' failed schema validation", key))
			} else {
				t := true
				r.SchemaValid = &t
			}
			results[key] = r
			continue
		}

		var parsed any
		if err := json.Unmarshal(bytes, &parsed); err != nil {
			f := false
			r.SchemaValid = &f
			r.Errors = append(r.Errors, "schema-backed evidence must be valid JSON")
			errs = append(errs, fmt.Sprintf("evidence '%s' is not valid JSON", key))
			results[key] = r
			continue
		}

		ok, schemaErrs := schema.ValidateAgainstSchema(schemaPath, parsed)
		if !ok {
			f := false
			r.SchemaValid = &f
			r.Errors = append(r.Errors, schemaErrs...)
			errs = append(errs, fmt.Sprintf("evidence '%s' failed schema validation", key))
		} else {
			t := true
			r.SchemaValid = &t
		}
		results[key] = r
	}

	return results, errs, warns
}
