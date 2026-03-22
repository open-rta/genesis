# Open RTA Genesis

## 1) What Genesis is

Open RTA Genesis is the canonical law, specification, conformance, and governance repository for Open RTA (Reliable Traceable Agents).

Genesis is intentionally **spec-first**. It defines law and conformance expectations, then provides schemas and examples to make those expectations testable.

## 2) What Genesis is not

Genesis is not:
- a runtime
- an SDK
- a framework
- a control plane
- a hosted validator service
- a replacement for manual/foundation certification review

Genesis does **not** define runtime architecture, internal repository layout, or implementation-specific engineering patterns.

## 3) The four Open RTA conditions

1. **No Rogue Autonomy**: autonomous action must have attributable origin.
2. **Observable Execution**: execution must be inspectable during or after operation.
3. **Interruptible Control**: authority must be able to intervene according to runtime constraints.
4. **Replayable Behavior**: execution must leave evidence for reconstruction or meaningful replay.

## 4) Manifest-first conformance model

Conformance is manifest-first:
- runtime authors publish an `open-rta-manifest.json`
- the manifest points to evidence artifacts
- validators/reviewers resolve and assess evidence from manifest references

> Open RTA requires discoverable evidence, not prescribed internal layout.

Evidence may live anywhere in a runtime repository (or exported evidence bundle). What matters is discoverability and conformance, not folder naming conventions.

## 5) How runtime authors use Genesis

1. Read laws in [`laws/`](laws/) and normative requirements in [`spec/`](spec/).
2. Produce an `open-rta-manifest.json` in your runtime repository.
3. Reference your evidence artifacts from that manifest.
4. Run the local support validator under [`tools/validator/`](tools/validator/).
5. Use validator output as pre-review input for manual/foundation review.

A runtime does not need to import Genesis code to be Open RTA-conformant.

## 6) What the validator does

The validator (currently inside Genesis, intentionally isolated under `tools/validator`) performs basic automatic checks:
- manifest existence + JSON parse
- manifest schema validation
- local evidence reference resolution
- schema validation for known evidence artifact categories
- basic law claim/evidence consistency checks
- L0-L3 prerequisite checks
- machine-readable validation report output

## 7) What the validator does not do

The validator does not:
- run runtime code
- inspect internal runtime architecture
- prove operational safety or production reliability
- prove oversight/control quality in real operations
- issue L4/foundation certification

## 8) Certification vs validator output

A passing validator report indicates **automatic pre-check readiness** only.

Validator pass != foundation certificate.

Manual/foundation review remains required for higher-assurance judgments and formal certification outcomes.

## 9) Repository structure guide

- [`laws/`](laws/): law-level constraints.
- [`concepts/`](concepts/): shared conceptual vocabulary.
- [`spec/`](spec/): normative requirements and validation scope boundaries.
- [`schemas/`](schemas/): machine-readable contracts.
- [`examples/`](examples/): illustrative manifest/evidence/report examples.
- [`governance/`](governance/): law and spec evolution process.
- [`tools/validator/`](tools/validator/): support tooling (Go validator), intentionally isolated.

## 10) Future ecosystem note

Validator tooling lives inside Genesis for now to support adoption, but it is deliberately isolated so it can be split later (for example to a dedicated `open-rta-validator` repository) with minimal disruption to the law/spec source of truth.
