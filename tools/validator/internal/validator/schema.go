package validator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

var artifactSchemas = map[string]string{
	"objective":          "schemas/objective.schema.json",
	"authority":          "schemas/authority.schema.json",
	"trace":              "schemas/trace-event.schema.json",
	"replay":             "schemas/replay-record.schema.json",
	"oversight":          "schemas/oversight-record.schema.json",
	"control":            "schemas/control-record.schema.json",
	"conformance_report": "schemas/conformance-report.schema.json",
}

func schemaPathForArtifact(repoRoot, artifact string) (string, bool) {
	rel, ok := artifactSchemas[artifact]
	if !ok {
		return "", false
	}
	return filepath.Join(repoRoot, rel), true
}

func validateDataAgainstSchema(schemaPath string, data any) (bool, []string) {
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return false, []string{fmt.Sprintf("schema read error: %v", err)}
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(schemaPath, strings.NewReader(string(schemaBytes))); err != nil {
		return false, []string{fmt.Sprintf("schema load error: %v", err)}
	}
	schema, err := compiler.Compile(schemaPath)
	if err != nil {
		return false, []string{fmt.Sprintf("schema compile error: %v", err)}
	}
	if err := schema.Validate(data); err != nil {
		return false, []string{err.Error()}
	}
	return true, nil
}
