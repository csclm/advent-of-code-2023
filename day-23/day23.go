package day23

import (
	"fmt"
	"os"
)

func Part1(f *os.File) {
	trails := parseInput(f)
	length := trails.findLongestWalkLength()
	fmt.Printf("Longest walk length is %d\n", length)
}

func Part2(f *os.File) {

}
