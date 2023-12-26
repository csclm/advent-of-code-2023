package day17

import (
	"aoc-2023/aoc-lib"
	"container/heap"

	"github.com/golang-collections/collections/set"
)

type SearchGraph[TNode comparable] interface {
	neighbors(node TNode) []SearchEdge[TNode]
}

type SearchEdge[TNode comparable] struct {
	node TNode
	cost int
}

// Implementation of Uniform Cost Search
// Returns the minimum cost to go from the start node to an end node
func graphSearch[TNode comparable](graph SearchGraph[TNode], start TNode, isEndState func(TNode) bool) int {
	node := SearchEdge[TNode]{node: start, cost: 0}
	frontier := aoc.SliceHeap[SearchEdge[TNode]]{
		Contents: make([]SearchEdge[TNode], 0),
		ElementLess: func(a, b SearchEdge[TNode]) bool {
			return a.cost < b.cost
		},
	}
	heap.Init(&frontier)
	heap.Push(&frontier, node)
	visited := set.New()
	for {
		if frontier.Len() == 0 {
			return -1
		}
		node = heap.Pop(&frontier).(SearchEdge[TNode])
		if isEndState(node.node) {
			return node.cost
		}
		visited.Insert(node.node)
		for _, neighbor := range graph.neighbors(node.node) {
			foundInFrontier := SearchEdge[TNode]{}
			foundInFrontierIndex := -1
			costForThisNeighbor := neighbor.cost + node.cost
			for i, frontierEdge := range frontier.Contents {
				if frontierEdge.node == neighbor.node {
					foundInFrontier = frontierEdge
					foundInFrontierIndex = i
					break
				}
			}
			if !visited.Has(neighbor.node) && foundInFrontierIndex == -1 {
				heap.Push(
					&frontier,
					SearchEdge[TNode]{
						node: neighbor.node,
						cost: neighbor.cost + node.cost,
					},
				)
			} else if foundInFrontierIndex != -1 && costForThisNeighbor < foundInFrontier.cost {
				frontier.Contents[foundInFrontierIndex].cost = costForThisNeighbor
				heap.Fix(&frontier, foundInFrontierIndex)
			}
		}
	}
}

// Iterating over all the possible edges wil be impractical for this problem
// Wikipedia mentions a Uniform Cost Search which can handle potentially infinite graphs
// and discovers edges gradually
/*
procedure uniform_cost_search(start) is
    node ← start
    frontier ← priority queue containing node only
    expanded ← empty set
    do
        if frontier is empty then
            return failure
        node ← frontier.pop()
        if node is a goal state then
            return solution(node)
        expanded.add(node)
        for each of node's neighbors n do
            if n is not in expanded and not in frontier then
                frontier.add(n)
            else if n is in frontier with higher cost
                replace existing node with n
*/

// First thing that comes to mind is dijkstra's algo. Pseudocode from wikipedia:
/*
 1  function Dijkstra(Graph, source):
 2
 3      for each vertex v in Graph.Vertices:
 4          dist[v] ← INFINITY
 5          prev[v] ← UNDEFINED
 6          add v to Q
 7      dist[source] ← 0
 8
 9      while Q is not empty:
10          u ← vertex in Q with min dist[u]
11          remove u from Q
12
13          for each neighbor v of u still in Q:
14              alt ← dist[u] + Graph.Edges(u, v)
15              if alt < dist[v]:
16                  dist[v] ← alt
17                  prev[v] ← u
18
19      return dist[], prev[]
*/
