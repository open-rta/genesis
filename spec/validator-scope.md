# Validator Scope

This document defines the role and boundaries of the Genesis validator.

## Role

The validator is support tooling for automatic pre-review conformance checks.

It helps runtime authors detect structural issues before manual/foundation review.

## In scope

- validating `open-rta-manifest.json` structure against schema
- resolving local evidence references
- validating known evidence artifacts against Open RTA schemas
- checking basic law/evidence consistency
- checking prerequisite level rules for L0-L3
- producing human-readable output and machine-readable report output

## Out of scope

- proving real-world safety and reliability
- proving operational quality of oversight/control practice
- semantic truth adjudication of all evidence content
- granting L4 or foundation certification
- prescribing runtime implementation architecture

## Positioning

Genesis remains the canonical law/spec/conformance repository.

Validator tooling is intentionally isolated under `tools/validator` and may later be extracted to a dedicated tooling repository without redefining Open RTA law.
