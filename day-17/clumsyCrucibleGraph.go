package day17

import "aoc-2023/aoc-lib"

type ClumsyCrucibleGraph struct {
	contents        [][]int
	maxStraightLine int // minimum number of squares before a crucible can turn
	minStraightLine int // maximum number of squares before a crucible can turn
}

func (g ClumsyCrucibleGraph) Width() int {
	return len(g.contents[0])
}

func (g ClumsyCrucibleGraph) Height() int {
	return len(g.contents)
}

type ClumsyCrucibleNode struct {
	location      aoc.Vec2
	lastDirection aoc.Vec2
	momentum      int
}

func (g ClumsyCrucibleGraph) neighbors(node ClumsyCrucibleNode) []SearchEdge[ClumsyCrucibleNode] {
	result := make([]SearchEdge[ClumsyCrucibleNode], 0)
	for _, dir := range aoc.CardinalDirections() {
		newLocation := node.location.Plus(dir)
		if !newLocation.IsInBoundingBox(g.Width(), g.Height()) {
			// Can't go out of bounds
			continue
		}
		if dir == node.lastDirection.Inverse() {
			// Can't backtrack
			continue
		}
		var newMomentum int
		if dir == node.lastDirection {
			newMomentum = node.momentum + 1
		} else {
			if node.momentum < g.minStraightLine && node.lastDirection != (aoc.Vec2{0, 0}) {
				// Can't turn until we've moved far enough
				continue
			}
			newMomentum = 1
		}
		if newMomentum > g.maxStraightLine {
			// Can't move in the same direction more than 3 squares
			continue
		}
		newNode := ClumsyCrucibleNode{
			location:      newLocation,
			lastDirection: dir,
			momentum:      newMomentum,
		}
		newCost := g.contents[newLocation.Y][newLocation.X]
		result = append(result, SearchEdge[ClumsyCrucibleNode]{newNode, newCost})
	}
	return result
}
