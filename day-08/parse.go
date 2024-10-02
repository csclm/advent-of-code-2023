package day8

import (
	"aoc-2023/aoc-lib"
	"os"
	"regexp"
)

// network, string of L/R directions
func parseInput(f *os.File) (Network, string) {

	reader := aoc.LineReader(f)
	instructions := <-reader
	<-reader // blank line

	nodeNamePattern := regexp.MustCompile(`[A-Z]{3}`)

	network := Network{
		nodes: make(map[string]NodeLinks),
	}

	for nodeLine := range reader {
		nodeNames := nodeNamePattern.FindAllString(nodeLine, -1)
		network.nodes[nodeNames[0]] = NodeLinks{nodeNames[1], nodeNames[2]}
	}

	return network, instructions
}
