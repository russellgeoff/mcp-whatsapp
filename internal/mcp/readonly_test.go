package mcp

import (
	"sort"
	"testing"
)

// TestReadOnlyToolSet checks that the read-only build exposes exactly the
// tools in readOnlyTools: nothing that sends, edits, deletes or changes
// settings, and none of the read tools missing.
func TestReadOnlyToolSet(t *testing.T) {
	s := NewServer(nil, nil)
	got := s.MCP().ListTools()

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
}
