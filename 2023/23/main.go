package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/23/input.txt")
	if err != nil {
		panic(err)
	}

	island := bytes.Split(file, []byte("\n"))

	cols, rows := len(island), len(island[0])

	startPoint := [2]int{0, bytes.Index(island[0], []byte("."))}
	endPoint := [2]int{len(island) - 1, bytes.Index(island[len(island)-1], []byte("."))}
	points := [][2]int{startPoint, endPoint}
	for i, line := range island {
		for j, ch := range line {
			if ch == '#' {
				continue
			}
			neighbours := 0
			for _, move := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
				ii, jj := i+move[0], j+move[1]
				if !(ii >= 0 && ii < cols && jj >= 0 && jj < rows) || island[ii][jj] == '#' {
					continue
				}
				neighbours++
				if neighbours >= 3 {
					points = append(points, [2]int{i, j})
				}
			}
		}
	}

	graph1 := buildGraph(island, points, false)
	graph2 := buildGraph(island, points, true)

	fmt.Println(dfs(startPoint, endPoint, graph1, map[[2]int]bool{}))
	fmt.Println(dfs(startPoint, endPoint, graph2, map[[2]int]bool{}))
}

func buildGraph(island [][]byte, points [][2]int, part2 bool) map[[2]int]map[[2]int]int {
	cols, rows := len(island), len(island[0])
	graph := map[[2]int]map[[2]int]int{}
	for _, point := range points {
		graph[point] = map[[2]int]int{}
	}

	dirs := map[uint8][][2]int{
		'>': {{0, 1}},
		'<': {{0, -1}},
		'^': {{-1, 0}},
		'v': {{1, 0}},
		'.': {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
	}
	for _, point := range points {
		stack := [][3]int{{point[0], point[1], 0}}
		seen := map[[2]int]int{}
		for len(stack) > 0 {
			i, j, d := stack[0][0], stack[0][1], stack[0][2]
			stack = stack[1:]

			if _, ok := graph[[2]int{i, j}]; ok && d != 0 {
				graph[point][[2]int{i, j}] = d
				continue
			}

			moves := dirs['.']
			if !part2 {
				moves = dirs[island[i][j]]
			}
			for _, move := range moves {
				ii, jj := i+move[0], j+move[1]
				if !(ii >= 0 && ii < cols && jj >= 0 && jj < rows) || island[ii][jj] == '#' {
					continue
				}
				if _, ok := seen[[2]int{ii, jj}]; ok {
					continue
				}
				stack = append(stack, [3]int{ii, jj, d + 1})
				seen[[2]int{ii, jj}] = d
			}
		}
	}

	return graph
}

func dfs(currentP [2]int, endP [2]int, graph map[[2]int]map[[2]int]int, visited map[[2]int]bool) int {
	if currentP == endP {
		return 0
	}

	maxSteps := -math.MaxInt
	visited[currentP] = true
	for point := range graph[currentP] {
		if _, ok := visited[point]; ok {
			continue
		}
		maxSteps = max(maxSteps, dfs(point, endP, graph, visited)+graph[currentP][point])
	}
	delete(visited, currentP)

	return maxSteps
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
