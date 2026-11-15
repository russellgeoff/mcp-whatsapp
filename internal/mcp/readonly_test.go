package mcp

import (
	"sort"
	"testing"
)

// TestReadOnlyToolSet checks that the read-only build exposes exactly the
// tools in readOnlyTools: nothing that sends, edits, deletes or changes
// settings, and none of the read tools missing. It checks every way the
// server can be built, so upstream's own modes (read-only mode, the opt-in
// WhatsApp Business label tools) cannot add anything to the list.
func TestReadOnlyToolSet(t *testing.T) {
	variants := []struct {
		name string
		srv  *Server
	}{
		{"default", NewServer(nil, nil, false)},
		{"labels enabled", NewServer(nil, nil, true)},
		{"upstream read-only mode", NewReadOnlyServer(nil, nil, false)},
		{"upstream read-only mode, labels enabled", NewReadOnlyServer(nil, nil, true)},
	}
	for _, v := range variants {
		t.Run(v.name, func(t *testing.T) {
			got := v.srv.MCP().ListTools()

			var names []string
			for name := range got {
				names = append(names, name)
				if !readOnlyTools[name] {
					t.Errorf("tool %q is exposed but is not in the read-only allowlist", name)
				}
			}
			for name := range readOnlyTools {
				if _, ok := got[name]; !ok {
					t.Errorf("read-only tool %q is missing (renamed upstream?)", name)
				}
			}
			sort.Strings(names)
			t.Logf("exposed tools (%d): %v", len(names), names)
		})
	}
}
