# Open RTA Charter

## Purpose
Autonomous systems increasingly act across software, enterprise, and physical domains, often without consistent foundations for accountability and reconstruction. Open RTA exists to define minimal, shared conditions for reliable autonomy across implementations.

## Scope
Open RTA defines minimal conditions under which autonomous execution remains attributable, observable, interruptible, and replayable.

## Foundational principle
The term RTA draws from a concept of underlying order. In physical systems, that order is inherent. In autonomous systems, it must be explicitly designed through objectives, authority boundaries, operational visibility, intervention capability, and trace evidence.

## The Four Laws
1. No Rogue Autonomy: autonomous action must have attributable origin.
2. Observable Execution: execution must be inspectable during or after operation.
3. Interruptible Control: authority must be able to intervene according to runtime constraints.
4. Replayable Behavior: execution must leave evidence for reconstruction or meaningful replay.

## Separation of law and implementation
Open RTA defines conditions, not implementations.

Genesis tooling (including validator tooling) supports conformance workflows but does not define law. Law and specification are normative; tooling is supportive and may evolve independently.

## Federation
Different runtimes and domains may independently implement Open RTA while adhering to shared laws, concepts, and conformance expectations.

## Non-goals
See [NON-GOALS.md](NON-GOALS.md).

## Evolution
See [GOVERNANCE.md](GOVERNANCE.md) and documents in [governance/](governance/).
