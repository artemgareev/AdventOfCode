package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/2023/4/input.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(file, []byte("\n"))

	totalScratchedCards := 0
	totalPoints := 0
	cardCopies := map[int]int{}
	for i, line := range lines {
		gameInfo := strings.Split(
			strings.Split(string(line), ":")[1],
			"|",
		)
		winningCards := map[int]struct{}{}
		for _, card := range strings.Split(gameInfo[0], " ") {
			if card == "" {
				continue
			}
			val, _ := strconv.Atoi(card)
			winningCards[val] = struct{}{}
		}
		winningCardsCount := 0
		for _, myCard := range strings.Split(gameInfo[1], " ") {
			if myCard == "" {
				continue
			}
			val, _ := strconv.Atoi(myCard)
			if _, ok := winningCards[val]; ok {
				winningCardsCount++
			}
		}

		for k := 0; k < winningCardsCount; k++ {
			cardCopies[k+i+2] += cardCopies[i+1] + 1
		}

		if winningCardsCount > 0 {
			totalPoints += 1 << (winningCardsCount - 1)
		}
		totalScratchedCards += 1 + cardCopies[i+1]
	}

	fmt.Println(totalPoints)
	fmt.Println(totalScratchedCards)
}
