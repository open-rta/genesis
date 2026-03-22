package manifest

import "github.com/open-rta/genesis/tools/validator/internal/schema"

func ValidateSchema(m Manifest, repoRoot string) (bool, []string) {
	return schema.ValidateAgainstSchema(schema.ManifestSchemaPath(repoRoot), m)
}
