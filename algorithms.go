package main

import (
	"fmt"
)

func (g *Graph) DFS(start string, visited map[string]bool) {
	visited[start] = true
	fmt.Println(start)

	for _, neighbor := range g.adj[start] {
		if !visited[neighbor] {
			g.DFS(neighbor, visited)
		}
	}
}

func (g *Graph) BFS(start string) {
	visited := make(map[string]bool)
	queue := []string{start}

	visited[start] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node)

		for _, neighbor := range g.adj[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
}

func (g *Graph) IsConnected() bool {
	visited := make(map[string]bool)

	var start string
	for k := range g.adj {
		start = k
		break
	}

	g.DFS(start, visited)

	return len(visited) == len(g.adj)
}

func (g *Graph) TransitiveClosure(start string) {
	visited := make(map[string]bool)
	g.DFS(start, visited)
}

func (g *Graph) InverseTransitiveClosure(start string) {
	rev := g.Reverse()
	visited := make(map[string]bool)
	rev.DFS(start, visited)
}
