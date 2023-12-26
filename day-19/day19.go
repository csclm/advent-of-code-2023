package day19

import (
	"fmt"
	"os"
)

func Part1(f *os.File) {
	rules, parts := parseInput(f)

	ratingsSum := 0
	for _, part := range parts {
		accepted := evaluateWorkflowChain(rules, part)
		if accepted {
			ratingsSum += part.x
			ratingsSum += part.m
			ratingsSum += part.a
			ratingsSum += part.s
		}
	}

	fmt.Printf("Sum of ratings %d\n", ratingsSum)
}
