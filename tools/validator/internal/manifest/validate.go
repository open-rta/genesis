package manifest

import "github.com/open-rta/genesis/tools/validator/internal/schema"

func ValidateSchema(m Manifest, repoRoot string) (bool, []string) {
	errs := schema.ValidateDataAgainstSchema(m, schema.ManifestSchemaPath(repoRoot))
	return len(errs) == 0, errs
}
