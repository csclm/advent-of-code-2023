package main

type ClumsyCrucibleGraph struct {
	contents [][]int
}

func (g ClumsyCrucibleGraph) Width() int {
	return len(g.contents[0])
}

func (g ClumsyCrucibleGraph) Height() int {
	return len(g.contents)
}

type Vec2 struct {
	x, y int
}

func (v Vec2) Plus(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

func (v Vec2) Inverse() Vec2 {
	return Vec2{-v.x, -v.y}
}

func (v Vec2) IsInBounds(width, height int) bool {
	return v.x >= 0 && v.y >= 0 && v.x < width && v.y < height
}

type ClumsyCrucibleNode struct {
	location      Vec2
	lastDirection Vec2
	momentum      int
}

func (g ClumsyCrucibleGraph) neighbors(node ClumsyCrucibleNode) []SearchEdge[ClumsyCrucibleNode] {
	result := make([]SearchEdge[ClumsyCrucibleNode], 0)
	for _, dir := range cardinalDirections() {
		newLocation := node.location.Plus(dir)
		if !newLocation.IsInBounds(g.Width(), g.Height()) {
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
			newMomentum = 1
		}
		if newMomentum > 3 {
			// Can't move in the same direction more than 3 squares
			continue
		}
		newNode := ClumsyCrucibleNode{
			location:      newLocation,
			lastDirection: dir,
			momentum:      newMomentum,
		}
		newCost := g.contents[newLocation.y][newLocation.x]
		result = append(result, SearchEdge[ClumsyCrucibleNode]{newNode, newCost})
	}
	return result
}

func cardinalDirections() [4]Vec2 {
	return [...]Vec2{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
}
