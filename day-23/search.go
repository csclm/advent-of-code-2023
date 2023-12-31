package day23

import (
	"aoc-2023/aoc-lib"
)

type SearchGraph[T comparable] interface {
	vertices() []T
	neighbors(T) []SearchEdge[T]
}

type SearchEdge[T comparable] struct {
	to   T
	cost int
}

type SearchPath[T comparable] struct {
	distances map[T]int
	links     map[T](*T)
}

func dijkstra[T comparable](graph SearchGraph[T], start T) SearchPath[T] {
	q := aoc.NewSet[T]()
	dist := make(map[T]int)
	prev := make(map[T](*T)) // nullable pointers to the previous node
	for _, v := range graph.vertices() {
		dist[v] = (1 << 62)
		prev[v] = nil
		q.Insert(v)
	}
	q.Insert(start)
	dist[start] = 0

	for q.Len() != 0 {
		minDist := (1 << 62)
		u := *new(T)
		for node := range q.Elements() {
			thisDist := dist[node]
			if thisDist < minDist {
				minDist = thisDist
				u = node
			}
		}
		if minDist == 1<<62 {
			break // search is over
		}
		q.Delete(u)
		for _, neighbor := range graph.neighbors(u) {
			v := neighbor.to
			if !q.Has(v) {
				continue
			}
			alt := dist[u] + neighbor.cost
			// Ensure this pointer chain doensn't already contain this node
			// TODO this doesn't seem to work - it still doesn't give me the longest path on the example
			check := u
			foundDuplicate := false
			for {
				if check == v {
					foundDuplicate = true
					break
				}
				ptr := prev[check]
				if ptr == nil {
					break
				}
				check = *ptr
			}
			if alt < dist[v] && !foundDuplicate {
				dist[v] = alt
				prev[v] = &u
			}
		}
	}
	return SearchPath[T]{
		distances: dist,
		links:     prev,
	}
}

/* From Wikipedia:

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
