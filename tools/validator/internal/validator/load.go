package validator

import (
	"fmt"
	"path/filepath"
)

type Context struct {
	RepoRoot     string
	ManifestPath string
	ManifestDir  string
}

func findRepoRoot(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		hasSchemas := fileExists(filepath.Join(current, "schemas"))
		hasLaws := fileExists(filepath.Join(current, "laws"))
		if hasSchemas && hasLaws {
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", fmt.Errorf("could not locate Genesis repository root")
		}
		current = parent
	}
}

func LoadManifest(manifestPathInput string) (Manifest, Context, error) {
	var m Manifest
	repoRoot, err := findRepoRoot(".")
	if err != nil {
		return m, Context{}, err
	}
	absManifest, err := filepath.Abs(manifestPathInput)
	if err != nil {
		return m, Context{}, err
	}
	if !fileExists(absManifest) {
		return m, Context{}, fmt.Errorf("manifest file not found: %s", absManifest)
	}
	if err := readJSONFile(absManifest, &m); err != nil {
		return m, Context{}, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}
	ctx := Context{RepoRoot: repoRoot, ManifestPath: absManifest, ManifestDir: filepath.Dir(absManifest)}
	return m, ctx, nil
}
