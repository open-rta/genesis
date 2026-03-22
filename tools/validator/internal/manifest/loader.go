package manifest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func findRepoRoot(start string) (string, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", err
	}
	for {
		if fileExists(filepath.Join(current, "schemas")) && fileExists(filepath.Join(current, "laws")) {
			return current, nil
		}
		parent := filepath.Dir(current)
		if parent == current {
			return "", fmt.Errorf("could not locate Genesis repository root")
		}
		current = parent
	}
}

func Load(manifestPathInput string) (Manifest, Context, error) {
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

	bytes, err := os.ReadFile(absManifest)
	if err != nil {
		return m, Context{}, fmt.Errorf("manifest read failed: %w", err)
	}
	if err := json.Unmarshal(bytes, &m); err != nil {
		return m, Context{}, fmt.Errorf("failed to parse manifest JSON: %w", err)
	}

	ctx := Context{RepoRoot: repoRoot, ManifestPath: absManifest, ManifestDir: filepath.Dir(absManifest)}
	return m, ctx, nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
