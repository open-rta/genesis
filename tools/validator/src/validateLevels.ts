import { Level, Manifest } from "./types";

const orderedLevels: Level[] = ["L0", "L1", "L2", "L3"];

function hasEvidence(manifest: Manifest, key: string): boolean {
  return Boolean(manifest.evidence && manifest.evidence[key]);
}

function qualifiesForLevel(manifest: Manifest, level: Exclude<Level, "L4">): boolean {
  const hasAttestation = Boolean(manifest.attestation?.by && manifest.attestation?.timestamp && manifest.attestation?.statement);

  if (level === "L0") {
    return hasAttestation;
  }
  if (level === "L1") {
    return hasAttestation && hasEvidence(manifest, "objective") && hasEvidence(manifest, "authority") && hasEvidence(manifest, "trace");
  }
  if (level === "L2") {
    return qualifiesForLevel(manifest, "L1") && hasEvidence(manifest, "control");
  }

  return qualifiesForLevel(manifest, "L2") && hasEvidence(manifest, "replay");
}

export function validateLevel(manifest: Manifest): { validatedLevel: Level; errors: string[]; notes: string[] } {
  const claimed = manifest.compliance.level_claimed;
  const errors: string[] = [];
  const notes: string[] = [];

  let validated: Level = "L0";
  for (const level of orderedLevels) {
    if (qualifiesForLevel(manifest, level)) {
      validated = level;
    }
  }

  if (claimed === "L4") {
    notes.push("L4 requires manual/foundation review and cannot be granted automatically.");
  } else {
    const claimedIndex = orderedLevels.indexOf(claimed as Exclude<Level, "L4">);
    const validatedIndex = orderedLevels.indexOf(validated as Exclude<Level, "L4">);
    if (claimedIndex >= 0 && validatedIndex < claimedIndex) {
      errors.push(`Claimed level ${claimed} is not supported by provided prerequisite evidence (validated: ${validated})`);
    }
  }

  return { validatedLevel: validated, errors, notes };
}
