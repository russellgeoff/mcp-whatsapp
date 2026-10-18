package mcp

// Read-only build (Russell_NAS, 2026-10-17).
//
// This build exposes only the tools below. Every other tool (sending,
// reacting, editing, deleting, marking read, group admin, blocking,
// presence, privacy and profile changes, number lookups, media download)
// is removed from the server after registration, so it cannot be called
// even by a client that holds the bearer token. The media resource
// template is not registered either.
var readOnlyTools = map[string]bool{
	"list_chats":          true,
	"get_chat":            true,
	"list_messages":       true,
	"get_message_context": true,
	"search_contacts":     true,
	"list_groups":         true,
	"get_group_info":      true,
	"request_sync":        true,
	"get_status":          true,
	"pairing_status":      true,
}

// applyReadOnly deletes every registered tool that is not in readOnlyTools.
func (s *Server) applyReadOnly() {
	var remove []string
	for name := range s.mcp.ListTools() {
		if !readOnlyTools[name] {
			remove = append(remove, name)
		}
	}
	s.mcp.DeleteTools(remove...)
}
