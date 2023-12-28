package day19

import (
	"fmt"
	"os"
)

func Part1(f *os.File) {
	workflows, parts := parseInput(f)

	ratingsSum := 0
	for _, part := range parts {
		accepted := evaluateWorkflowChain(workflows, part)
		if accepted {
			ratingsSum += part.x
			ratingsSum += part.m
			ratingsSum += part.a
			ratingsSum += part.s
		}
	}

	fmt.Printf("Sum of ratings %d\n", ratingsSum)
}

func Part2(f *os.File) {
	workflows, _ := parseInput(f)
	accepted := findNumberOfAcceptedParts("in", workflows, NewPartGamut())
	fmt.Printf("Possible accepted parts %d\n", accepted)
}

func findNumberOfAcceptedParts(workflowName string, workflows map[string]Workflow, priorGamut PartGamut) int {
	workflow := workflows[workflowName]
	currentGamut := priorGamut
	result := 0
	for _, rule := range workflow.rules {
		gamutSatisfyingRule := currentGamut.ThatSatisfiesCondition(rule.condition)
		if rule.action.accept {
			result += gamutSatisfyingRule.Cardinality()
		} else if !rule.action.reject {
			result += findNumberOfAcceptedParts(rule.action.next, workflows, gamutSatisfyingRule)
		}
		currentGamut = currentGamut.ThatDoesNotSatisfyCondition(rule.condition)
	}
	if workflow.defaultAction.accept {
		result += currentGamut.Cardinality()
	} else if !workflow.defaultAction.reject {
		result += findNumberOfAcceptedParts(workflow.defaultAction.next, workflows, currentGamut)
	}
	return result
}
