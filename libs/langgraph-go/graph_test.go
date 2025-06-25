package langgraph

import "testing"

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	if g == nil || len(g.Nodes) != 0 {
		t.Fatalf("unexpected graph: %+v", g)
	}
}
