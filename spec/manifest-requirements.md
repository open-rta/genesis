# Manifest Requirements

## Purpose

`open-rta-manifest.json` is the runtime author's declaration and evidence index for Open RTA conformance checks.

It is a portable entrypoint that lets validators and reviewers discover evidence without prescribing runtime internals.

> Open RTA requires discoverable evidence, not prescribed internal layout.

## Required sections

A conformant manifest includes:
- `manifest_version`
- `open_rta_version`
- `runtime`
- `compliance`
- `evidence`
- `attestation`

Optional but recommended:
- `test_targets`
- `limitations`

See [`../schemas/open-rta-manifest.schema.json`](../schemas/open-rta-manifest.schema.json).

## Evidence reference model

The manifest references evidence artifacts by logical key (for example, `objective`, `trace`, `control`) and metadata including:
- reference `type`
- reference `ref`
- optional schema hint and digest metadata

Evidence may reside anywhere in the runtime repository or adjacent export location, provided paths are resolvable to the reviewer/validator context.

## Reference types

Manifest evidence entries support these `type` values:
- `file` (local file path)
- `artifact` (named runtime artifact, usually a local file reference)
- `doc` (documentation artifact)
- `url` (future/optional remote reference)

## Validator support scope (v0.2)

The v0.2 validator supports **local file-based** resolution for `file`, `artifact`, and `doc` references.

Remote retrieval (`url`) is out of scope for automatic checks in v0.2 and remains a manual review concern.
