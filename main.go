package main

import "fmt"

type StringSet map[string]struct{}

type Graph struct {
	Nodes map[string]StringSet
}

func (g Graph) Push(node string) {
	if _, x := g.Nodes[node]; !x {
		g.Nodes[node] = make(StringSet)
	}
}

func (g Graph) Connect(from string, to string) {
	g.Nodes[from][to] = struct{}{}
}

func (g Graph) EdgesFrom(from string) StringSet {
	return g.Nodes[from]
}

func findSCC(g Graph) map[string]int {
	visitOrder := dfsVisitOrder(transpose(g))
	result := make(map[string]int)

	i := 0
	for len(visitOrder) > 0 {
		n := len(visitOrder) - 1
		node := visitOrder[n]
		if _, x := result[node]; x {
			visitOrder = visitOrder[:n]
			continue
		}
		markReachableNodes(node, &g, &result, i)
		i++
		visitOrder = visitOrder[:n]
	}
	return result
}

func markReachableNodes(node string, g *Graph, result *map[string]int, label int) {
	if _, x := (*result)[node]; x {
		return
	}
	(*result)[node] = label
	for e := range g.EdgesFrom(node) {
		markReachableNodes(e, g, result, label)
	}
}

func dfsVisitOrder(g Graph) []string {
	var stack = make([]string, 0)
	var visited = make(StringSet)
	for n := range g.Nodes {
		recExplore(n, &g, &stack, &visited)
	}
	return stack
}

func recExplore(node string, g *Graph, stack *[]string, visited *StringSet) {
	if _, x := (*visited)[node]; x {
		return
	}
	(*visited)[node] = struct{}{}
	for e := range g.EdgesFrom(node) {
		recExplore(e, g, stack, visited)
	}
	*stack = append(*stack, node)
}

func transpose(g Graph) Graph {
	t := Graph{Nodes: make(map[string]StringSet)}
	for k := range g.Nodes {
		t.Push(k)
	}
	for node := range g.Nodes {
		for e := range g.EdgesFrom(node) {
			t.Connect(e, node)
		}
	}
	return t
}

func main() {
	g := Graph{Nodes: make(map[string]StringSet)}
	g.Push("foo")
	g.Push("bar")
	g.Push("baz")
	g.Push("moo")
	g.Connect("foo", "bar")
	g.Connect("bar", "baz")
	g.Connect("baz", "foo")
	g.Connect("foo", "moo")
	fmt.Println("SCC: ", findSCC(g))
}
