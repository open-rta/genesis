package evidence

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
