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

var directionNextMove = map[int][2]int{
	upDir:    {-1, 0},
	downDir:  {1, 0},
	leftDir:  {0, -1},
	rightDir: {0, 1},
}

var objectDirection = map[int]map[uint8][][3]int{
	upDir: {
		'-':  {b(directionNextMove[leftDir], leftDir), b(directionNextMove[rightDir], rightDir)},
		'|':  {b(directionNextMove[upDir], upDir)},
		'\\': {b(directionNextMove[leftDir], leftDir)},
		'/':  {b(directionNextMove[rightDir], rightDir)},
	},
	downDir: {
		'-':  {b(directionNextMove[leftDir], leftDir), b(directionNextMove[rightDir], rightDir)},
		'|':  {b(directionNextMove[downDir], downDir)},
		'\\': {b(directionNextMove[rightDir], rightDir)},
		'/':  {b(directionNextMove[leftDir], leftDir)},
	},
	leftDir: {
		'-':  {b(directionNextMove[leftDir], leftDir)},
		'|':  {b(directionNextMove[upDir], upDir), b(directionNextMove[downDir], downDir)},
		'\\': {b(directionNextMove[upDir], upDir)},
		'/':  {b(directionNextMove[downDir], downDir)},
	},
	rightDir: {
		'-':  {b(directionNextMove[rightDir], rightDir)},
		'|':  {b(directionNextMove[upDir], upDir), b(directionNextMove[downDir], downDir)},
		'\\': {b(directionNextMove[downDir], downDir)},
		'/':  {b(directionNextMove[upDir], upDir)},
	},
}

func main() {
	val, ok := objectDirection[rightDir]['|']
	_, _ = val, ok

	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/16/input.txt")
	if err != nil {
		panic(err)
	}

	labyrinth := [][]byte{}
	for _, line := range bytes.Split(file, []byte("\n")) {
		labyrinth = append(labyrinth, line)
	}

	fmt.Println(bfs(labyrinth, 0, 0, rightDir))

	p2 := -1
	for i := 0; i < len(labyrinth); i++ {
		p2 = max(p2, bfs(labyrinth, 0, i, downDir))
		p2 = max(p2, bfs(labyrinth, len(labyrinth)-1, i, upDir))

	}
	for i := 0; i < len(labyrinth[0]); i++ {
		p2 = max(p2, bfs(labyrinth, i, 0, rightDir))
		p2 = max(p2, bfs(labyrinth, 0, len(labyrinth[0])-1, leftDir))
	}

	fmt.Println(p2)
}

func bfs(lab [][]byte, iS, iJ, dir int) int {
	stack := [][3]int{{iS, iJ, dir}}
	visited := map[[3]int]bool{}
	for len(stack) > 0 {
		i, j, dir := stack[0][0], stack[0][1], stack[0][2]
		stack = stack[1:]

		if !(i >= 0 && j >= 0 && i < len(lab) && j < len(lab[0])) {
			continue
		}
		if _, ok := visited[[3]int{i, j, dir}]; ok {
			continue
		}

		curr := lab[i][j]
		visited[[3]int{i, j, dir}] = true
		if nextDirs, ok := objectDirection[dir][curr]; ok {
			for _, nextDir := range nextDirs {
				stack = append(stack, [3]int{i + nextDir[0], j + nextDir[1], nextDir[2]})
			}
		} else {
			stack = append(stack, [3]int{i + directionNextMove[dir][0], j + directionNextMove[dir][1], dir})
		}
	}

	res := map[[2]int]struct{}{}
	for dot := range visited {
		res[[2]int{dot[0], dot[1]}] = struct{}{}
	}

	return len(res)
}

func b(arr [2]int, dir int) [3]int {
	return [3]int{arr[0], arr[1], dir}
}
