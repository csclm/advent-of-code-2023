package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("./input.txt")
	sequences := parseInputFile(f)

	sumOfRightExtrapolated := 0
	for _, sequence := range sequences {
		sumOfRightExtrapolated += extrapolateRight(sequence)
	}

	sumOfLeftExtrapolated := 0
	for _, sequence := range sequences {
		sumOfLeftExtrapolated += extrapolateLeft(sequence)
	}

	fmt.Printf("Sum of extrapolated values to the right: %d \n", sumOfRightExtrapolated)
	fmt.Printf("Sum of extrapolated values to the left: %d \n", sumOfLeftExtrapolated)
}

func extrapolateRight(sequence []int) int {
	if isAllZeroes(sequence) {
		return 0
	}
	derivative := calcDerivative(sequence)
	return sequence[len(sequence)-1] + extrapolateRight(derivative)
}

func extrapolateLeft(sequence []int) int {
	if isAllZeroes(sequence) {
		return 0
	}
	derivative := calcDerivative(sequence)
	return sequence[0] - extrapolateLeft(derivative)
}

func calcDerivative(sequence []int) []int {
	derivative := make([]int, len(sequence)-1)
	for i := range derivative {
		derivative[i] = sequence[i+1] - sequence[i]
	}
	return derivative
}

func isAllZeroes(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}
