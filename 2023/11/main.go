package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/11/input.txt")
	if err != nil {
		panic(err)
	}

	galaxy := [][]byte{}
	for _, line := range bytes.Split(file, []byte("\n")) {
		galaxy = append(galaxy, line)
	}

	emptyRows := []int{}
	emptyCols := []int{}
	for i, line := range galaxy {
		if string(line) == string(bytes.Repeat([]byte("."), len(line))) {
			emptyRows = append(emptyRows, i)
		}
	}
	for i := 0; i < len(galaxy[0]); i++ {
		col := []byte{}
		for j := 0; j < len(galaxy); j++ {
			col = append(col, galaxy[j][i])
		}
		if string(col) == string(bytes.Repeat([]byte("."), len(col))) {
			emptyCols = append(emptyCols, i)
		}
	}

	planets := map[[2]int]map[[2]int]int{}
	for i, line := range galaxy {
		for j, ch := range line {
			if ch == '#' {
				planets[[2]int{i, j}] = map[[2]int]int{}
			}
		}
	}
	for planet1 := range planets {
		for planet2 := range planets {
			if planet2 != planet1 {
				if _, ok := planets[planet2][planet1]; !ok {
					planets[planet1][planet2] = 0
				}
			}
		}
	}

	calcDist := func(expansionFactor int) int {
		sum := 0
		for planet, subPlanets := range planets {
			for connectedPlanet := range subPlanets {
				dist := shortestDist(planet, connectedPlanet)
				for _, row := range emptyRows {
					if max(planet[0], connectedPlanet[0]) >= row && min(planet[0], connectedPlanet[0]) <= row {
						dist += expansionFactor
					}
				}
				for _, col := range emptyCols {
					if max(planet[1], connectedPlanet[1]) >= col && min(planet[1], connectedPlanet[1]) <= col {
						dist += expansionFactor
					}
				}

				sum += dist
			}
		}
		return sum
	}

	fmt.Println(calcDist(1))
	fmt.Println(calcDist(1000000 - 1))
}

func shortestDist(p1, p2 [2]int) int {
	dst := abc(p2[0]-p1[0]) + abc(p2[1]-p1[1])
	return dst
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func abc(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
