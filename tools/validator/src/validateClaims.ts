import { Manifest } from "./types";

export function validateClaimConsistency(manifest: Manifest): string[] {
  const errors: string[] = [];
  const claims = new Set(manifest.compliance?.laws_claimed ?? []);
  const evidenceKeys = new Set(Object.keys(manifest.evidence ?? {}));

  if (claims.has("no-rogue-autonomy")) {
    if (!evidenceKeys.has("objective") || !evidenceKeys.has("authority")) {
      errors.push("Claim 'no-rogue-autonomy' requires objective and authority evidence");
    }
  }

  if (claims.has("observable-execution") && !evidenceKeys.has("trace")) {
    errors.push("Claim 'observable-execution' requires trace evidence");
  }

  if (claims.has("interruptible-control") && !evidenceKeys.has("control")) {
    errors.push("Claim 'interruptible-control' requires control evidence");
  }

  if (claims.has("replayable-behavior") && !evidenceKeys.has("replay")) {
    errors.push("Claim 'replayable-behavior' requires replay evidence");
  }

  return errors;
}
