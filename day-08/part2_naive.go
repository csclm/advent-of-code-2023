package day8

// Naive solution, takes a long time to return an answer
func numberOfTraversalStepsPart2Naive(directions string, network Network) int {
	currentNodes := make([]string, 0)
	for k := range network.nodes {
		if k[2] == 'A' {
			currentNodes = append(currentNodes, k)
		}
	}
	step := 0
	for {
		anyNodeNotAtFinish := false
		for _, node := range currentNodes {
			if node[2] != 'Z' {
				anyNodeNotAtFinish = true
				break
			}
		}
		if !anyNodeNotAtFinish {
			return step
		}
		if directions[step%len(directions)] == 'L' {
			for i, node := range currentNodes {
				currentNodes[i] = network.nodes[node].left
			}
		} else {
			for i, node := range currentNodes {
				currentNodes[i] = network.nodes[node].right
			}
		}
		step++
	}
}
