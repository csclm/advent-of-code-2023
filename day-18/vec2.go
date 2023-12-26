package day18

import "aoc-2023/aoc-lib"

func perimiterCoords(width, height int) []aoc.Vec2 {
	result := make([]aoc.Vec2, 0)
	for i := 0; i < height; i++ {
		result = append(result, aoc.NewVec2(0, i))
		result = append(result, aoc.NewVec2(width-1, i))
	}
	for i := 1; i < width-1; i++ {
		result = append(result, aoc.NewVec2(i, 0))
		result = append(result, aoc.NewVec2(i, height-1))
	}
	return result
}
