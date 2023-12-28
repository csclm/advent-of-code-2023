package day14

import (
	"slices"
)

func spinABillionTimes(grid Grid) Grid {
	// 2 pointers method to find a cycle
	cycleLength := 0
	slowGrid := gridDeepCopy(grid)
	fastGrid := gridDeepCopy(grid)
	for cycleLength == 0 || !gridsAreEqual(slowGrid, fastGrid) {
		spinCycle(slowGrid)
		spinCycle(fastGrid)
		spinCycle(fastGrid)
		cycleLength++
	}

	fastGrid = gridDeepCopy(grid)
	cycleStart := 0
	for !gridsAreEqual(slowGrid, fastGrid) {
		spinCycle(slowGrid)
		spinCycle(fastGrid)
		cycleStart++
	}

	offsetFromCycleStart := (1000000000 - cycleStart) % cycleLength
	for step := 0; step < offsetFromCycleStart; step++ {
		spinCycle(slowGrid)
	}

	// Should be equivalent to a billion spins
	return slowGrid
}

func spinCycle(grid Grid) {
	orientedGrid := grid // North
	slideStonesNorth(orientedGrid)

	orientedGrid = orientedGrid.RotatedLeft() // West
	slideStonesNorth(orientedGrid)

	orientedGrid = orientedGrid.RotatedLeft() // South
	slideStonesNorth(orientedGrid)

	orientedGrid = orientedGrid.RotatedLeft() // East
	slideStonesNorth(orientedGrid)
}

func gridDeepCopy(grid Grid) Grid {
	newGridContents := make([][]rune, grid.Height())
	for i := range grid.contents {
		newGridContents[i] = slices.Clone(grid.contents[i])
	}
	return Grid{newGridContents, grid.orientation}
}

func gridsAreEqual(g1 Grid, g2 Grid) bool {
	if g1.Width() != g2.Width() || g1.Height() != g2.Height() {
		return false
	}
	for ri := 0; ri < g1.Height(); ri++ {
		for ci := 0; ci < g1.Width(); ci++ {
			if g1.RuneAt(ri, ci) != g2.RuneAt(ri, ci) {
				return false
			}
		}
	}
	return true
}
