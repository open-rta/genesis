package schema

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
)

func SchemaLoader(schemaFilePath string) (gojsonschema.JSONLoader, error) {
	absPath, err := filepath.Abs(schemaFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not resolve schema path %q: %w", schemaFilePath, err)
	}
	if _, err := os.Stat(absPath); err != nil {
		return nil, fmt.Errorf("schema file not found %q: %w", absPath, err)
	}
	return gojsonschema.NewReferenceLoader("file://" + filepath.ToSlash(absPath)), nil
}
