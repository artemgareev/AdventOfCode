package main

import (
	"bytes"
	"fmt"
	"os"
)

const (
	upDir    = 0
	downDir  = 1
	leftDir  = 2
	rightDir = 3
)

var directionNextMove = map[int]map[uint8][3]int{
	upDir: { // bottom-up
		'|': {-1, 0, upDir},
		'7': {0, -1, leftDir},
		'F': {0, 1, rightDir},
	},
	downDir: { // up-bottom
		'|': {1, 0, downDir},
		'L': {0, 1, rightDir},
		'J': {0, -1, leftDir},
	},
	leftDir: { // right-left
		'-': {0, -1, leftDir},
		'L': {-1, 0, upDir},
		'F': {1, 0, downDir},
	},
	rightDir: { // left-right
		'-': {0, 1, rightDir},
		'J': {-1, 0, upDir},
		'7': {1, 0, downDir},
	},
}

var maxSteps = 0

var labyrinth = [][]uint8{}
var labyrinthStr = [][]string{}

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/10/input.txt")
	if err != nil {
		panic(err)
	}

	iS, jS := 0, 0
	for ii, line := range bytes.Split(file, []byte("\n")) {
		labyrinthLine := []uint8{}
		labyrinthStrLine := []string{}
		for jj, ch := range line {
			if ch == 'S' {
				iS = ii
				jS = jj
			}
			labyrinthLine = append(labyrinthLine, ch)
			labyrinthStrLine = append(labyrinthStrLine, string(ch))
		}
		labyrinthStr = append(labyrinthStr, labyrinthStrLine)
		labyrinth = append(labyrinth, labyrinthLine)
	}

	bfs(iS, jS)
	p2 := 0

	// for any point that not in the loop check all |, L, J pipes from the left of it
	// if there is an odd number, then it must be inside the loop
	for i := 0; i < len(labyrinth); i++ {
		for j := 0; j < len(labyrinth[0]); j++ {
			if visited[[2]int{i, j}] {
				continue
			}
			visitedPipesFromLeft := 0
			for k := j - 1; k >= 0; k-- {
				if visited[[2]int{i, k}] && (labyrinth[i][k] == '|' || labyrinth[i][k] == 'L' || labyrinth[i][k] == 'J') {
					visitedPipesFromLeft++
				}
			}
			if visitedPipesFromLeft > 0 && visitedPipesFromLeft%2 != 0 {
				labyrinthStr[i][j] = "\033[94m" + labyrinthStr[i][j] + "\033[00m"
				p2++
			}

		}
	}

	for _, labLine := range labyrinthStr {
		for _, ch := range labLine {
			fmt.Print(ch)
		}
		fmt.Println()
	}
	fmt.Println(maxSteps)
	fmt.Println(p2)
}

var visited = map[[2]int]bool{}

func bfs(iS, jS int) {
	stack := [][4]int{{iS, jS, 0, -1}}
	// {-1, 0}, {1, 0}, {0, -1}, {0, 1} <-> up, down, left, right
	for direction, nextMove := range [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		stack = append(stack, [4]int{iS + nextMove[0], jS + nextMove[1], 1, direction})
	}
	for len(stack) != 0 {
		ii, jj, steps, direction := stack[0][0], stack[0][1], stack[0][2], stack[0][3]
		stack = stack[1:]
		if isBadPath(ii, jj) {
			continue
		}
		maxSteps = max(maxSteps, steps)

		if nextMove, ok := directionNextMove[direction][labyrinth[ii][jj]]; ok {
			visited[[2]int{ii, jj}] = true

			labyrinthStr[ii][jj] = "\033[33m" + labyrinthStr[ii][jj] + "\033[00m"
			stack = append(stack, [4]int{ii + nextMove[0], jj + nextMove[1], steps + 1, nextMove[2]})
		}
	}
}

func isBadPath(i, j int) bool {
	return !(i >= 0 && i < len(labyrinth) && j >= 0 && j < len(labyrinth[0])) ||
		visited[[2]int{i, j}] ||
		labyrinth[i][j] == '.'
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
