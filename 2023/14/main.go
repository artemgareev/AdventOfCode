package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/14/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(
		score(tiltNorth(bytes.Split(file, []byte("\n")))),
	)
	platform := bytes.Split(file, []byte("\n"))
	cycles := 1000000000
	hashes := map[string]int{}
	cycleFound := false
	for i := 0; i < cycles; i++ {
		for j := 0; j < 4; j++ {
			platform = tiltNorth(platform)
			platform = rotateClockwise(platform)
		}
		hash := string(bytes.Join(platform, []byte("")))
		if idx, ok := hashes[hash]; ok && !cycleFound {
			cycleLen := i - idx
			cycleToSkip := (cycles - i) / cycleLen
			i += cycleLen * cycleToSkip
			cycleFound = true

		}
		hashes[hash] = i
	}
	fmt.Println(score(platform))
}

func tiltNorth(platform [][]byte) [][]byte {
	rows := len(platform)
	cols := len(platform[0])

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			for row2 := 0; row2 < rows; row2++ {
				if platform[row2][col] == 'O' && row2 > 0 && platform[row2-1][col] == '.' {
					platform[row2][col] = '.'
					platform[row2-1][col] = 'O'
				}
			}
		}
	}
	return platform
}

func rotateClockwise(platform [][]byte) [][]byte {
	rotated := [][]byte{}
	for col := 0; col < len(platform[0]); col++ {
		colLine := []byte{}
		for row := len(platform) - 1; row >= 0; row-- {
			colLine = append(colLine, platform[row][col])
		}
		rotated = append(rotated, colLine)
	}
	return rotated
}

func score(platform [][]byte) int {
	sum := 0
	for i := 0; i < len(platform); i++ {
		rocks := 0
		for j := 0; j < len(platform[0]); j++ {
			if platform[i][j] == 'O' {
				rocks++
			}
		}
		sum += rocks * (len(platform) - i)
	}

	return sum
}
