# Certification Levels

Open RTA levels provide an incremental conformance model.

## L0 — Declared

Requirements:
- valid `open-rta-manifest.json`
- valid attestation section

## L1 — Attributable + Observable

Requirements:
- L0
- objective evidence
- authority evidence
- trace evidence

## L2 — Controlled

Requirements:
- L1
- control evidence

## L3 — Replayable

Requirements:
- L2
- replay evidence

## L4 — Reviewed / Foundation-certified

Requirements:
- L3 baseline artifacts
- successful manual/foundation review under governance policy

## Automatic evidence checks by level

The validator can automatically verify prerequisite evidence for L0-L3.

L4 cannot be automatically granted by validator output alone.
