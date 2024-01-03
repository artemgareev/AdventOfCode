package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

var dir = map[uint8][2]int{
	'D': {1, 0}, 'U': {-1, 0}, 'L': {0, -1}, 'R': {0, 1},
}

func main() {
	directory, _ := os.Getwd()
	file, err := os.ReadFile(directory + "/18/input.txt")
	if err != nil {
		panic(err)
	}

	perimetr1 := 0
	points := [][2]int{{0, 0}}
	for _, line := range bytes.Split(file, []byte("\n")) {
		p := bytes.Split(line, []byte(" "))
		move, val := dir[p[0][0]], mustAtoi(p[1])
		perimetr1 += val
		lastPoint := points[len(points)-1]
		points = append(points, [2]int{lastPoint[0] + move[0]*val, lastPoint[1] + move[1]*val})
	}
	fmt.Println(calcArea(points, perimetr1))

	perimetr2 := 0
	points2 := [][2]int{{0, 0}}
	for _, line := range bytes.Split(file, []byte("\n")) {
		hex := bytes.Split(line, []byte(" "))[2]
		h := hex[2 : len(hex)-2]
		d := hex[len(hex)-2]

		val64, _ := strconv.ParseInt(string(h), 16, 64)
		val := int(val64)
		move := dir[[]byte("RDLU")[mustAtoi2(d)]]
		perimetr2 += val

		lastPoint := points2[len(points2)-1]
		points2 = append(points2, [2]int{lastPoint[0] + move[0]*val, lastPoint[1] + move[1]*val})
	}
	fmt.Println(calcArea(points2, perimetr2))
}

func calcArea(points [][2]int, perimeter int) int {
	sum := 0
	sum += points[0][0] * (points[len(points)-1][1] - points[1][1])
	for i := 1; i < len(points)-1; i++ {
		sum += points[i][0] * (points[i-1][1] - points[i+1][1])
	}
	sum += points[len(points)-1][0] * (points[len(points)-2][1] - points[0][1])

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	sum = abc(sum / 2)
	area := sum - perimeter/2 + 1
	return area + perimeter
}

func mustAtoi(str []byte) int {
	val, _ := strconv.Atoi(string(str))
	return val
}
func mustAtoi2(str byte) int {
	val, _ := strconv.Atoi(string(str))
	return val
}

func abc(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
