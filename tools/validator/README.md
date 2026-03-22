# Open RTA Validator (Genesis)

This tool provides lightweight, local, automatic validation for `open-rta-manifest.json` and referenced evidence artifacts.

Genesis defines the laws and conformance expectations of Open RTA. It does not define runtime architecture.

> Open RTA requires discoverable evidence, not prescribed internal layout.

## What this validator does

- validates manifest presence and JSON structure
- validates manifest against `schemas/open-rta-manifest.schema.json`
- resolves local evidence references
- validates known schema-backed artifacts
- checks basic law/evidence consistency
- checks claimed level prerequisites for L0-L3
- writes a machine-readable validation report

## What it does not do

- does not run runtime code
- does not introspect deployment internals
- does not evaluate operational quality or safety
- does not grant L4 or formal certification

## Build and run

From `tools/validator/`:

```bash
go build -o bin/open-rta-validate ./cmd/open-rta-validate
./bin/open-rta-validate <path-to-open-rta-manifest.json>
```

Optional report output:

```bash
./bin/open-rta-validate <manifest-path> --report <report-output.json>
```

Without building, you can also run directly:

```bash
go run ./cmd/open-rta-validate ../../examples/open-rta-manifest.json --report ../../examples/validation-report.generated.json
```

## Example

```bash
./bin/open-rta-validate ../../examples/open-rta-manifest.json --report ../../examples/validation-report.generated.json
```

## Output shape

The JSON report includes:
- `manifest_path`
- `runtime_name`
- `runtime_version`
- `claimed_level`
- `validated_level`
- `passed`
- `errors[]`
- `warnings[]`
- `evidence_results{}`
- `notes[]`

## Certification note

A passing validator report indicates basic conformance readiness only. It is not equivalent to foundation/manual certification issuance.
