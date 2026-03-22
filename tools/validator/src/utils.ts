import * as fs from "node:fs";
import * as path from "node:path";

export function readJsonFile(filePath: string): unknown {
  const raw = fs.readFileSync(filePath, "utf8");
  return JSON.parse(raw);
}

export function isLikelyUrl(ref: string): boolean {
  return /^https?:\/\//i.test(ref);
}

export function resolveLocalPath(baseDir: string, ref: string): string {
  if (path.isAbsolute(ref)) {
    return path.normalize(ref);
  }
  return path.resolve(baseDir, ref);
}

export function pathExists(filePath: string): boolean {
  return fs.existsSync(filePath);
}

export function formatAjvErrors(errors: Array<{ instancePath?: string; message?: string }> | null | undefined): string[] {
  if (!errors) {
    return [];
  }

  return errors.map((error) => {
    const where = error.instancePath && error.instancePath.length > 0 ? error.instancePath : "/";
    return `${where}: ${error.message ?? "validation error"}`;
  });
}
