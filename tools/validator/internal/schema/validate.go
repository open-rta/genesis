package schema

import (
	"fmt"
	"os"
	"strings"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

func ValidateAgainstSchema(schemaPath string, data any) (bool, []string) {
	schemaBytes, err := os.ReadFile(schemaPath)
	if err != nil {
		return false, []string{fmt.Sprintf("schema read error: %v", err)}
	}

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource(schemaPath, strings.NewReader(string(schemaBytes))); err != nil {
		return false, []string{fmt.Sprintf("schema load error: %v", err)}
	}

	compiled, err := compiler.Compile(schemaPath)
	if err != nil {
		return false, []string{fmt.Sprintf("schema compile error: %v", err)}
	}

	if err := compiled.Validate(data); err != nil {
		return false, []string{err.Error()}
	}
	return true, nil
}
