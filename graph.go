package main

type Graph struct {
	adj      map[string][]string
	directed bool
}

// Graph Constructor
func NewGraph(directed bool) *Graph {
	return &Graph{
		adj:      make(map[string][]string),
		directed: directed,
	}
}

func (g *Graph) AddVertex(v string) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = []string{}
	}
}

func (g *Graph) RemoveVertex(v string) {
	return
}

func (g *Graph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)

	if !g.directed {
		g.adj[u] = append(g.adj[v], u)
	}
}

func (g *Graph) RemoveEdge() { return }
