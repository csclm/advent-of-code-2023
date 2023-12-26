package day13

func findReflectionColumnsWithSmudge(grid Grid) []int {
	result := make([]int, 0)
	for col := 0; col < (grid.Width() - 1); col++ {
		if isSymmetricalAboutColumnWithSmudge(grid, col) {
			result = append(result, col)
		}
	}
	return result
}

// Defining line N as the line between the N and N+1 indices

func isSymmetricalAboutColumnWithSmudge(grid Grid, column int) bool {
	foundSmudge := false
	for row := 0; row < grid.Height(); row++ {
		for i := 0; i <= column; i++ {
			if grid.Width() < column+i+2 {
				break
			}
			if grid.RuneAt(row, column-i) != grid.RuneAt(row, column+i+1) {
				if !foundSmudge {
					foundSmudge = true
				} else {
					return false
				}
			}
		}
	}
	return foundSmudge
}
