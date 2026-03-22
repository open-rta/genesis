# Automatic Validation Tests

This document defines what the Genesis validator checks automatically.

## Automatic checks

1. Manifest file exists.
2. Manifest parses as JSON.
3. Manifest validates against `open-rta-manifest.schema.json`.
4. Evidence references are resolved for supported local references.
5. Known schema-backed evidence artifacts are validated when present:
   - objective
   - authority
   - trace
   - replay
   - oversight
   - control
   - conformance_report
6. Basic claim/evidence consistency is checked:
   - `no-rogue-autonomy` requires objective + authority evidence
   - `observable-execution` requires trace evidence
   - `interruptible-control` requires control evidence
   - `replayable-behavior` requires replay evidence
7. Claimed compliance level prerequisites are checked:
   - L0: valid manifest + attestation
   - L1: L0 + objective + authority + trace
   - L2: L1 + control
   - L3: L2 + replay
   - L4: cannot be auto-granted

## Out of scope for automatic checks

Automatic validation does not establish:
- true operational quality
- meaningfulness of oversight practice
- real-world safety
- practical effectiveness of controls in production
- replay usefulness quality beyond basic structure
- semantic truthfulness of evidence beyond structural/basic consistency
- formal foundation certification issuance

See [`manual-review.md`](manual-review.md) and [`certificate-issuance.md`](certificate-issuance.md).
