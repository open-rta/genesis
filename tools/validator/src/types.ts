export type Level = "L0" | "L1" | "L2" | "L3" | "L4";

export interface EvidenceReference {
  type: "file" | "artifact" | "doc" | "url";
  ref: string;
  schema?: string;
  description?: string;
  required?: boolean;
}

export interface Manifest {
  manifest_version: string;
  open_rta_version: string;
  runtime: {
    name: string;
    version: string;
    vendor?: string;
    homepage?: string;
  };
  compliance: {
    laws_claimed: string[];
    level_claimed: Level;
  };
  evidence: Record<string, EvidenceReference>;
  test_targets?: Array<{ id: string; description: string }>;
  attestation: {
    by: string;
    timestamp: string;
    statement: string;
    signature?: string;
  };
  limitations?: string[];
}

export interface EvidenceResult {
  found: boolean;
  path?: string;
  schema_valid?: boolean;
  errors?: string[];
  warnings?: string[];
}

export interface ValidationReport {
  manifest_path: string;
  runtime_name: string;
  runtime_version: string;
  claimed_level: Level;
  validated_level: Level;
  passed: boolean;
  errors: string[];
  warnings: string[];
  evidence_results: Record<string, EvidenceResult>;
  notes: string[];
  summary: {
    checked_evidence_count: number;
    schema_validated_count: number;
    hard_error_count: number;
    warning_count: number;
  };
}

export interface ValidationContext {
  repoRoot: string;
  manifestPath: string;
  manifestDir: string;
}
