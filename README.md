# Open RTA Genesis

Open RTA Genesis is the canonical law/spec/conformance repository for Open RTA.

Genesis is **spec-first**: it defines laws, concepts, conformance expectations, schemas, and governance. Tooling inside Genesis is support tooling, not the standard itself.

> Open RTA requires discoverable evidence, not prescribed internal layout.

## What Genesis is

- Open RTA laws and concepts
- conformance requirements and review boundaries
- canonical schemas and examples
- governance policy for law/spec evolution
- lightweight support tooling under `tools/validator`

## What Genesis is not

Genesis is not a runtime, framework, SDK, or control plane.

Genesis does not prescribe runtime architecture, internal repository structure, or engineering patterns.

## Four Open RTA conditions

1. No Rogue Autonomy
2. Observable Execution
3. Interruptible Control
4. Replayable Behavior

## Manifest-first conformance model

Runtime authors publish an `open-rta-manifest.json` that references evidence artifacts.

Those artifacts can live anywhere in a runtime repository; what matters is that evidence is discoverable and conforms to Open RTA expectations.

## First-time workflow (generate -> validate)

From `tools/validator`:

```bash
# 1) Create starter manifest interactively
go run . init-manifest

# 2) Validate it
go run . validate --manifest ./open-rta-manifest.json
```

Or generate non-interactively:

```bash
go run . init-manifest \
  --non-interactive \
  --runtime-name "ExampleRuntime" \
  --declared-by "compliance@example.org" \
  --level L1 \
  --laws "no-rogue-autonomy,observable-execution" \
  --output ./open-rta-manifest.json
```

## Validator scope (support tooling)

Validator automatic checks include:
- manifest existence + JSON parse
- manifest schema validation
- local evidence resolution
- schema validation for known evidence categories
- basic claims/evidence consistency checks
- L0-L3 prerequisite checks (L4 remains manual/foundation review)

## Certification boundary

Validator pass is pre-review readiness only.

Validator pass != foundation certificate.

Manual/foundation review remains required for higher-trust judgments and formal certification outcomes.

## Repo structure

- `laws/` — legal constraints
- `concepts/` — shared definitions
- `spec/` — normative requirements and review boundaries
- `schemas/` — canonical contracts
- `examples/` — illustrative manifests/evidence/reports
- `governance/` — evolution process and policy
- `tools/validator/` — isolated support tooling (Go)
