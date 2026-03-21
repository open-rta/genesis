# Law 4: Replayable Behavior

## Law statement
Autonomous execution must leave enough evidence to reconstruct or meaningfully replay what happened.

## Rationale
Without reconstruction evidence, verification, audit, and continuous improvement are degraded.

## Minimum operational meaning
The runtime preserves records linking objectives, actions, and outcomes so that a reviewer can reconstruct causal sequence and major decisions.

## Examples
- A replay record references source trace events and summarizes reconstruction scope.
- A reviewer can explain the action sequence that produced a final state.

## Non-examples
- Logs contain fragmented entries with no linkage to objective or action IDs.
- Only aggregate metrics are retained, with no event-level reconstruction path.

## What this law does NOT imply
- It does not require bit-exact deterministic replay for every runtime.
- It does not require full simulation of all external dependencies.
- It does not require exposing sensitive data beyond governance and legal constraints.
