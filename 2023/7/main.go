package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/7/input.txt")
	if err != nil {
		panic(err)
	}
	cards := [][2][]byte{}
	for _, line := range bytes.Split(file, []byte("\n")) {
		parts := bytes.Split(line, []byte(" "))
		// [2][]byte{hand, bid}
		cards = append(cards, [2][]byte{parts[0], parts[1]})
	}

	fmt.Println(TotalWinnings(cards, RankHand, CompareHands))
	fmt.Println(TotalWinnings(cards, RankHandWithJoker, CompareHandsWithJoker))
}

func TotalWinnings(cards [][2][]byte, rankFN func(hand []byte) int, cmpFn func(hand1 []byte, hand2 []byte) bool) int {
	// sort from the lowest to the highest hand
	sort.SliceStable(cards, func(i, j int) bool {
		rank1 := rankFN(cards[i][0])
		rank2 := rankFN(cards[j][0])

		if rank1 == rank2 {
			return cmpFn(cards[i][0], cards[j][0])
		}
		return rank1 < rank2
	})

	totalWinnings := 0
	for i := 1; i <= len(cards); i++ {
		totalWinnings += mustAtoi(string(cards[i-1][1])) * i
	}
	return totalWinnings
}

func CompareHands(hand1 []byte, hand2 []byte) bool {
	var cardRank = map[byte]int{
		'A': 12, 'K': 11, 'Q': 10, 'J': 9, 'T': 8,
		'9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0,
	}
	for i := 0; i < len(hand1); i++ {
		if cardRank[hand1[i]] == cardRank[hand2[i]] {
			continue
		}
		return cardRank[hand1[i]] < cardRank[hand2[i]]
	}

	return true
}

func RankHand(hand []byte) int {
	cards := make(map[string]int, 5)
	for _, card := range hand {
		cards[string(card)]++
	}
	sortedVal := make([]int, 0)
	for _, val := range cards {
		sortedVal = append(sortedVal, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedVal)))

	// Five of a kind
	if sortedVal[0] == 5 {
		return 6
	}
	// Four of a kind
	if sortedVal[0] == 4 {
		return 5
	}
	// Full house
	if sortedVal[0] == 3 && sortedVal[1] == 2 {
		return 4
	}
	// Three of a kind
	if sortedVal[0] == 3 && sortedVal[1] == 1 {
		return 3
	}
	// Two pair
	for sortedVal[0] == 2 && sortedVal[1] == 2 {
		return 2
	}
	// One pair
	for sortedVal[0] == 2 {
		return 1
	}
	// No pairs
	return 0
}

func CompareHandsWithJoker(hand1 []byte, hand2 []byte) bool {
	var cardRankWithJoker = map[byte]int{
		'A': 12, 'K': 11, 'Q': 10, 'T': 8,
		'9': 7, '8': 6, '7': 5, '6': 4, '5': 3, '4': 2, '3': 1, '2': 0, 'J': -1,
	}
	for i := 0; i < len(hand1); i++ {
		if cardRankWithJoker[hand1[i]] == cardRankWithJoker[hand2[i]] {
			continue
		}
		return cardRankWithJoker[hand1[i]] < cardRankWithJoker[hand2[i]]
	}

	return true
}

func RankHandWithJoker(hand []byte) int {
	numOfJokers := 0
	cards := make(map[string]int, 5)
	for _, card := range hand {
		if card == 'J' {
			numOfJokers++
		} else {
			cards[string(card)]++
		}
	}
	sortedVal := make([]int, 0)
	for _, val := range cards {
		sortedVal = append(sortedVal, val)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedVal)))

	if numOfJokers == 5 {
		return 6
	}
	// Five of a kind
	if sortedVal[0]+numOfJokers == 5 {
		return 6
	}
	// Four of a kind
	if sortedVal[0]+numOfJokers == 4 {
		return 5
	}
	// Full house
	if (sortedVal[0]+numOfJokers == 3 && sortedVal[1] == 2) || (sortedVal[0] == 3 && sortedVal[1]+numOfJokers == 2) {
		return 4
	}
	// Three of a kind
	if (sortedVal[0]+numOfJokers == 3 && sortedVal[1] == 1) || (sortedVal[0] == 3 && sortedVal[1]+numOfJokers == 1) {
		return 3
	}
	// Two pair
	for (sortedVal[0]+numOfJokers == 2 && sortedVal[1] == 2) || (sortedVal[0] == 2 && sortedVal[1]+numOfJokers == 2) {
		return 2
	}
	// One pair
	for sortedVal[0]+numOfJokers == 2 {
		return 1
	}
	// No pairs
	return 0
}

func mustAtoi(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}
