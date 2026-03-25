package main

import (
	"fmt"
)

func main() {
	obj := NewGraph(true)
	obj.AddVertex("a")
	obj.AddVertex("b")
	obj.AddVertex("c")

	obj.AddEdge("a", "b")
	obj.AddEdge("b", "c")
	obj.AddEdge("a", "c")

	fmt.Println(obj.adj)
}
