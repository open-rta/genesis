package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/open-rta/genesis/tools/validator/internal/report"
	"github.com/open-rta/genesis/tools/validator/internal/validator"
)

func main() {
	manifestPath := flag.String("manifest", "", "Path to open-rta-manifest.json")
	reportPath := flag.String("report", "", "Path to write machine-readable validation report JSON")
	flag.Parse()

	if *manifestPath == "" && flag.NArg() > 0 {
		*manifestPath = flag.Arg(0)
	}

	if *manifestPath == "" {
		fmt.Fprintln(os.Stderr, "Usage: go run ./tools/validator --manifest ./examples/open-rta-manifest.json [--report ./out.json]")
		os.Exit(2)
	}

	r, err := validator.Run(*manifestPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Validation failed: %v\n", err)
		os.Exit(1)
	}

	if *reportPath != "" {
		out := *reportPath
		if !filepath.IsAbs(out) {
			out = filepath.Clean(out)
		}
		if err := report.WriteJSON(out, r); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Report written to %s\n", out)
	}

	if r.Passed {
		os.Exit(0)
	}
	os.Exit(1)
}
