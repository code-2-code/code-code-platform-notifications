# Agent Rules

- This repository owns notification dispatch and inbound callback handling only.
- Do not edit protobuf source or generated contracts here; change `code-code-contracts` first.
- Do not add provider, auth, agent runtime, egress, model catalog, profile, or deployment behavior here.
- Keep credentials and callback secrets out of logs, tests, docs, and committed config.
- Run the notification service tests before delivery.
