import { EvidenceResult, Manifest, ValidationContext } from "./types";
import { getSchemaPathForArtifact, validateAgainstSchema } from "./validateSchema";
import { isLikelyUrl, pathExists, readJsonFile, resolveLocalPath } from "./utils";

export function resolveAndValidateEvidence(manifest: Manifest, context: ValidationContext): {
  results: Record<string, EvidenceResult>;
  errors: string[];
  warnings: string[];
} {
  const results: Record<string, EvidenceResult> = {};
  const errors: string[] = [];
  const warnings: string[] = [];

  for (const [artifactKey, evidence] of Object.entries(manifest.evidence ?? {})) {
    const evidenceResult: EvidenceResult = { found: false, errors: [], warnings: [] };

    if (isLikelyUrl(evidence.ref) || evidence.type === "url") {
      evidenceResult.warnings?.push("URL evidence is not automatically fetched in v0.2");
      warnings.push(`Evidence '${artifactKey}' uses URL reference and was not auto-validated: ${evidence.ref}`);
      results[artifactKey] = evidenceResult;
      continue;
    }

    const resolvedPath = resolveLocalPath(context.manifestDir, evidence.ref);
    evidenceResult.path = resolvedPath;

    if (!pathExists(resolvedPath)) {
      evidenceResult.errors?.push("Referenced evidence file does not exist");
      errors.push(`Evidence '${artifactKey}' not found at ${resolvedPath}`);
      results[artifactKey] = evidenceResult;
      continue;
    }

    evidenceResult.found = true;

    const schemaPath = getSchemaPathForArtifact(artifactKey, context.repoRoot);
    if (!schemaPath) {
      evidenceResult.warnings?.push("No known schema mapping for this evidence key");
      results[artifactKey] = evidenceResult;
      continue;
    }

    let parsed: unknown;
    try {
      parsed = readJsonFile(resolvedPath);
    } catch (error) {
      const message = error instanceof Error ? error.message : String(error);
      evidenceResult.schema_valid = false;
      evidenceResult.errors?.push(`Could not parse evidence as JSON: ${message}`);
      errors.push(`Evidence '${artifactKey}' at ${resolvedPath} is not valid JSON`);
      results[artifactKey] = evidenceResult;
      continue;
    }

    if (artifactKey === "trace") {
      if (!Array.isArray(parsed)) {
        evidenceResult.schema_valid = false;
        evidenceResult.errors?.push("Trace evidence must be a JSON array of trace events");
        errors.push(`Evidence '${artifactKey}' must be an array of trace events`);
        results[artifactKey] = evidenceResult;
        continue;
      }

      const traceErrors: string[] = [];
      for (let idx = 0; idx < parsed.length; idx += 1) {
        const item = parsed[idx];
        const schemaValidation = validateAgainstSchema(item, schemaPath);
        if (!schemaValidation.valid) {
          traceErrors.push(...schemaValidation.errors.map((entry) => `[index ${idx}] ${entry}`));
        }
      }

      if (traceErrors.length > 0) {
        evidenceResult.schema_valid = false;
        evidenceResult.errors?.push(...traceErrors);
        errors.push(`Evidence '${artifactKey}' failed schema validation`);
      } else {
        evidenceResult.schema_valid = true;
      }

      results[artifactKey] = evidenceResult;
      continue;
    }

    const schemaValidation = validateAgainstSchema(parsed, schemaPath);
    evidenceResult.schema_valid = schemaValidation.valid;
    if (!schemaValidation.valid) {
      evidenceResult.errors?.push(...schemaValidation.errors);
      errors.push(`Evidence '${artifactKey}' failed schema validation`);
    }

    results[artifactKey] = evidenceResult;
  }

  return { results, errors, warnings };
}
