# Law 3: Interruptible Control

## Law statement
Authority must be able to pause, stop, or otherwise intervene in execution according to runtime constraints.

## Rationale
Autonomy without intervention capability creates unacceptable operational and governance risk.

## Minimum operational meaning
A runtime provides a control pathway to issue intervention commands and records resulting effects.

## Examples
- An authorized operator issues a stop instruction for an objective in progress.
- A policy engine applies an override that constrains subsequent actions.

## Non-examples
- Execution cannot be halted once started.
- Control commands exist but have no observable effect or record.

## What this law does NOT imply
- It does not mandate exact verbs (`pause`, `stop`, `override`) if equivalents exist.
- It does not guarantee instantaneous interruption in all physical systems.
- It does not prescribe control interface type.
