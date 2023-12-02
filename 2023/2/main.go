package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	regID    = regexp.MustCompile("Game (\\d+):")
	regBlue  = regexp.MustCompile("(\\d+) blue")
	redRed   = regexp.MustCompile("(\\d+) red")
	regGreen = regexp.MustCompile("(\\d+) green")
)
var redCap, greenCap, blueCap = 12, 13, 14

func main() {
	file, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	validGamesNum := 0
	sumPowerSets := 0

	for _, line := range bytes.Split(file, []byte("\n")) {
		gameIsValid, minRed, minBlue, minGreen := gameValid(string(line))
		if gameIsValid {
			gameID := regID.FindStringSubmatch(string(line))
			validGamesNum += mustAtoi(gameID[1])
		}
		sumPowerSets += minRed * minBlue * minGreen
	}

	fmt.Println("Part1:", validGamesNum)
	fmt.Println("Part2:", sumPowerSets)
}

func gameValid(line string) (gamePossible bool, minRed int, minBlue int, minGreen int) {
	gamePossible = true
	for _, subGame := range strings.Split(line, ";") {
		redNum := calc(redRed.FindAllStringSubmatch(subGame, -1))
		blueNum := calc(regBlue.FindAllStringSubmatch(subGame, -1))
		greenNum := calc(regGreen.FindAllStringSubmatch(subGame, -1))

		minRed = max(minRed, redNum)
		minBlue = max(minBlue, blueNum)
		minGreen = max(minGreen, greenNum)
		if gamePossible && (redNum > redCap || blueNum > blueCap || greenNum > greenCap) {
			gamePossible = false
		}
	}
	return
}

func calc(matches [][]string) (total int) {
	for _, match := range matches {
		total += mustAtoi(match[1])
	}
	return
}

func mustAtoi(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
