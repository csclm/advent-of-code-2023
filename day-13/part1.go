package day13

func findReflectionColumns(grid Grid) []int {
	result := make([]int, 0)
	for col := 0; col < (grid.Width() - 1); col++ {
		if isSymmetricalAboutColumn(grid, col) {
			result = append(result, col)
		}
	}
	return result
}

// Defining line N as the line between the N and N+1 indices

func isSymmetricalAboutColumn(grid Grid, column int) bool {
	for row := 0; row < grid.Height(); row++ {
		for i := 0; i <= column; i++ {
			if grid.Width() < column+i+2 {
				break
			}
			if grid.RuneAt(row, column-i) != grid.RuneAt(row, column+i+1) {
				return false
			}
		}
	}
	return true
}
