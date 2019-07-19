package graph

import (
	"errors"
	"fmt"
)

type stringSet map[string]struct{}

type Graph struct {
	nodes map[string]stringSet
}

func (g Graph) Push(node string) {
	if _, x := g.nodes[node]; !x {
		g.nodes[node] = make(stringSet)
	}
}

func (g Graph) Connect(from string, to string) error {
	if from == to {
		return errors.New("can't connect to self")
	}
	if _, x := g.nodes[from]; !x {
		return errors.New("from node must be in the graph")
	}
	if _, x := g.nodes[to]; !x {
		return errors.New("to node must be in the graph")
	}
	g.nodes[from][to] = struct{}{}
	return nil
}

func (g Graph) edgesFrom(from string) stringSet {
	return g.nodes[from]
}

func NewGraph() Graph {
	return Graph{nodes: make(map[string]stringSet)}
}

func FindStronglyConnectedComponents(g Graph) map[int][]string {
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
		markReachableNodes(node, g, result, i)
		i++
		visitOrder = visitOrder[:n]
	}

	scc := make(map[int][]string)
	for k, v := range result {
		scc[v] = append(scc[v], k)
	}
	return scc
}

func markReachableNodes(node string, g Graph, result map[string]int, label int) {
	if _, x := result[node]; x {
		return
	}
	result[node] = label
	for e := range g.edgesFrom(node) {
		markReachableNodes(e, g, result, label)
	}
}

func dfsVisitOrder(g Graph) []string {
	var stack = make([]string, 0)
	var visited = make(stringSet)
	for n := range g.nodes {
		recExplore(n, &g, &stack, visited)
	}
	return stack
}

func recExplore(node string, g *Graph, stack *[]string, visited stringSet) {
	if _, x := visited[node]; x {
		return
	}
	visited[node] = struct{}{}
	for e := range g.edgesFrom(node) {
		recExplore(e, g, stack, visited)
	}
	*stack = append(*stack, node)
}

func transpose(g Graph) Graph {
	t := Graph{nodes: make(map[string]stringSet)}
	for k := range g.nodes {
		t.Push(k)
	}
	for node := range g.nodes {
		for e := range g.edgesFrom(node) {
			conErr := t.Connect(e, node)
			if conErr != nil {
				fmt.Println(conErr)
			}
		}
	}
	return t
}
