# Open RTA Genesis

Open RTA Genesis is the canonical law, specification, and conformance repository for Open RTA (Reliable Traceable Agents).

Genesis defines the laws and conformance expectations of Open RTA. It does not define runtime architecture.

> Open RTA requires discoverable evidence, not prescribed internal layout.

## What Genesis is

Genesis defines and publishes:
- foundational laws
- core concepts
- normative conformance requirements
- machine-readable JSON Schemas
- examples and governance process
- a lightweight validator tool for pre-review conformance checks

## What Genesis is not

Genesis is not:
- a runtime
- an SDK
- an agent framework
- a control plane
- a hosted service
- a final certification authority

## How runtime authors use Genesis

1. Read laws in [`laws/`](laws/) and requirements in [`spec/`](spec/).
2. Produce an `open-rta-manifest.json` in your own repository.
3. Reference your own evidence artifacts from the manifest.
4. Run the local validator from [`tools/validator/`](tools/validator/).
5. Use the report as a pre-certification readiness input before manual/foundation review.

Genesis does not require importing Genesis code into your runtime.

## Where key pieces live

- Laws: [`laws/`](laws/)
- Spec requirements: [`spec/`](spec/)
- Schemas: [`schemas/`](schemas/)
- Examples: [`examples/`](examples/)
- Validator tool: [`tools/validator/`](tools/validator/)

## Automatic validation workflow

From `tools/validator/`:

```bash
npm install
npm run build
node dist/cli.js ../../examples/open-rta-manifest.json --report ../../examples/validation-report.generated.json
```

The validator supports local filesystem evidence references in this initial release.

## What automatic validation covers

- manifest presence and JSON parseability
- manifest schema validation (`schemas/open-rta-manifest.schema.json`)
- evidence reference resolution for local files
- schema validation of known evidence artifacts:
  - objective
  - authority
  - trace
  - replay
  - oversight
  - control
  - conformance_report
- basic laws/evidence consistency checks
- claimed compliance level prerequisite checks (L0-L3)

See [`spec/automatic-tests.md`](spec/automatic-tests.md) for normative details.

## What automatic validation does not cover

Automatic checks do **not** prove:
- operational safety or reliability in production
- quality or meaningfulness of oversight
- practical effectiveness of controls
- replay usefulness beyond structural evidence
- semantic truthfulness of all evidence contents
- L4 certification status

These are handled by manual/foundation review. See:
- [`spec/manual-review.md`](spec/manual-review.md)
- [`spec/certificate-issuance.md`](spec/certificate-issuance.md)

## Compliance levels at a glance

- **L0 — Declared**: valid manifest + attestation
- **L1 — Attributable + Observable**: L0 + objective + authority + trace
- **L2 — Controlled**: L1 + control
- **L3 — Replayable**: L2 + replay
- **L4 — Reviewed / Foundation-certified**: manual review required

Full requirements: [`spec/certification-levels.md`](spec/certification-levels.md).

## Recommended reading order

1. [`CHARTER.md`](CHARTER.md)
2. [`laws/`](laws/)
3. [`concepts/`](concepts/)
4. [`spec/manifest-requirements.md`](spec/manifest-requirements.md)
5. [`spec/conformance-model.md`](spec/conformance-model.md)
6. [`schemas/`](schemas/)
7. [`tools/validator/README.md`](tools/validator/README.md)
8. [`examples/`](examples/)
