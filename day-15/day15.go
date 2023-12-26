package day15

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func Part1(f *os.File) {
	inputs := parseInput(f)
	totalHash := 0
	for _, input := range inputs {
		totalHash += HolidayHash(input)
	}
	fmt.Printf("Total of hashes %d\n", totalHash)
}

func Part2(f *os.File) {
	inputs := parseInput(f)
	hashMap := createHolidayHashmap(inputs)
	fmt.Printf("Total focusing power %d\n", hashMap.TotalFocusingPower())
}

func parseInput(f *os.File) []string {
	result := make([]string, 0)
	for input := range iochan.DelimReader(f, ',') {
		if input[len(input)-1] == ',' {
			result = append(result, input[:len(input)-1])
		} else {
			result = append(result, input)
		}
	}
	return result
}

func createHolidayHashmap(lenses []string) HolidayHashmap {
	holidayHashmap := HolidayHashmap{}
	for _, lens := range lenses {
		if lens[len(lens)-1] == '-' {
			label := lens[:len(lens)-1]
			holidayHashmap.Remove(label)
		} else {
			components := strings.Split(lens, "=")
			label := components[0]
			focalLengthString := components[1]
			focalLength, _ := strconv.ParseInt(focalLengthString, 10, 0)
			holidayHashmap.Insert(label, int(focalLength))
		}
	}
	return holidayHashmap
}
