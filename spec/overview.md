# Specification Overview

Open RTA Genesis is structured as:

1. **Laws**: minimal constraints on autonomous behavior.
2. **Concepts**: precise terms used by laws and requirements.
3. **Requirements**: normative expectations for runtime-exposed artifacts.
4. **Schemas**: machine-readable contracts for artifact structures.
5. **Evidence**: concrete exported records demonstrating conformance.
6. **Validation guidance**: automatic checks and manual review boundaries.

The stack is directional: laws -> concepts -> requirements -> schemas -> evidence.

Genesis defines the laws and conformance expectations of Open RTA. It does not define runtime architecture.

## Key specification documents

- [`manifest-requirements.md`](manifest-requirements.md)
- [`automatic-tests.md`](automatic-tests.md)
- [`manual-review.md`](manual-review.md)
- [`certification-levels.md`](certification-levels.md)
- [`certificate-issuance.md`](certificate-issuance.md)
- [`validator-scope.md`](validator-scope.md)
