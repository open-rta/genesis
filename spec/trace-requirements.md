# Trace Requirements

Trace events MUST include at least:
- `event_id`
- `objective_id`
- `timestamp`
- `actor` or component identifier
- `event_type`
- `status` or outcome indicator

Trace events SHOULD include:
- `action_id` when a specific action is represented
- concise event summary

Trace exports MUST allow linking execution steps to objective context.
