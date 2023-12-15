package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	f, _ := os.Open("./input.txt")
	handsAndBids := parseInput(f)
	totalWinningsPart1 := calculateWinnings(handsAndBids)

	handsAndBidsWithJokers := make([]HandAndBid, len(handsAndBids))
	for i := range handsAndBids {
		handsAndBidsWithJokers[i] = handsAndBids[i]
		handsAndBidsWithJokers[i].hand = handsAndBidsWithJokers[i].hand.WithJacksAsJokers()
	}
	totalWinningsPart2 := calculateWinnings(handsAndBidsWithJokers)

	fmt.Printf("Total Winnings Part 1 %d\n", totalWinningsPart1)
	fmt.Printf("Total Winnings Part 2 %d\n", totalWinningsPart2)
}

func calculateWinnings(handsAndBids []HandAndBid) int {
	slices.SortFunc(handsAndBids, func(a, b HandAndBid) int {
		return a.hand.Compare(b.hand)
	})
	totalWinnings := 0
	for i, handAndBid := range handsAndBids {
		totalWinnings += handAndBid.bid * (i + 1)
	}
	return totalWinnings
}
