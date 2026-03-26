package main

import (
	"fmt"
)

func main() {
	obj := NewGraph(true)
	obj.AddVertex("a")
	obj.AddVertex("b")
	obj.AddVertex("c")
	obj.AddVertex("d")

	obj.AddEdge("a", "b")
	obj.AddEdge("b", "d")
	obj.AddEdge("a", "c")

	fmt.Println(obj.adj)

	visited := make(map[string]bool)
	obj.DFS("a", visited)

	fmt.Println("BFS:")
	obj.BFS("a")

	rev := obj.Reverse()
	rev.TransitiveClosure("a")

	fmt.Println("REV:")
	rev.Print()
}
