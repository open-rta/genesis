package schema

import "path/filepath"

var artifactSchemas = map[string]string{
	"objective":          "schemas/objective.schema.json",
	"authority":          "schemas/authority.schema.json",
	"trace":              "schemas/trace-event.schema.json",
	"replay":             "schemas/replay-record.schema.json",
	"oversight":          "schemas/oversight-record.schema.json",
	"control":            "schemas/control-record.schema.json",
	"conformance_report": "schemas/conformance-report.schema.json",
}

func ArtifactSchemaPath(repoRoot, artifact string) (string, bool) {
	rel, ok := artifactSchemas[artifact]
	if !ok {
		return "", false
	}
	return filepath.Join(repoRoot, rel), true
}

func ManifestSchemaPath(repoRoot string) string {
	return filepath.Join(repoRoot, "schemas", "open-rta-manifest.schema.json")
}
