package day2

import (
	"fmt"
	"os"

	"github.com/mitchellh/iochan"
)

type DrawBag struct {
	reds   int
	greens int
	blues  int
}

type DrawGame struct {
	id    int
	draws []Draw
}

type Draw struct {
	quantity int
	color    string
}

func Part1(f *os.File) {
	sumOfIds := 0
	part1Bag := DrawBag{
		reds:   12,
		greens: 13,
		blues:  14,
	}
	for line := range iochan.DelimReader(f, '\n') {
		game := parseGame(line)
		if game.isPossibleWithBag(part1Bag) {
			sumOfIds += game.id
		}
	}
	fmt.Printf("Sum of possible IDs is: %d\n", sumOfIds)
}

func Part2(f *os.File) {
	sumOfPowers := 0
	for line := range iochan.DelimReader(f, '\n') {
		game := parseGame(line)
		sumOfPowers += game.power()
	}
	fmt.Printf("Sum of all powers is: %d\n", sumOfPowers)
}

func (game DrawGame) isPossibleWithBag(bag DrawBag) bool {
	minRed, minGreen, minBlue := game.minimumCubesRequired()
	return minRed <= bag.reds && minGreen <= bag.greens && minBlue <= bag.blues
}

func (game DrawGame) power() int {
	minRed, minGreen, minBlue := game.minimumCubesRequired()
	return minRed * minGreen * minBlue
}

// red,green,blue
func (game DrawGame) minimumCubesRequired() (int, int, int) {
	minRed := 0
	minGreen := 0
	minBlue := 0
	for _, draw := range game.draws {
		switch draw.color {
		case "red":
			minRed = max(minRed, draw.quantity)
		case "green":
			minGreen = max(minGreen, draw.quantity)
		case "blue":
			minBlue = max(minBlue, draw.quantity)
		default:
			panic("shouldn't happen")
		}
	}

	return minRed, minGreen, minBlue
}
