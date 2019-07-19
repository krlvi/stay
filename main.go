package main

import "fmt"
import "stay/graph"

func main() {
	scc := graph.FindStronglyConnectedComponents(testGraph())
	fmt.Println("SCC: ", scc)
}

func testGraph() graph.Graph {
	g := graph.NewGraph()
	g.Push("foo")
	g.Push("bar")
	g.Push("baz")
	g.Push("moo")
	g.Connect("foo", "bar")
	g.Connect("bar", "baz")
	g.Connect("baz", "foo")
	g.Connect("foo", "moo")
	return g
}
