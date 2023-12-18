package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

const DamagedSpring = '#'
const OperationalSpring = '.'
const UnknownSpring = '?'

type SpringRecord struct {
	damagedSpringGroups []int
	springs             []rune
}

func (sr SpringRecord) Unfolded() SpringRecord {
	springs := make([]rune, len(sr.springs)*5+4)
	for i := 0; i < len(springs); i++ {
		mod := i % (len(sr.springs) + 1)
		if mod == len(sr.springs) {
			springs[i] = UnknownSpring
		} else {
			springs[i] = sr.springs[mod]
		}
	}
	groups := make([]int, len(sr.damagedSpringGroups)*5)
	for i := range groups {
		groups[i] = sr.damagedSpringGroups[i%len(sr.damagedSpringGroups)]
	}
	return SpringRecord{
		damagedSpringGroups: groups,
		springs:             springs,
	}
}

func main() {
	f, _ := os.Open("./input.txt")
	records := parseInput(f)
	fmt.Printf("%d Records\n", len(records))
	sumOfPossibilities := 0
	for i, record := range records {
		solutions := possibleSolutions(record)
		sumOfPossibilities += solutions
		fmt.Printf("Record %d has %d possibilities\n", i, solutions)
	}
	fmt.Printf("Sum of possibilities: %d\n", sumOfPossibilities)

	sumOfUnfoldedPossibilities := 0
	for i, record := range records {
		solutions := possibleSolutions(record.Unfolded())
		sumOfUnfoldedPossibilities += solutions
		fmt.Printf("Unfolded record %d has %d possibilities\n", i, solutions)
	}
	fmt.Printf("Sum of unfolded possibilities: %d\n", sumOfUnfoldedPossibilities)
}

func parseInput(f *os.File) []SpringRecord {
	result := make([]SpringRecord, 0)
	for line := range iochan.DelimReader(f, '\n') {
		components := strings.Split(strings.TrimSpace(line), " ")
		groupsStrings := strings.Split(components[1], ",")
		groupsSlice := make([]int, len(groupsStrings))
		for i, g := range groupsStrings {
			groupNum, _ := strconv.ParseInt(g, 10, 0)
			groupsSlice[i] = int(groupNum)
		}
		result = append(result, SpringRecord{
			damagedSpringGroups: groupsSlice,
			springs:             []rune(components[0]),
		})
	}
	return result
}

type SpringRecordPiece struct {
	springStart int // records.springs[springStart:]
	groupStart  int // records.damagedSpringGroups[groupStart:]
}

func possibleSolutions(record SpringRecord) int {
	memo := make(map[SpringRecordPiece]int)
	return possibleSolutionsWithMemo(record, SpringRecordPiece{0, 0}, memo)
}

func possibleSolutionsWithMemo(baseRecord SpringRecord, piece SpringRecordPiece, memo map[SpringRecordPiece]int) int {
	memoResult, hasMemo := memo[piece]
	if hasMemo {
		return memoResult
	}

	record := SpringRecord{
		damagedSpringGroups: baseRecord.damagedSpringGroups[piece.groupStart:],
		springs:             baseRecord.springs[piece.springStart:],
	}

	// base case
	if len(record.springs) == 0 {
		if len(record.damagedSpringGroups) == 0 {
			return 1
		} else {
			return 0
		}
	} else if len(record.damagedSpringGroups) > 0 && len(record.springs) < record.damagedSpringGroups[0] {
		// No way to satisfy the remaining groups
		return 0
	}

	possibilitiesIfOperational := func() int {
		return possibleSolutionsWithMemo(baseRecord, SpringRecordPiece{
			springStart: piece.springStart + 1,
			groupStart:  piece.groupStart,
		}, memo)
	}

	possibilitiesIfDamaged := func() int {
		if len(record.damagedSpringGroups) == 0 {
			return 0 // couldn't be damaged - no groups left
		}
		nextGroupWidth := record.damagedSpringGroups[0]
		// guaranteed that record.springs is long enough because of the check in the base case
		for i := 1; i < nextGroupWidth; i++ {
			if record.springs[i] == OperationalSpring {
				// No way this spring can be broken - an operational spring is too close
				return 0
			}
		}
		if len(record.springs) > nextGroupWidth {
			if record.springs[nextGroupWidth] == DamagedSpring {
				// This group is too wide
				return 0
			}
			// Otherwise, the following spring must be operational
			return possibleSolutionsWithMemo(baseRecord, SpringRecordPiece{
				springStart: piece.springStart + nextGroupWidth + 1,
				groupStart:  piece.groupStart + 1,
			}, memo)
		} else {
			// No spring after this list to check, hit the base case to find out if there are remaining groups
			return possibleSolutionsWithMemo(baseRecord, SpringRecordPiece{
				springStart: piece.springStart + nextGroupWidth,
				groupStart:  piece.groupStart + 1,
			}, memo)
		}
	}

	switch record.springs[0] {
	case OperationalSpring:
		result := possibilitiesIfOperational()
		memo[piece] = result
		return result
	case DamagedSpring:
		result := possibilitiesIfDamaged()
		memo[piece] = result
		return result
	case UnknownSpring:
		result := possibilitiesIfOperational() + possibilitiesIfDamaged()
		memo[piece] = result
		return result
	default:
		panic("Unknown spring type " + string(record.springs[0]))
	}
}
