package validator

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func readJSONFile(path string, out any) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isURL(ref string) bool {
	u, err := url.Parse(ref)
	if err != nil {
		return false
	}
	return strings.HasPrefix(u.Scheme, "http")
}

func resolveLocalPath(baseDir, ref string) string {
	if filepath.IsAbs(ref) {
		return filepath.Clean(ref)
	}
	return filepath.Clean(filepath.Join(baseDir, ref))
}
