# Open RTA Genesis

## Short definition
Open RTA Genesis is the canonical law, specification, and conformance repository for Open RTA (Reliable Traceable Agents). It defines the minimum accountable conditions that autonomous runtimes must satisfy and the evidence artifacts they should expose from their own project roots.
RTA also references an underlying-order concept: in physical systems order is often inherent, while in autonomous systems it must be explicitly designed and evidenced.

## What Genesis is
Genesis defines:
- foundational laws
- core concepts
- normative conformance requirements
- machine-readable schemas
- examples and governance processes

Genesis is intended to be read by runtime authors, reviewers, and operators who need a neutral and durable source of truth.

## What Genesis is not
Genesis is:
- not a runtime
- not an SDK
- not a CLI
- not a workflow engine
- not a control plane

## The four conditions of Open RTA
- Every action must be attributable to a declared objective or authority.
- Execution must remain observable during or after operation.
- Authority must be able to intervene or stop execution.
- Execution must leave sufficient trace for replay or reconstruction.

## How a runtime becomes Open RTA aligned
A runtime does not import Genesis code. A runtime becomes Open RTA aligned by exposing required artifacts from its own project root and demonstrating that these artifacts satisfy Open RTA laws, requirements, and schemas.

## Required project-root artifacts
A runtime may use a project-root layout similar to:

```text
project-root/
└── rta/
    ├── objective.yaml
    ├── authority.yaml
    ├── trace/
    │   └── trace-events.json
    ├── replay/
    │   └── replay-record.json
    ├── oversight/
    │   └── oversight-record.json
    ├── control/
    │   └── control-record.json
    └── conformance/
        └── conformance-report.json
```

## Runtime author checklist
- [ ] Declare objectives.
- [ ] Maintain authority lineage.
- [ ] Export execution trace.
- [ ] Expose supervision and control evidence.
- [ ] Preserve replay or reconstruction evidence.
- [ ] Produce a conformance report.

## What Genesis does not prescribe
Genesis does not prescribe:
- runtime architecture
- orchestration style
- programming language
- storage engine
- transport protocol
- UI
- workflow DSL
- simulation engine
- model provider

## Repo navigation guide
- [`laws/`](laws/): minimal legal constraints for Open RTA.
- [`concepts/`](concepts/): precise definitions of terms used by the laws and spec.
- [`spec/`](spec/): normative requirements, conformance model, and evidence model.
- [`schemas/`](schemas/): JSON Schema contracts for required artifacts.
- [`examples/`](examples/): minimal data examples and compliant vs noncompliant comparison.
- [`governance/`](governance/): process documents for proposing and evolving the standard.
- Root docs:
  - [`CHARTER.md`](CHARTER.md)
  - [`GOVERNANCE.md`](GOVERNANCE.md)
  - [`NON-GOALS.md`](NON-GOALS.md)
  - [`ROADMAP.md`](ROADMAP.md)
  - [`CHANGELOG.md`](CHANGELOG.md)

## Conformance path
Recommended reading order for runtime authors:
1. [`CHARTER.md`](CHARTER.md)
2. [`laws/`](laws/)
3. [`concepts/`](concepts/)
4. [`spec/conformance-model.md`](spec/conformance-model.md)
5. Artifact requirement files in [`spec/`](spec/)
6. JSON Schemas in [`schemas/`](schemas/)
7. Concrete examples in [`examples/`](examples/)

## One-line summary
Genesis defines the laws and conformance expectations of Open RTA. It is not a runtime, framework, SDK, or control plane.
