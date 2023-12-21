package main

import "slices"

// func spinABillionTimes(grid Grid) Grid {
// 	hashSet := make(map[int][]Grid, 0)

// 	for {
// 		slideStonesNorth(grid)
// 		grid = grid.RotatedRight()
// 	}
// }

func gridDeepCopy(grid Grid) Grid {
	newGridContents := slices.Clone(grid.contents)
	for i, s := range newGridContents {
		newGridContents[i] = slices.Clone(s)
	}
	return Grid{newGridContents, grid.orientation}
}

func hashRoundStonePositions(gridContents [][]rune) int {
	var result int32 = 0
	for ri := range gridContents {
		for ci := range gridContents[ri] {
			if gridContents[ri][ci] == RoundStone {
				result += int32(ri*(1<<8) + ci*(1<<16))
				result = result<<4 | result>>28 // Rotate left 4 bits
			}
		}
	}
	return int(result)
}
