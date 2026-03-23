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

func (g *Graph) AddVertex() {
	return
}

func (g *Graph) RemoveVertex() { return }

func (g *Graph) AddEdge() { return }

func (g *Graph) RemoveEdge() { return }
