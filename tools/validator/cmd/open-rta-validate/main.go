package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/open-rta/genesis/tools/validator/internal/validator"
)

func main() {
	reportPath := flag.String("report", "", "Path to write machine-readable validation report JSON")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "Usage: open-rta-validate <manifest-path> [--report <report-output.json>]")
		os.Exit(2)
	}

	report, err := validator.Validate(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		os.Exit(1)
	}

	if *reportPath != "" {
		bytes, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not encode report: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(*reportPath, append(bytes, '\n'), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "Could not write report: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Report written to %s\n", *reportPath)
	}

	if report.Passed {
		os.Exit(0)
	}
	os.Exit(1)
}
