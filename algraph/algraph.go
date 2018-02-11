// Package algraph provides implementations of graph
// algorithms and related primitives such
// a graph data structure implementation.
package algraph

import (
	"errors"
	"fmt"
)

// Node struct represents graph node
// with some data. A more agile way
// in case of changing data type than
// specifing concrete data type in
// Graph structure
type Node struct {
	data string
}

// Graph contains neighbors and
// costs for each Node that
// is a part of the graph
type Graph struct {
	neighbors map[Node][]Node
	costs     map[Node]map[Node]float32
}

// A local function for implement some public methods.
// Return true is the specified list of Nodes contains
// the specified Node
func contains(lst []Node, node Node) bool {
	for _, n := range lst {
		if node == n {
			return true
		}
	}
	return false
}

func deleteInArray(lst []Node, node Node) ([]Node, error) {
	delIndex := -1
	for i, value := range lst {
		if value == node {
			delIndex = i
		}
	}
	if delIndex != -1 {
		return nil, errors.New(fmt.Sprintf("There is no %+v in %v", node, lst))
	}
	return append(lst[:delIndex], lst[delIndex+1:]...), nil
}

// Adjacent tests whether there is an edge
// from the vertex x to the vertex y
func (g *Graph) Adjacent(x Node, y Node) (bool, error) {
	if nodeList, ok := g.neighbors[x]; ok {
		return contains(nodeList, y), nil
	}
	return false, errors.New(fmt.Sprintf("There is no %+v in the node list", x))
}

// Neighbors returns a list of all vertices y such
// that there is an edge from the vertex x to the vertex y;
func (g *Graph) Neighbors(x Node) ([]Node, error) {
	if nodeList, ok := g.neighbors[x]; ok {
		return nodeList, nil
	}
	return nil, errors.New(fmt.Sprintf("There is no %+v in the node list", x))
}

// AddVertex adds the vertex x, if it is not there.
// Adds just a single vertex without any edge.
func (g *Graph) AddVertex(x Node) {
	g.neighbors[x] = []Node{}
	g.costs[x] = map[Node]float32{}
}

// RemoveVertex removes the vertex x, if it is there.
// Removes all edges as well.
func (g *Graph) RemoveVertex(x Node) error {
	if _, ok := g.neighbors[x]; ok {
		delete(g.neighbors, x)
		delete(g.costs, x)

		for nodeFrom, nodeLst := range g.neighbors {
			if contains(nodeLst, x) {
				g.neighbors[nodeFrom], _ = deleteInArray(nodeLst, x)
			}
		}

		for nodeFrom, nodeMap := range g.costs {
			if nodeMap[x] != 0 {
				delete(g.costs[nodeFrom], x)
				g.neighbors[nodeFrom], _ = deleteInArray(g.neighbors[nodeFrom], x)
			}
		}

		return nil
	}
	return errors.New(fmt.Sprintf("There is no %+v in the node list", x))
}

// AddEdge adds the edge from the vertex x to the vertex y.
// If x -> y egde already exists, the edge value will be reassigned.
// It returns an error of the value of the edge is zero.
func (g *Graph) AddEdge(x Node, y Node, v float32) error {
	if v == 0 {
		return errors.New(fmt.Sprintf("The specified value is %f but it can't be 0", x))
	}

	g.neighbors[x] = append(g.neighbors[x], y)
	if nodeMap, ok := g.costs[x]; ok {
		nodeMap[y] = v
	} else {
		g.costs[x] = map[Node]float32{y: v}
	}
	return nil
}

// RemoveEdge removes the edge from the vertex x to
// the vertex y, if it is there;
func (g *Graph) RemoveEdge(x Node, y Node) error {
	if nodeList, ok := g.neighbors[x]; ok {
		if contains(nodeList, y) {
			g.neighbors[x], _ = deleteInArray(g.neighbors[x], y)
		}
		return errors.New(fmt.Sprintf("There is no %+v in the node list", x))
	}

	if _, ok := g.costs[x]; ok {
		delete(g.costs[x], y)
	}
	return nil
}

// GetEdgeValue returns the value associated with the edge (x, y).
// It returns 0 in case of there is no edge x -> y and
// returns an error if x doesn't have any edge at all.
func (g *Graph) GetEdgeValue(x Node, y Node) (float32, error) {
	if nodeMap, ok := g.costs[x]; ok {
		// 0 is a default value, so it will be
		// return here if nodeMap doesn't have
		// y Node as a key
		return nodeMap[y], nil
	}
	return -1, errors.New(fmt.Sprintf("There is no %+v in the node list", x))
}

// SetEdgeValue sets the value associated with the edge (x, y) to v.
// It returns an erorr if there is no such edge.
func (g *Graph) SetEdgeValue(x Node, y Node, v float32) error {
	if nodeMap, ok := g.costs[x]; ok {
		if nodeMap[y] == 0 {
			nodeMap[y] = v
		}
		return errors.New("There is no x -> y edge in the graph")
	}
	return errors.New(fmt.Sprintf("There is no %+v in the node list", x))
}
