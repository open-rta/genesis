# Examples

This directory provides minimal Open RTA evidence artifacts and a manifest-first validation flow.

## Core artifact examples

- `minimal-objective.json`: objective declaration linked to authority.
- `minimal-authority.json`: authority lineage and constraints.
- `minimal-trace.json`: trace events showing execution progression.
- `minimal-replay-record.json`: reconstruction metadata referencing trace events.
- `minimal-oversight-record.json`: oversight observation record.
- `minimal-control-record.json`: intervention record with observed effect.
- `minimal-conformance-report.json`: runtime declaration of Open RTA alignment.

Legacy YAML variants are included for readability:
- `minimal-objective.yaml`
- `minimal-authority.yaml`

## Manifest + report examples

- `open-rta-manifest.json`: sample runtime manifest referencing local evidence artifacts.
- `validation-report.json`: example machine-readable output from the validator.

## How these relate to real runtimes

These examples demonstrate structure and discoverability expectations only.

Real runtimes may store evidence in any internal layout. Open RTA requires discoverable evidence, not prescribed internal folder names.

See `compliant-vs-noncompliant.md` for practical interpretation.
