package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func parseInput(f *os.File) (map[string]Workflow, []Part) {
	workflowMap := make(map[string]Workflow, 0)
	pattern := regexp.MustCompile(`(\w+)\{((?:[xmas][><=]\d+:\w+,)+(?:\w+))\}`)

	lineReader := iochan.DelimReader(f, '\n')

	for line := range lineReader {
		if len(strings.TrimSpace(line)) == 0 {
			break
		}
		matches := pattern.FindStringSubmatch(strings.TrimSpace(line))
		workflowName := matches[1]
		rules := make([]WorkflowRule, 0)
		workflowComponents := strings.Split(matches[2], ",")
		for _, ruleString := range workflowComponents[:len(workflowComponents)-1] {
			rules = append(rules, parseRule(ruleString))
		}
		defaultAction := parseAction(workflowComponents[len(workflowComponents)-1])
		workflowMap[workflowName] = Workflow{rules: rules, defaultAction: defaultAction}
	}

	parts := make([]Part, 0)
	for line := range lineReader {
		parts = append(parts, parsePart(line))
	}

	return workflowMap, parts
}

func parseRule(s string) WorkflowRule {
	components := strings.Split(s, ":")
	condition := parseCondition(components[0])
	action := parseAction(components[1])
	return WorkflowRule{
		condition: condition,
		action:    action,
	}
}

func parseAction(s string) WorkflowAction {
	switch s {
	case "A":
		return WorkflowAction{accept: true}
	case "R":
		return WorkflowAction{reject: true}
	default:
		return WorkflowAction{next: s}
	}
}

func parseCondition(s string) WorkflowCondition {
	valNum, _ := strconv.ParseInt(s[2:], 10, 0)
	return WorkflowCondition{
		property: []rune(s)[0],
		operator: []rune(s)[1],
		value:    int(valNum),
	}
}

func parsePart(str string) Part {
	pattern := regexp.MustCompile(`\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}`)
	matches := pattern.FindStringSubmatch(str)
	x, _ := strconv.ParseInt(matches[1], 10, 0)
	m, _ := strconv.ParseInt(matches[2], 10, 0)
	a, _ := strconv.ParseInt(matches[3], 10, 0)
	s, _ := strconv.ParseInt(matches[4], 10, 0)
	return Part{
		x: int(x),
		m: int(m),
		a: int(a),
		s: int(s),
	}
}
