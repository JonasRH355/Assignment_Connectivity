package main

import (
	"fmt"
)

func main() {
	obj := NewGraph(true)
	// obj.AddVertex("a")
	// obj.AddVertex("b")
	// obj.AddVertex("c")
	// obj.AddVertex("d")

	// obj.AddEdge("a", "b")
	// obj.AddEdge("b", "d")
	// obj.AddEdge("a", "c")

	// fmt.Println(obj.adj)

	// visited := make(map[string]bool)
	// obj.DFS("a", visited)

	// fmt.Println("BFS:")
	// obj.BFS("a")

	// rev := obj.Reverse()
	// rev.TransitiveClosure("a")

	// fmt.Println("REV:")
	// rev.Print()

	// fmt.Println("is conected:")
	// obj.IsConnected()

	// fmt.Println("InverseTransitiveClosure:")
	// obj.InverseTransitiveClosure("a")
	// fmt.Println("InverseTransitiveClosure:")
	// rev.InverseTransitiveClosure("a")
	for {
		fmt.Println("\n1. Add Vertex")
		fmt.Println("2. Add Edge")
		fmt.Println("3. DFS")
		fmt.Println("4. BFS")
		fmt.Println("5. Transitive closure")
		fmt.Println("6. Check connectivity")
		fmt.Println("7. Print graph")
		fmt.Println("8. Remove Vertex")
		fmt.Println("9. Remove Edge")
		fmt.Println("0. close")

		var op int
		fmt.Scan(&op)

		switch op {
		case 1:
			var v string
			fmt.Print("Vertex: ")
			fmt.Scan(&v)
			obj.AddVertex(v)

		case 2:
			var u, v string
			fmt.Print("Edge (u v): ")
			fmt.Scan(&u, &v)
			obj.AddEdge(u, v)

		case 3:
			var start string
			fmt.Print("DFS start: ")
			fmt.Scan(&start)
			visited := make(map[string]bool)
			obj.DFS(start, visited)

		case 4:
			var start string
			fmt.Print("BFS start: ")
			fmt.Scan(&start)
			obj.BFS(start)

		case 5:
			var start string
			fmt.Print("Vertex: ")
			fmt.Scan(&start)
			obj.TransitiveClosure(start)

		case 6:
			fmt.Println("Connected:", obj.IsConnected())

		case 7:

			fmt.Print("Graph: ")
			obj.Print()

		case 8:
			var v string
			fmt.Print("Vertex: ")
			fmt.Scan(&v)
			obj.RemoveVertex(v)

		case 9:
			var u, v string
			fmt.Print("Edge (u v): ")
			fmt.Scan(&u, &v)
			obj.RemoveEdge(u, v)

		case 0:
			return
		}
	}
}
