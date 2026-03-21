# Law 2: Observable Execution

## Law statement
Autonomous execution must be inspectable while running or reviewable after execution.

## Rationale
Operational trust and incident analysis require visibility into what occurred and why.

## Minimum operational meaning
A runtime provides trace records sufficient to inspect state transitions, actions, and outcomes during operation or after completion.

## Examples
- A trace stream exposes event records during execution.
- A trace export allows post-run review when live inspection is unavailable.

## Non-examples
- Only final success/failure is retained, with no intermediate evidence.
- Execution is opaque and cannot be reviewed after completion.

## What this law does NOT imply
- It does not require real-time dashboards.
- It does not require full observability of internal implementation details.
- It does not prescribe storage format beyond required artifact compatibility.
