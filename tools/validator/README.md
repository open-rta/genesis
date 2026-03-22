# Open RTA Validator (Genesis Support Tooling)

The Open RTA validator is a small Go tool that performs automatic pre-review checks for `open-rta-manifest.json` and referenced evidence artifacts.

Genesis remains the canonical law/spec/conformance repository. This validator is support tooling only.

## Why this tool lives inside Genesis (for now)

- to provide practical conformance pre-checks close to the canonical schemas/spec
- to improve adoption while keeping Open RTA law/spec central
- to preserve a staged path toward future extraction if tooling grows

The validator is intentionally isolated under `tools/validator` so it can later move to a dedicated repository (for example `open-rta-validator`) with minimal disruption.

## What the validator does

- checks manifest file existence and JSON parsing
- validates manifest against `schemas/open-rta-manifest.schema.json`
- resolves local evidence references
- validates known schema-backed evidence artifacts when possible
- checks basic claim/evidence consistency
- checks claimed level prerequisites for L0-L3
- writes machine-readable validation report output (optional)
- uses `github.com/xeipuuv/gojsonschema` in an isolated schema layer for JSON Schema checks

## What the validator does not do

- does not run or introspect runtime internals
- does not define runtime architecture or folder layout
- does not prove operational quality or real-world safety
- does not grant L4/foundation certification

Validator pass != foundation certificate.

## Build and run

From `tools/validator`:

```bash
go run . --manifest ../../examples/open-rta-manifest.json
```

Build binary (optional):

```bash
cd tools/validator
go build -o bin/open-rta-validate .
./bin/open-rta-validate --manifest ../../examples/open-rta-manifest.json
```

Write report JSON:

```bash
cd tools/validator && go run . --manifest ../../examples/open-rta-manifest.json --report ../../examples/validation-report.generated.json
```

## Output fields

Validation report fields include:
- `manifest_path`
- `runtime_name`
- `runtime_version`
- `claimed_level`
- `validated_level`
- `passed`
- `errors`
- `warnings`
- `evidence_results`
- `notes`
