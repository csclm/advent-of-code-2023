package day8

import (
	"os"
	"regexp"
	"strings"

	"github.com/mitchellh/iochan"
)

// network, string of L/R directions
func parseInput(f *os.File) (Network, string) {

	reader := iochan.DelimReader(f, '\n')
	instructions := strings.TrimSpace(<-reader)
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
