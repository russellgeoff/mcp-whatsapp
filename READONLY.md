# Read-only build (Russell_NAS)

This fork builds an image of [Sealjay/mcp-whatsapp](https://github.com/Sealjay/mcp-whatsapp)
that can only **read** WhatsApp. It is used as a Claude connector on a home NAS.

What differs from upstream (the `readonly` branch: tag `v0.6.0` plus the changes
below, with upstream tag `v0.7.2` merged in on 2026-11-15):

- `internal/mcp/readonly.go` — after the tools are registered, every tool not
  on a short allowlist is deleted. Exposed: `list_chats`, `get_chat`,
  `list_messages`, `get_message_context`, `search_contacts`, `list_groups`,
  `get_group_info`, `request_sync`, `get_status`, `pairing_status`.
  Removed: everything that sends, reacts, edits, deletes, marks as read,
  changes groups, blocks, sets presence, privacy or profile, checks numbers,
  or downloads media.
- `internal/mcp/server.go` — calls `applyReadOnly()` after upstream's own
  read-only filter (`serve -read-only`, 15 tools from v0.7.0), whatever the
  mode, and no longer registers the `whatsapp://media/…` resource template.
  So this build stays at the 10 tools above even where upstream's read-only
  mode would keep media downloads, poll results or the label tools.
- `internal/mcp/readonly_test.go` — fails if the exposed tool set differs from
  the allowlist (for example if an upstream update adds or renames a tool),
  checked in all four ways the server can be built (default, labels on,
  upstream read-only mode, both).
- `.github/workflows/build-readonly-image.yml` — on every push to `readonly`,
  runs that test, builds the image for linux/amd64, boot-tests it, and
  publishes `ghcr.io/<owner>/mcp-whatsapp:<version>`.
- `internal/client/client.go` — pairing fix (from `v0.6.0-ro2`): after a
  pairing succeeds (phone-number code or QR), `login` now stays connected until the
  post-pairing reconnect completes, then 30 s more. Upstream exits 0.5 s after
  success, which left the phone reporting a failed link and WhatsApp removing
  the device ("device_removed") on the next connect. Still needed at v0.7.2,
  which keeps the 0.5 s exit (its pairing fix, #77, is a different one).
- `READONLY_VERSION` — the image tag to publish (`v0.7.2-ro1`; earlier
  `v0.6.0-ro1`, `v0.6.0-ro2`).

To update to a new upstream release: merge the new upstream tag into
`readonly` (merging keeps the branch's history, so no force-push is needed),
set `READONLY_VERSION` (e.g. `v0.7.3-ro1`) and push the branch.
