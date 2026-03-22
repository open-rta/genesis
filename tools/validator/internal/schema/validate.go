package schema

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSONFileAgainstSchema(jsonFilePath string, schemaFilePath string) error {
	absJSONPath, err := filepath.Abs(jsonFilePath)
	if err != nil {
		return fmt.Errorf("could not resolve JSON path %q: %w", jsonFilePath, err)
	}

	jsonBytes, err := os.ReadFile(absJSONPath)
	if err != nil {
		return fmt.Errorf("could not read JSON file %q: %w", absJSONPath, err)
	}

	var payload any
	if err := json.Unmarshal(jsonBytes, &payload); err != nil {
		return fmt.Errorf("invalid JSON in %q: %w", absJSONPath, err)
	}

	errs := ValidateDataAgainstSchema(payload, schemaFilePath)
	if len(errs) == 0 {
		return nil
	}
	return fmt.Errorf(strings.Join(errs, "; "))
}

func ValidateDataAgainstSchema(data any, schemaFilePath string) []string {
	schemaLoader, err := SchemaLoader(schemaFilePath)
	if err != nil {
		return []string{fmt.Sprintf("schema loader error (%s): %v", schemaFilePath, err)}
	}

	result, err := gojsonschema.Validate(schemaLoader, gojsonschema.NewGoLoader(data))
	if err != nil {
		return []string{fmt.Sprintf("schema validation runtime error (schema: %s): %v", schemaFilePath, err)}
	}

	if result.Valid() {
		return nil
	}

	details := make([]string, 0, len(result.Errors()))
	for _, schemaErr := range result.Errors() {
		details = append(details, fmt.Sprintf("schema=%s field=%s error=%s", schemaFilePath, schemaErr.Field(), schemaErr.Description()))
	}
	return details
}
