package algraph

import (
	"testing"
	"reflect"
)

func TestAdjacent(t *testing.T) {
	a, b, c, d := Node{"a"}, Node{"b"}, Node{"c"}, Node{"d"}
	neighbors := map[Node][]Node{
		a: {b},
		b: {a, c},
		c: {b, d},
		d: {c},
	}
	costs := map[Node]map[Node]float32{}
	g := Graph{neighbors: neighbors, costs: costs}

	successAdjacent := func(x Node, y Node) {
		r, _ := g.Adjacent(x, y)
		if r != true {
			t.Errorf("Expected true but got false for nodes '%+v' and '%+v'", x, y)
		}
	}
	successItems := [][]Node{
		{a, b}, {b, a},
		{b, c}, {c, b},
		{c, d}, {d, c},
	}
	for _, nodes := range successItems {
		successAdjacent(nodes[0], nodes[1])
	}

	failAdjacent := func(x Node, y Node) {
		r, _ := g.Adjacent(x, y)
		if r == true {
			t.Errorf("Expected false but got true for nodes '%+v' and '%+v'", x, y)
		}
	}
	failItems := [][]Node{
		{a, c}, {b, d},
		{c, a}, {d, b},
	}
	for _, nodes := range failItems {
		failAdjacent(nodes[0], nodes[1])
	}
}

func TestNeighbors(t *testing.T) {
	a, b, c, d := Node{"a"}, Node{"b"}, Node{"c"}, Node{"d"}
	neighbors := map[Node][]Node{
		a: {b},
		b: {a, c},
		c: {b, d},
		d: {c},
	}
	costs := map[Node]map[Node]float32{}
	g := Graph{neighbors: neighbors, costs: costs}

	for node, expectNeighs := range neighbors {
		resultNeighs, _ := g.Neighbors(node)

		if !reflect.DeepEqual(expectNeighs, resultNeighs) {
			t.Errorf("Expected '%+v' but got '%+v' for node '%+v'",
				expectNeighs, resultNeighs, node)
		}
	}

	// check an absent node
	singleNode := Node{"notingraph"}
	resultEmpty, ok := g.Neighbors(singleNode)
	if resultEmpty != nil && ok == nil {
		t.Errorf("Expected an empty neighbors list but got '%+v' for node '%+v'",
			resultEmpty, singleNode)
	}
}

func TestAddVertex(t *testing.T) {
	a, b := Node{"a"}, Node{"b"}
	neighbors := map[Node][]Node{
		a: {b},
	}
	costs := map[Node]map[Node]float32{}
	g := Graph{neighbors: neighbors, costs: costs}

	newNode := Node{"c"}
	g.AddVertex(newNode)

	nodeList, err := g.Neighbors(newNode)
	if nodeList == nil || err != nil {
		t.Errorf("Can't find the node '%+v' int the graph '%+v'",
			newNode, g)
	}
}

func TestRemoveVertex(t *testing.T) {
	
}

func TestAddEdge(t *testing.T) {
	
}

func TestRemoveEdge(t *testing.T) {
	
}

func TestGetEdgeValue(t *testing.T) {
	
}

func TestSetEdgeValue(t *testing.T) {
	
}
