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

func (g *Graph) TransitiveClosure(start string) {
	visited := make(map[string]bool)
	g.DFS(start, visited)
}
