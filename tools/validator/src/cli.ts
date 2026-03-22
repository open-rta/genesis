#!/usr/bin/env node
import * as fs from "node:fs";
import * as path from "node:path";
import { runAndPrint } from "./index";

function parseArgs(argv: string[]): { manifestPath?: string; reportPath?: string } {
  const positional: string[] = [];
  let reportPath: string | undefined;

  for (let i = 0; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === "--report") {
      reportPath = argv[i + 1];
      i += 1;
      continue;
    }

    positional.push(arg);
  }

  return {
    manifestPath: positional[0],
    reportPath
  };
}

function main(): void {
  const { manifestPath, reportPath } = parseArgs(process.argv.slice(2));

  if (!manifestPath) {
    console.error("Usage: open-rta-validate <manifest-path> [--report <report-output.json>]");
    process.exit(2);
  }

  const manifestArg = manifestPath;

  try {
    const report = runAndPrint(manifestArg);

    if (reportPath) {
      const outPath = path.resolve(reportPath);
      fs.writeFileSync(outPath, `${JSON.stringify(report, null, 2)}\n`, "utf8");
      console.log(`Report written to ${outPath}`);
    }

    process.exit(report.passed ? 0 : 1);
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    console.error(`Validation failed: ${message}`);
    process.exit(1);
  }
}

main();
