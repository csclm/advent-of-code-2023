package aoc

import (
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

func ReadGrid(r io.Reader) (Grid[rune], error) {
	width := 0
	contents := make([]rune, 0)
	for line := range LineReader(r) {
		if len(line) == 0 {
			break
		}
		if width == 0 {
			width = len(line)
		} else if len(line) != width {
			return *new(Grid[rune]), errors.New("input is not rectangular")
		}
		contents = append(contents, []rune(line)...)
	}
	if width == 0 {
		return *new(Grid[rune]), errors.New("no grid to read")
	}
	return NewGridWithContents(width, len(contents)/width, contents), nil
}

func LineReader(r io.Reader) chan string {
	c := make(chan string)
	go func() {
		for line := range iochan.DelimReader(r, '\n') {
			c <- strings.TrimSpace(line)
		}
		close(c)
	}()
	return c
}

func MustParseInt(str string) int {
	num, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		panic("Could not parse int " + err.Error())
	}
	return int(num)
}

func MustParseDigit(char rune) int {
	return MustParseInt(string(char))
}

func MustParseListOfNums(str string, sep string) []int {
	components := strings.Split(str, sep)
	result := make([]int, len(components))
	for i := range components {
		result[i] = MustParseInt(components[i])
	}
	return result
}
