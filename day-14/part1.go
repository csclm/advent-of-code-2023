package main

func slideStonesNorth(grid Grid) {
	for col := 0; col < grid.Width(); col++ {
		northShiftStonesInGridColumn(grid, col)
	}
}

// TODO this is borked now. Why? I thought slices were opaque pointers?

// Shifts stones toward the zero index
func northShiftStonesInGridColumn(grid Grid, column int) {
	roundStonesInSegment := 0
	segmentLength := 0
	assembled := 0

	flushSegment := func() {
		for i := 0; i < segmentLength; i++ {
			if i < roundStonesInSegment {
				grid.SetRuneAt(assembled+i, column, RoundStone)
			} else {
				grid.SetRuneAt(assembled+i, column, EmptySpace)
			}
		}
		assembled += segmentLength
		segmentLength = 0
		roundStonesInSegment = 0
	}

	for cell := 0; cell < grid.Height(); cell++ {
		if cell == SquareStone {
			flushSegment()
			assembled++ // Skip over this stone
		} else if cell == RoundStone {
			segmentLength++
			roundStonesInSegment++
		} else {
			segmentLength++
		}
	}
	flushSegment()
}
