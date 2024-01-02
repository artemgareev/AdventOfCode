package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/12/input.txt")
	if err != nil {
		panic(err)
	}

	p1, p2 := 0, 0
	for _, line := range bytes.Split(file, []byte("\n")) {
		parts := bytes.Split(line, []byte(" "))

		springs := parts[0]
		arrangements := toArray(bytes.Split(parts[1], []byte(",")))
		p1 += dfs(springs, arrangements, map[string]int{})

		springsExp, arrangementsExp := expand(springs, arrangements)
		p2 += dfs(springsExp, arrangementsExp, map[string]int{})
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func dfs(springs []byte, arrangements []int, memo map[string]int) int {
	if len(springs) == 0 {
		if len(arrangements) == 0 {
			return 1
		}
		return 0
	}

	if len(arrangements) == 0 {
		if slices.Contains(springs, '#') {
			return 0
		}
		return 1
	}

	key := string(springs) + toString(arrangements)
	if r, ok := memo[key]; ok {
		return r
	}

	result := 0
	if springs[0] == '.' || springs[0] == '?' {
		result += dfs(springs[1:], arrangements, memo)
	}

	if springs[0] == '#' || springs[0] == '?' {
		if arrangements[0] <= len(springs) && !slices.Contains(springs[:arrangements[0]], '.') && (arrangements[0] == len(springs) || springs[arrangements[0]] != '#') {
			if arrangements[0] == len(springs) {
				result += dfs([]byte{}, arrangements[1:], memo)
			} else {
				result += dfs(springs[arrangements[0]+1:], arrangements[1:], memo)
			}
		}
	}

	memo[key] = result

	return result
}

func expand(springs []byte, arrangements []int) ([]byte, []int) {
	arrangementsCopy := slices.Clone(arrangements)
	for i := 0; i < 4; i++ {
		arrangements = append(arrangements, arrangementsCopy...)
	}

	springsCopy := slices.Clone(springs)
	for i := 0; i < 4; i++ {
		springs = append(springs, '?')
		springs = append(springs, springsCopy...)
	}

	return springs, arrangements
}

func toString(arr []int) string {
	result := ""
	for _, el := range arr {
		result += strconv.Itoa(el)
	}
	return result
}

func toArray(arrB [][]byte) (arr []int) {
	for _, val := range arrB {
		arr = append(arr, mustAtoi(val))
	}

	return arr
}

func mustAtoi(str []byte) int {
	val, _ := strconv.Atoi(string(str))
	return val
}
