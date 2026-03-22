package report

import "fmt"

func PrintSummary(r ValidationReport) {
	fmt.Println("Open RTA Validation Summary")
	fmt.Println("===========================")
	fmt.Printf("Manifest: %s\n", r.ManifestPath)
	fmt.Printf("Runtime: %s @ %s\n", r.RuntimeName, r.RuntimeVersion)
	fmt.Printf("Claimed level: %s\n", r.ClaimedLevel)
	fmt.Printf("Validated level: %s\n", r.ValidatedLevel)
	fmt.Printf("Passed: %v\n", r.Passed)
	if len(r.Errors) > 0 {
		fmt.Println("Errors:")
		for _, e := range r.Errors {
			fmt.Printf("  - %s\n", e)
		}
	}
	if len(r.Warnings) > 0 {
		fmt.Println("Warnings:")
		for _, w := range r.Warnings {
			fmt.Printf("  - %s\n", w)
		}
	}
	if len(r.Notes) > 0 {
		fmt.Println("Notes:")
		for _, n := range r.Notes {
			fmt.Printf("  - %s\n", n)
		}
	}
}
