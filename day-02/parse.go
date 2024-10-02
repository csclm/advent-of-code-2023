package day2

import (
	"aoc-2023/aoc-lib"
	"regexp"
	"strings"
)

func parseGame(game string) DrawGame {
	pattern := regexp.MustCompile("[:,;]\\s?")
	components := pattern.Split(game, -1)
	gameId := aoc.MustParseInt(strings.Split(components[0], " ")[1])
	parsedDraws := make([]Draw, max(0, len(components)-1))
	for i, drawString := range components {
		if i == 0 {
			// This is the game ID label
			continue
		}
		components := strings.Split(strings.TrimSpace(drawString), " ")
		quantity := aoc.MustParseInt(components[0])
		color := components[1]
		parsedDraws[i-1] = Draw{
			quantity: quantity,
			color:    color,
		}
	}
	return DrawGame{
		id:    gameId,
		draws: parsedDraws,
	}
}
