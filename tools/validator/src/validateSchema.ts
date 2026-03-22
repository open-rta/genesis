import * as path from "node:path";
import Ajv from "ajv";
import addFormats from "ajv-formats";
import { formatAjvErrors, readJsonFile } from "./utils";

const schemaPathByArtifact: Record<string, string> = {
  objective: "schemas/objective.schema.json",
  authority: "schemas/authority.schema.json",
  trace: "schemas/trace-event.schema.json",
  replay: "schemas/replay-record.schema.json",
  oversight: "schemas/oversight-record.schema.json",
  control: "schemas/control-record.schema.json",
  conformance_report: "schemas/conformance-report.schema.json"
};

export function getSchemaPathForArtifact(artifactKey: string, repoRoot: string): string | undefined {
  const relative = schemaPathByArtifact[artifactKey];
  if (!relative) {
    return undefined;
  }
  return path.resolve(repoRoot, relative);
}

export function validateAgainstSchema(data: unknown, schemaPath: string): { valid: boolean; errors: string[] } {
  const schema = readJsonFile(schemaPath);

  const ajv = new Ajv({ allErrors: true, strict: false });
  addFormats(ajv);
  const validate = ajv.compile(schema);
  const valid = validate(data);

  return {
    valid: Boolean(valid),
    errors: formatAjvErrors(validate.errors as Array<{ instancePath?: string; message?: string }> | null)
  };
}

export function validateManifestSchema(manifest: unknown, repoRoot: string): { valid: boolean; errors: string[] } {
  const manifestSchemaPath = path.resolve(repoRoot, "schemas/open-rta-manifest.schema.json");
  return validateAgainstSchema(manifest, manifestSchemaPath);
}
