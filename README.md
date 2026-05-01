# code-code-platform-notifications

Notification services split from the Code Code platform.

This repository owns:

- `packages/platform-k8s/internal/notificationdispatcher`: outbound notification delivery from platform events.
- `packages/platform-k8s/internal/wecomcallback`: verified WeCom inbound callback adapter.
- `packages/platform-k8s/cmd/notification-dispatcher`: notification dispatcher runtime entrypoint.
- `packages/platform-k8s/cmd/wecom-callback-adapter`: WeCom callback adapter runtime entrypoint.

Contracts are consumed through the `code-code-contracts` submodule.

Useful checks:

```bash
git submodule update --init --recursive
cd packages/platform-k8s
go test ./internal/notificationdispatcher/... ./internal/wecomcallback/... ./cmd/notification-dispatcher ./cmd/wecom-callback-adapter
```
