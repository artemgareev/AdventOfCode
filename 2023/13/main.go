package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/13/input.txt")
	if err != nil {
		panic(err)
	}

	p1 := 0
	p2 := 0

	for _, wallsRaw := range bytes.Split(file, []byte("\n\n")) {
		walls := [][]byte{}
		for _, line := range bytes.Split(wallsRaw, []byte("\n")) {
			walls = append(walls, line)
		}
		p1 += horizontal(walls, 0)*100 + vertical(walls, 0)
		p2 += horizontal(walls, 1)*100 + vertical(walls, 1)
	}

	fmt.Println(p1)
	fmt.Println(p2)
}

func vertical(walls [][]byte, badMirrorsAccepted int) int {
	cols := len(walls[0])
	for col := 0; col < cols-1; col++ {
		badMirror := 0
		for col2 := 0; col2 < cols; col2++ {
			left := col - col2
			right := col + 1 + col2
			if left >= 0 && left < right && right < cols {
				badMirror += compare2Columns(walls, left, right)
			}
		}

		if badMirror == badMirrorsAccepted {
			return col + 1
		}
	}
	return 0
}

func horizontal(walls [][]byte, badMirrorsAccepted int) int {
	rows := len(walls)

	for row := 0; row < rows-1; row++ {
		badMirror := 0
		for row2 := 0; row2 < rows; row2++ {
			up := row - row2
			down := row + 1 + row2
			if up >= 0 && up < down && down < rows {
				badMirror += compare2Rows(walls, up, down)
			}
		}
		if badMirror == badMirrorsAccepted {
			return 100 * (row + 1)
		}
	}
	return 0
}

func compare2Columns(walls [][]byte, l, r int) int {
	cnt := 0
	for i := 0; i < len(walls); i++ {
		if walls[i][l] != walls[i][r] {
			cnt++
		}
	}
	return cnt
}

func compare2Rows(walls [][]byte, l, r int) int {
	cnt := 0
	for i := 0; i < len(walls[0]); i++ {
		if walls[l][i] != walls[r][i] {
			cnt++
		}
	}
	return cnt
}
