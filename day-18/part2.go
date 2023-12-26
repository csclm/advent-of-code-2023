package day18

import (
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

// TODO

type RectangularTrough struct {
	x, y, width, height int
}

func (t RectangularTrough) Volume() int {
	return t.width * t.height
}

func parseInputWithHexInstructions(f *os.File) []DigInstruction {
	pattern := regexp.MustCompile(`([UDRL]) (\d+) \((#\w+)\)`)
	result := make([]DigInstruction, 0)
	for line := range iochan.DelimReader(f, '\n') {
		match := pattern.FindStringSubmatch(strings.TrimSpace(line))
		distanceNum, _ := strconv.ParseInt(match[3], 16, 0)
		result = append(result, DigInstruction{
			direction: []rune(match[1])[0],
			distance:  int(distanceNum),
			color:     "", // irrelevant
		})
	}
	return result
}
