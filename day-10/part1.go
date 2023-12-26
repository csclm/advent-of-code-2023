package day10

import "aoc-2023/aoc-lib"

func CalculateFurthestStep(grid PipeGrid) int {
	loopLength := 0
	for range LoopSteps(&grid) {
		loopLength++
	}
	return loopLength / 2
}

func LoopSteps(grid *PipeGrid) chan aoc.Vec2 {
	result := make(chan aoc.Vec2)
	go func() {
		start := grid.MustFindStartingPoint()
		previousLocation := start
		currentLocation := start
		firstStep := true
		for firstStep || currentLocation != start {
			newLocation := MustStepAlongPath(grid, currentLocation, previousLocation)
			previousLocation = currentLocation
			currentLocation = newLocation
			result <- currentLocation
			firstStep = false
		}
		close(result)
	}()
	return result
}

func MustStepAlongPath(grid *PipeGrid, location aoc.Vec2, previousLocation aoc.Vec2) aoc.Vec2 {
	runeAtLocation, inBounds := grid.RuneAtLocation(location)
	if !inBounds {
		panic("stepped out of bounds!")
	}
	for _, dir := range aoc.CardinalDirections() {
		spotInDirection := location.Plus(dir)
		if spotInDirection == previousLocation {
			continue // don't step backwards
		}
		if !runeAtLocation.ConnectsTo(dir) {
			continue // doesn't connect
		}
		runeInDirection, directionInBounds := grid.RuneAtLocation(spotInDirection)
		if !directionInBounds {
			continue // would step out of bounds
		}
		if !runeInDirection.ConnectsTo(dir.Inverse()) {
			continue // next rune doesn't connect back to this one
		}
		return spotInDirection
	}
	panic("Hit dead end stepping along path")
}
