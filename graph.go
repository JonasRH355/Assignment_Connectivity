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
	delete(g.adj, v)

	for key, neighbors := range g.adj {
		newList := []string{}
		for _, n := range neighbors {
			if n != v {
				newList = append(newList, n)
			}
		}
		g.adj[key] = newList
	}
}

func (g *Graph) AddEdge(u, v string) {
	g.adj[u] = append(g.adj[u], v)

	if !g.directed {
		g.adj[v] = append(g.adj[v], u)
	}
}

func (g *Graph) RemoveEdge(u, v string) {
	newList := []string{}
	for _, n := range g.adj[u] {
		if n != v {
			newList = append(newList, n)
		}
	}
	g.adj[u] = newList

	if !g.directed {
		newList = []string{}
		for _, n := range g.adj[v] {
			if n != u {
				newList = append(newList, n)
			}
		}
		g.adj[v] = newList
	}
}
