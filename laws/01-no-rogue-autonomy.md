# Law 1: No Rogue Autonomy

## Law statement
No autonomous action should occur without a declared objective, authority boundary, or attributable origin.

## Rationale
Unattributed autonomous behavior prevents accountability, governance, and safe operation across domains.

## Minimum operational meaning
At runtime or post-hoc, each material action can be linked to:
- an objective identifier
- an authority reference
- an actor or component identity

## Examples
- A runtime executes `action_id=A-42` with `objective_id=OBJ-100` and `authority_ref=AUTH-9`.
- A manually approved objective is logged before autonomous execution begins.

## Non-examples
- A system initiates external actions with no objective declaration.
- Actions are logged but cannot be associated with any authority source.

## What this law does NOT imply
- It does not require human approval for every action.
- It does not prescribe a single authority model.
- It does not prohibit automation; it requires attributable automation.
