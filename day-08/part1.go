package day8

func numberOfTraversalStepsPart1(directions string, network Network) int {
	currentNode := "AAA"
	step := 0
	for {
		if currentNode == "ZZZ" {
			return step
		}
		if directions[step%len(directions)] == 'L' {
			currentNode = network.nodes[currentNode].left
		} else {
			currentNode = network.nodes[currentNode].right
		}
		step++
	}
}
