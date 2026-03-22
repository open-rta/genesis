import * as fs from "node:fs";
import * as path from "node:path";
import { loadManifest } from "./loadManifest";
import { printReportSummary } from "./report";
import { resolveAndValidateEvidence } from "./resolveEvidence";
import { Manifest, ValidationReport } from "./types";
import { validateClaimConsistency } from "./validateClaims";
import { validateLevel } from "./validateLevels";
import { validateManifestSchema } from "./validateSchema";

function findRepoRoot(startDir: string): string {
  let current = path.resolve(startDir);

  while (true) {
    const hasSchemas = fs.existsSync(path.join(current, "schemas"));
    const hasLaws = fs.existsSync(path.join(current, "laws"));
    if (hasSchemas && hasLaws) {
      return current;
    }

    const parent = path.dirname(current);
    if (parent === current) {
      throw new Error("Could not locate Genesis repository root from current directory.");
    }
    current = parent;
  }
}

export function runValidation(manifestPathInput: string): ValidationReport {
  const repoRoot = findRepoRoot(process.cwd());
  const { manifest, context } = loadManifest(manifestPathInput, repoRoot);

  const errors: string[] = [];
  const warnings: string[] = [];
  const notes: string[] = [];

  const manifestSchemaResult = validateManifestSchema(manifest, repoRoot);
  if (!manifestSchemaResult.valid) {
    errors.push(...manifestSchemaResult.errors.map((entry) => `Manifest schema: ${entry}`));
  }

  const evidenceValidation = resolveAndValidateEvidence(manifest as Manifest, context);
  errors.push(...evidenceValidation.errors);
  warnings.push(...evidenceValidation.warnings);

  const claimErrors = validateClaimConsistency(manifest as Manifest);
  errors.push(...claimErrors);

  const levelValidation = validateLevel(manifest as Manifest);
  errors.push(...levelValidation.errors);
  notes.push(...levelValidation.notes);

  const report: ValidationReport = {
    manifest_path: context.manifestPath,
    runtime_name: manifest.runtime?.name ?? "unknown",
    runtime_version: manifest.runtime?.version ?? "unknown",
    claimed_level: manifest.compliance?.level_claimed ?? "L0",
    validated_level: levelValidation.validatedLevel,
    passed: errors.length === 0,
    errors,
    warnings,
    evidence_results: evidenceValidation.results,
    notes,
    summary: {
      checked_evidence_count: Object.keys(evidenceValidation.results).length,
      schema_validated_count: Object.values(evidenceValidation.results).filter((result) => result.schema_valid === true).length,
      hard_error_count: errors.length,
      warning_count: warnings.length
    }
  };

  return report;
}

export function runAndPrint(manifestPathInput: string): ValidationReport {
  const report = runValidation(manifestPathInput);
  printReportSummary(report);
  return report;
}
