package day2

import (
	"regexp"
	"strconv"
	"strings"
)

func parseGame(game string) DrawGame {
	pattern := regexp.MustCompile("[:,;]\\s?")
	components := pattern.Split(game, -1)
	gameId, _ := strconv.ParseInt(strings.Split(components[0], " ")[1], 10, 0)
	parsedDraws := make([]Draw, max(0, len(components)-1))
	for i, drawString := range components {
		if i == 0 {
			// This is the game ID label
			continue
		}
		components := strings.Split(strings.TrimSpace(drawString), " ")
		quantity, _ := strconv.ParseInt(components[0], 10, 0)
		color := components[1]
		parsedDraws[i-1] = Draw{
			quantity: int(quantity),
			color:    color,
		}
	}
	return DrawGame{
		id:    int(gameId),
		draws: parsedDraws,
	}
}
