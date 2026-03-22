package validator

import "path/filepath"

func validateManifestSchema(m Manifest, repoRoot string) (bool, []string) {
	schemaPath := filepath.Join(repoRoot, "schemas", "open-rta-manifest.schema.json")
	return validateDataAgainstSchema(schemaPath, m)
}
