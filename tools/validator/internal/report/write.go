package report

import (
	"encoding/json"
	"fmt"
	"os"
)

func WriteJSON(path string, r ValidationReport) error {
	bytes, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return fmt.Errorf("could not encode report: %w", err)
	}
	if err := os.WriteFile(path, append(bytes, '\n'), 0o644); err != nil {
		return fmt.Errorf("could not write report: %w", err)
	}
	return nil
}
