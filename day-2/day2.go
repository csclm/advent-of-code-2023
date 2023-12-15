package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/mitchellh/iochan"
)

type drawBag struct {
	reds   int
	greens int
	blues  int
}

type drawGame struct {
	id    int
	draws []draw
}

type draw struct {
	quantity int
	color    string
}

func main() {
	f, _ := os.Open("../puzzle-inputs/day-2.txt")
	sumOfIds := 0
	sumOfPowers := 0
	part1Bag := drawBag{
		reds:   12,
		greens: 13,
		blues:  14,
	}
	for line := range iochan.DelimReader(f, '\n') {
		game := parseGame(line)
		if game.isPossibleWithBag(part1Bag) {
			sumOfIds += game.id
		}
		sumOfPowers += game.power()
	}
	fmt.Printf("Sum of possible IDs is: %d\n", sumOfIds)
	fmt.Printf("Sum of all powers is: %d\n", sumOfPowers)
}

func parseGame(game string) drawGame {
	pattern := regexp.MustCompile("[:,;]\\s?")
	components := pattern.Split(game, -1)
	gameId, _ := strconv.ParseInt(strings.Split(components[0], " ")[1], 10, 0)
	parsedDraws := make([]draw, max(0, len(components)-1))
	for i, drawString := range components {
		if i == 0 {
			// This is the game ID label
			continue
		}
		components := strings.Split(strings.TrimSpace(drawString), " ")
		quantity, _ := strconv.ParseInt(components[0], 10, 0)
		color := components[1]
		parsedDraws[i-1] = draw{
			quantity: int(quantity),
			color:    color,
		}
	}
	return drawGame{
		id:    int(gameId),
		draws: parsedDraws,
	}
}

func (game drawGame) isPossibleWithBag(bag drawBag) bool {
	minRed, minGreen, minBlue := game.minimumCubesRequired()
	return minRed <= bag.reds && minGreen <= bag.greens && minBlue <= bag.blues
}

func (game drawGame) power() int {
	minRed, minGreen, minBlue := game.minimumCubesRequired()
	return minRed * minGreen * minBlue
}

// red,green,blue
func (game drawGame) minimumCubesRequired() (int, int, int) {
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
