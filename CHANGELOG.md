# Changelog

All notable changes to this repository are documented here.

## v0.2.1
- Refactored validator tooling to a staged structure under `tools/validator` with clear package separation (`manifest`, `evidence`, `levels`, `report`, `schema`).
- Removed legacy TypeScript validator artifacts and kept a single primary Go validator implementation.
- Clarified root and governance docs so Genesis remains spec-first and tooling remains support-only.
- Added explicit staged-positioning language for future extraction of validator tooling to a dedicated repo.

## v0.2.0
- Added a manifest-based validator tool under `tools/validator` implemented in Go.
- Added `schemas/open-rta-manifest.schema.json`.
- Added `schemas/open-rta-certificate.schema.json`.
- Added automatic validation and manual review specification documents.
- Added certification-level and certificate-issuance documentation.
- Added `examples/open-rta-manifest.json` and `examples/validation-report.json`.
- Updated README and examples to a manifest-first, layout-neutral model.

## v0.1.0
- Initial charter.
- Initial laws.
- Initial concepts.
- Initial specification documents.
- Initial JSON Schemas.
- Initial examples.
