# Open RTA Validator (Support Tooling)

This Go tool provides support workflows for Open RTA runtime authors:

1. `init-manifest` — generate a starter `open-rta-manifest.json`
2. `validate` — validate a manifest and referenced evidence

Genesis remains spec-first; this tool is intentionally secondary support tooling.

## Build / run

From `tools/validator`:

```bash
go build -o bin/open-rta .
./bin/open-rta validate --manifest ../../examples/open-rta-manifest.json
```

Or without building:

```bash
go run . validate --manifest ../../examples/open-rta-manifest.json
```

## Command: init-manifest

### Interactive mode

```bash
go run . init-manifest
```

The tool prompts for runtime identity, compliance claim, evidence refs, attestation, limitations, and output path.

### Flag-based mode

```bash
go run . init-manifest \
  --non-interactive \
  --runtime-name "ExampleRuntime" \
  --runtime-version "0.1.0" \
  --open-rta-version "0.1.0" \
  --level L1 \
  --laws "no-rogue-autonomy,observable-execution" \
  --declared-by "compliance@example.org" \
  --objective-ref "./evidence/objective.json" \
  --authority-ref "./evidence/authority.json" \
  --trace-ref "./evidence/trace.json" \
  --output ./open-rta-manifest.json
```

### Key flags

- `--output`
- `--force`
- `--non-interactive`
- `--runtime-name`
- `--runtime-version`
- `--open-rta-version`
- `--level`
- `--description`
- `--repository`
- `--homepage`
- `--declared-by`
- `--contact`
- `--laws`
- `--applies-to`
- `--exclusions`
- `--objective-ref`
- `--authority-ref`
- `--trace-ref`
- `--control-ref`
- `--oversight-ref`
- `--replay-ref`
- `--conformance-ref`
- `--limitations`

### Overwrite protection

`init-manifest` will not overwrite an existing file unless:
- interactive confirmation is given, or
- `--force` is passed.

## Command: validate

```bash
go run . validate --manifest ../../examples/open-rta-manifest.json --report ../../examples/validation-report.generated.json
```

Automatic checks:
- manifest parse + schema validation
- evidence existence (for local refs)
- known evidence schema validation
- laws/evidence consistency
- level prerequisite checks (L0-L3)

## Important limitation

Generated manifests are a **starting point**, not automatic compliance.

Validator pass != foundation certificate.

Manual/foundation review is still required.
