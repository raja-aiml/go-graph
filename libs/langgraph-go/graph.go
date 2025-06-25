package langgraph

// Graph represents a simple directed graph placeholder.
type Graph struct {
	Nodes []string
}

func NewGraph() *Graph {
	return &Graph{Nodes: []string{}}
}
