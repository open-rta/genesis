import * as fs from "node:fs";
import * as path from "node:path";
import { readJsonFile } from "./utils";
import { Manifest, ValidationContext } from "./types";

export function loadManifest(manifestPathInput: string, repoRoot: string): { manifest: Manifest; context: ValidationContext } {
  const manifestPath = path.resolve(manifestPathInput);

  if (!fs.existsSync(manifestPath)) {
    throw new Error(`Manifest file not found: ${manifestPath}`);
  }

  let parsed: unknown;
  try {
    parsed = readJsonFile(manifestPath);
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    throw new Error(`Failed to parse manifest JSON at ${manifestPath}: ${message}`);
  }

  const context: ValidationContext = {
    repoRoot,
    manifestPath,
    manifestDir: path.dirname(manifestPath)
  };

  return { manifest: parsed as Manifest, context };
}
