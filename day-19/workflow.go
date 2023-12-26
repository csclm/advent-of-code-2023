package day19

type Part struct {
	x, m, a, s int
}

func (p Part) GetProperty(name rune) int {
	switch name {
	case 'x':
		return p.x
	case 'm':
		return p.m
	case 'a':
		return p.a
	case 's':
		return p.s
	default:
		panic("Invalid property name for part")
	}
}

type Workflow struct {
	rules         []WorkflowRule
	defaultAction WorkflowAction
}

type WorkflowRule struct {
	condition WorkflowCondition
	action    WorkflowAction
}

type WorkflowAction struct {
	// accept = true OR reject = true OR next = name of next rule
	accept bool
	reject bool
	next   string
}

type WorkflowCondition struct {
	property rune // x, m, a, or s
	operator rune // '>', '<', '=', or '!' (not equal)
	value    int
}

func (wc WorkflowCondition) Inverted() WorkflowCondition {
	switch wc.operator {
	case '>':
		wc.operator = '<'
	case '<':
		wc.operator = '>'
	case '=':
		wc.operator = '!'
	case '!':
		wc.operator = '='
	default:
		panic("Invalid operator")
	}
	return wc
}

func (w Workflow) evaluate(part Part) WorkflowAction {
	for _, rule := range w.rules {
		if rule.condition.evaluate(part) {
			return rule.action
		}
	}
	return w.defaultAction
}

func (wc WorkflowCondition) evaluate(part Part) bool {
	valueToTest := part.GetProperty(wc.property)
	switch wc.operator {
	case '=':
		return valueToTest == wc.value
	case '!':
		return valueToTest != wc.value
	case '>':
		return valueToTest > wc.value
	case '<':
		return valueToTest < wc.value
	default:
		panic("Invalid operator for workflow condition")
	}
}

// True for accept, false for reject
func evaluateWorkflowChain(workflows map[string]Workflow, part Part) bool {
	workflow := workflows["in"]
	for {
		action := workflow.evaluate(part)
		if action.accept {
			return true
		} else if action.reject {
			return false
		} else {
			workflow = workflows[action.next]
		}
	}
}
