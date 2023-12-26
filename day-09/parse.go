package day9

import (
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func parseInputFile(f *os.File) [][]int {
	sequences := make([][]int, 0)
	for line := range iochan.DelimReader(f, '\n') {
		numStrings := strings.Split(strings.TrimSpace(line), " ")
		nums := make([]int, 0)
		for _, numStr := range numStrings {
			num, _ := strconv.ParseInt(numStr, 10, 0)
			nums = append(nums, int(num))
		}
		sequences = append(sequences, nums)
	}
	return sequences
}
