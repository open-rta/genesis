import { ValidationReport } from "./types";

export function printReportSummary(report: ValidationReport): void {
  console.log("Open RTA Validation Summary");
  console.log("===========================");
  console.log(`Manifest: ${report.manifest_path}`);
  console.log(`Runtime: ${report.runtime_name} @ ${report.runtime_version}`);
  console.log(`Claimed level: ${report.claimed_level}`);
  console.log(`Validated level: ${report.validated_level}`);
  console.log(`Passed: ${report.passed ? "yes" : "no"}`);

  if (report.errors.length > 0) {
    console.log("Errors:");
    report.errors.forEach((error) => console.log(`  - ${error}`));
  }

  if (report.warnings.length > 0) {
    console.log("Warnings:");
    report.warnings.forEach((warning) => console.log(`  - ${warning}`));
  }

  if (report.notes.length > 0) {
    console.log("Notes:");
    report.notes.forEach((note) => console.log(`  - ${note}`));
  }
}
