package day22

import (
	"aoc-2023/aoc-lib"
	"fmt"
	"os"
	"strconv"
)

func Part1(f *os.File) {
	rects := parseInput(f)
	graph := makeDependencyGraph(rects)
	graph = graph.Fallen()
	checkForIntersections(graph.nodes) // Testing assertion
	graph = graph.PruneEdgesForSeparatedBlocks()
	loadBearingBlocks := graph.FindLoadBearingBlocks()
	fmt.Printf("%d blocks can safely be disintegrated\n", len(rects)-loadBearingBlocks.Len())
}

func Part2(f *os.File) {
	rects := parseInput(f)
	graph := makeDependencyGraph(rects)
	graph = graph.Fallen()
	checkForIntersections(graph.nodes) // Testing assertion
	graph = graph.PruneEdgesForSeparatedBlocks()
	totalFalling := 0
	for node := range graph.nodes {
		totalFalling += FallingBricksOnFall(graph, node)
	}
	fmt.Printf("Total number of falling blocks: %d", totalFalling)
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func checkForIntersections(rects map[string]Rect3) {
	for i1 := 0; i1 < len(rects); i1++ {
		for i2 := 0; i2 < i1; i2++ {
			r1 := rects[itoa(i1)]
			r2 := rects[itoa(i2)]
			if r1.x.Intersects(r2.x) && r1.y.Intersects(r2.y) && r1.z.Intersects(r2.z) {
				panic("blocks are intersecting!")
			}
		}
	}
}

func makeDependencyGraph(rects []Rect3) BlockSupportGraph {
	nodes := make(map[string]Rect3)
	edges := make(map[string](aoc.Set[string]))
	for i, rect := range rects {
		nodeName := itoa(i)
		nodes[nodeName] = rect
		edges[nodeName] = aoc.NewSet[string]()
	}
	for i1, r1 := range rects {
		for i2 := 0; i2 < i1; i2++ {
			r2 := rects[i2]
			if !r1.x.Intersects(r2.x) || !r1.y.Intersects(r2.y) {
				// 2 rectangles can't ever stack on each other
				continue
			}
			// Assumes none of the rectangles fully intersect each other
			if r1.z.Max > r2.z.Max {
				edges[itoa(i1)].Insert(itoa(i2))
			} else {
				edges[itoa(i2)].Insert(itoa(i1))
			}
		}
	}
	return BlockSupportGraph{
		nodes: nodes,
		edges: edges,
	}
}

type Rect3 struct {
	x, y, z aoc.Range
}
