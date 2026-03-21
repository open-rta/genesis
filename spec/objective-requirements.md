# Objective Requirements

An objective declaration MUST include:
- `objective_id`
- `description`
- `authority_ref`
- timestamp field (`created_at`)
- `status`

An objective declaration MAY include:
- `scope` or boundary metadata

`objective_id` MUST be stable enough to link trace, control, oversight, and replay artifacts for the same objective.
