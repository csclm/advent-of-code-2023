package day22

import (
	"aoc-2023/aoc-lib"

	"github.com/golang-collections/collections/queue"
)

type BlockSupportGraph struct {
	edges map[string]aoc.Set[string]
	nodes map[string]Rect3
}

func FallingBricksOnFall(graph BlockSupportGraph, node string) int {
	q := queue.New()
	q.Enqueue(node)

	reversed := graph.ReversedEdges()
	supports := make(map[string]int)
	for n, edges := range graph.edges {
		supports[n] = edges.Len()
	}

	// Start at -1 to exclude the brick that's being disintegrated
	totalFalling := -1

	for q.Len() != 0 {
		totalFalling += 1
		node := q.Dequeue().(string)
		for edge := range reversed.edges[node].Elements() {
			supports[edge]--
			if supports[edge] == 0 {
				q.Enqueue(edge) // no supports left - it needs to fall too
			}
		}
	}
	return totalFalling
}

func FallingBricksOnFallHelper(graph BlockSupportGraph, inverted BlockSupportGraph, node string) int {
	totalFalling := 0
	supportedByMe := inverted.edges[node]
	for supportedNode := range supportedByMe.Elements() {
		if graph.edges[supportedNode].Len() == 1 {
			// If solely supported by this brick, it falls and cause any others to fall too
			totalFalling += 1 + FallingBricksOnFallHelper(graph, inverted, supportedNode)
		}
	}
	return totalFalling
}

func (graph BlockSupportGraph) FindLoadBearingBlocks() aoc.Set[string] {
	loadBearingBlocks := aoc.NewSet[string]()
	for _, outgoingEdges := range graph.edges {
		if outgoingEdges.Len() == 1 {
			loadBearingBlocks.Insert(outgoingEdges.TakeOne())
		}
	}
	return loadBearingBlocks
}

func (graph BlockSupportGraph) Fallen() BlockSupportGraph {
	fallenNodes := make(map[string]Rect3)
	evalOrder := graph.ReversedEdges().InOrderTraversal()
	for _, node := range evalOrder {
		nodesBelow := graph.edges[node]
		maxZExtent := 1 // Ground is at level 0 (1 exclusive)
		for nodeBelow := range nodesBelow.Elements() {
			maxZExtent = max(maxZExtent, fallenNodes[nodeBelow].z.Max)
		}
		unFallenBlock := graph.nodes[node]
		fallenNodes[node] = Rect3{
			x: unFallenBlock.x,
			y: unFallenBlock.y,
			z: unFallenBlock.z.Plus(-(unFallenBlock.z.Min - maxZExtent)),
		}
	}
	graph.nodes = fallenNodes
	return graph
}

func (graph BlockSupportGraph) PruneEdgesForSeparatedBlocks() BlockSupportGraph {
	newEdges := make(map[string]aoc.Set[string])
	for node, outgoingEdges := range graph.edges {
		r1 := graph.nodes[node]
		newOutgoingEdges := outgoingEdges.Clone()
		for edge := range outgoingEdges.Elements() {
			r2 := graph.nodes[edge]
			if !AreBlocksVerticallyAdjacent(r1, r2) {
				newOutgoingEdges.Delete(edge)
			}
		}
		newEdges[node] = newOutgoingEdges
	}
	graph.edges = newEdges
	return graph
}

func AreBlocksVerticallyAdjacent(b1, b2 Rect3) bool {
	if b1.z.Max > b2.z.Max {
		return b1.z.Min == b2.z.Max
	} else {
		return b2.z.Min == b1.z.Max
	}
}

func (graph BlockSupportGraph) InOrderTraversal() []string {
	q := queue.New()
	traversal := make([]string, 0)

	// Start at all nodes that have no incoming edges
	reversed := graph.ReversedEdges()
	indegrees := make(map[string]int)
	for node, outgoingEdges := range reversed.edges {
		indegrees[node] = outgoingEdges.Len()
		if outgoingEdges.Len() == 0 {
			q.Enqueue(node)
		}
	}
	for q.Len() != 0 {
		node := q.Dequeue().(string)
		traversal = append(traversal, node)
		for edge := range graph.edges[node].Elements() {
			newDegree := indegrees[edge] - 1
			indegrees[edge] = newDegree
			if newDegree == 0 {
				q.Enqueue(edge)
			}
		}
	}
	return traversal
}

func (graph BlockSupportGraph) ReversedEdges() BlockSupportGraph {
	newEdges := make(map[string]aoc.Set[string])
	for node := range graph.nodes {
		newEdges[node] = aoc.NewSet[string]()
	}
	for node, outgoingEdges := range graph.edges {
		for outgoingEdge := range outgoingEdges.Elements() {
			newEdges[outgoingEdge].Insert(node)
		}
	}
	graph.edges = newEdges
	return graph
}
