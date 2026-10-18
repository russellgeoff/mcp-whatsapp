# Read-only build (Russell_NAS)

This fork builds an image of [Sealjay/mcp-whatsapp](https://github.com/Sealjay/mcp-whatsapp)
that can only **read** WhatsApp. It is used as a Claude connector on a home NAS.

What differs from upstream (the `readonly` branch, based on tag `v0.6.0`):

- `internal/mcp/readonly.go` — after the tools are registered, every tool not
  on a short allowlist is deleted. Exposed: `list_chats`, `get_chat`,
  `list_messages`, `get_message_context`, `search_contacts`, `list_groups`,
  `get_group_info`, `request_sync`, `get_status`, `pairing_status`.
  Removed: everything that sends, reacts, edits, deletes, marks as read,
  changes groups, blocks, sets presence, privacy or profile, checks numbers,
  or downloads media.
- `internal/mcp/server.go` — calls `applyReadOnly()` and no longer registers
  the `whatsapp://media/…` resource template.
- `internal/mcp/readonly_test.go` — fails if the exposed tool set differs from
  the allowlist (for example if an upstream update adds or renames a tool).
- `.github/workflows/build-readonly-image.yml` — runs that test, builds the
  image for linux/amd64, boot-tests it, and publishes
  `ghcr.io/<owner>/mcp-whatsapp:<tag>` when a `v*-ro*` tag is pushed.

To update to a new upstream release: rebase `readonly` onto the new upstream
tag, push, then push a new tag such as `v0.7.0-ro1`.
