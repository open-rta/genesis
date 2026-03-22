# Manifest Requirements

## Purpose

`open-rta-manifest.json` is the runtime author's declaration and evidence index for Open RTA conformance checks.

It is the manifest-first entrypoint for automatic tooling and manual/foundation review.

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

Canonical schema: [`../schemas/open-rta-manifest.schema.json`](../schemas/open-rta-manifest.schema.json).

## Evidence model

The manifest references evidence artifacts by logical keys (for example `objective`, `trace`, `control`) and includes metadata such as:
- reference `type`
- reference `ref`
- optional `schema`
- optional `format`

Evidence may live anywhere in a runtime repository or exported evidence bundle if it remains discoverable.

## Reference types

Supported evidence reference types:
- `file`
- `artifact`
- `doc`
- `url`

## Generation support

Genesis validator tooling supports creating starter manifests via `init-manifest` (interactive or flag-based) and then validating them via `validate`.

Generated manifests are starter declarations and should be reviewed/edited as needed.

## Validator support scope (0.1.1)

Automatic validation in 0.1.1 focuses on local file-based evidence resolution and schema-backed conformance checks.

Remote retrieval (`url`) remains out of scope for automatic retrieval and is handled in manual review contexts.
