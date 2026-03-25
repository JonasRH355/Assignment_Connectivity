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
