package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func main() {
	f, _ := os.Open("./input.txt")
	inputs := parseInput(f)

	totalHash := 0
	for _, input := range inputs {
		totalHash += HolidayHash(input)
	}

	hashMap := CreateHolidayHashmap(inputs)

	fmt.Printf("Total of hashes %d\n", totalHash)
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

func CreateHolidayHashmap(lenses []string) HolidayHashmap {
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
