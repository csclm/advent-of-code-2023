package day25

import (
	"math"
)

const INF = math.MaxInt

type StoerWagnerGraph struct {
	V        int
	adj      [][]int
	inMinCut []bool
}

func newGraph(V int) *StoerWagnerGraph {
	g := &StoerWagnerGraph{
		V:        V,
		adj:      make([][]int, V),
		inMinCut: make([]bool, V),
	}

	for i := range g.adj {
		g.adj[i] = make([]int, V)
	}

	return g
}

func (g *StoerWagnerGraph) addEdge(u, v, w int) {
	g.adj[u][v] = w
	g.adj[v][u] = w
}

func (g *StoerWagnerGraph) minCutPhase() (int, []int) {
	merged := make([]bool, g.V)
	included := make([]bool, g.V)
	dist := make([]int, g.V)
	best := -1
	var path []int

	for phase := g.V - 1; phase > 0; phase-- {
		for i := range dist {
			dist[i] = 0
			merged[i] = false
			included[i] = false
		}

		u := 0
		for i := 0; i < phase; i++ {
			last := u
			u = -1

			for v := 0; v < g.V; v++ {
				if !merged[v] && !included[v] {
					if u == -1 || dist[v] > dist[u] {
						u = v
					}
				}
			}

			if i == phase-1 {
				if dist[u] > dist[last] {
					best = last
					path = append(path, u)
				} else {
					best = u
					path = append([]int{last}, path...)
				}
				break
			}

			included[u] = true

			for v := 0; v < g.V; v++ {
				if !merged[v] && !included[v] {
					dist[v] += g.adj[u][v]
				}
			}
		}

		if phase > 1 {
			for v := 0; v < g.V; v++ {
				if !merged[v] {
					g.adj[best][v] += g.adj[u][v]
					g.adj[v][best] += g.adj[u][v]
				}
			}
			merged[u] = true
		}
	}

	return best, path
}

func (g *StoerWagnerGraph) minCut() int {
	minCut := INF

	for i := 0; i < g.V-1; i++ {
		u, path := g.minCutPhase()
		minCut = int(math.Min(float64(minCut), float64(g.adj[u][path[0]])))
		g.inMinCut[u] = true

		for j := 0; j < len(path)-1; j++ {
			g.adj[path[j]][path[j+1]] = g.adj[path[j+1]][path[j]]
			g.adj[path[j+1]][path[j]] = 0
		}
	}

	return minCut
}
