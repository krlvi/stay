package graph

import "testing"
import "reflect"

func TestGraph_PushOk(t *testing.T) {
	g := NewGraph()
	g.Push("foo")
	if len(g.nodes) != 1 {
		t.Fail()
	}
	if _, x := g.nodes["foo"]; !x {
		t.Fail()
	}
}

func TestGraph_PushTwiceOk(t *testing.T) {
	g := NewGraph()
	g.Push("foo")
	g.Push("foo")
	if len(g.nodes) != 1 {
		t.Fail()
	}
}

func TestGraph_ConnectOk(t *testing.T) {
	g := NewGraph()
	g.Push("foo")
	g.Push("bar")
	if g.Connect("foo", "bar") != nil {
		t.Fail()
	}
	if _, x := g.edgesFrom("foo")["bar"]; !x {
		t.Fail()
	}
}

func TestGraph_ConnectFromNotInGraphErrors(t *testing.T) {
	g := NewGraph()
	g.Push("to")
	if g.Connect("from", "to") == nil {
		t.Fail()
	}
}

func TestGraph_ConnectToNotInGraphErrors(t *testing.T) {
	g := NewGraph()
	g.Push("from")
	if g.Connect("from", "to") == nil {
		t.Fail()
	}
}

func TestGraph_ConnectSelfErrors(t *testing.T) {
	g := NewGraph()
	g.Push("foo")
	if g.Connect("foo", "foo") == nil {
		t.Fail()
	}
}

func TestFindStronglyConnectedComponents(t *testing.T) {
	g := NewGraph()
	g.Push("foo")
	g.Push("bar")
	g.Push("baz")
	g.Push("moo")
	g.Connect("foo", "bar")
	g.Connect("bar", "baz")
	g.Connect("baz", "foo")
	g.Connect("foo", "moo")
	components := FindStronglyConnectedComponents(g)
	if !reflect.DeepEqual(components, map[int][]string{0: {"moo"}, 1: {"foo", "bar", "baz"}}) {
		t.Fail()
	}
}
