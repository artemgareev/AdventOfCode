package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/9/input.txt")
	if err != nil {
		panic(err)
	}

	nums := [][]int{}
	for _, line := range bytes.Split(file, []byte("\n")) {
		numbersRaw := bytes.Split(line, []byte(" "))
		numbers := []int{}
		for _, numberRaw := range numbersRaw {
			numbers = append(numbers, mustAtoi(string(numberRaw)))
		}
		nums = append(nums, numbers)
	}

	p1, p2 := 0, 0
	for _, line := range nums {
		pyramid := [][]int{}
		pyramid = append(pyramid, line)

		// build pyramid
		pointer := 0
		for {
			sumAllElem := 0
			pyramidLine := []int{}
			for i := 1; i < len(pyramid[pointer]); i++ {
				diff := pyramid[pointer][i] - pyramid[pointer][i-1]
				pyramidLine = append(pyramidLine, diff)
				sumAllElem += diff
			}
			pyramid = append(pyramid, pyramidLine)
			pointer++

			if sumAllElem == 0 {
				break
			}
		}

		// calculate first and last val
		for i := len(pyramid) - 1; i > 0; i-- {
			pyramid[i-1] = append(
				[]int{pyramid[i-1][0] - pyramid[i][0]},
				pyramid[i-1]...,
			)
			pyramid[i-1] = append(
				pyramid[i-1],
				lastElem(pyramid[i])+lastElem(pyramid[i-1]),
			)
		}
		p1 += pyramid[0][len(pyramid[0])-1]
		p2 += pyramid[0][0]
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func lastElem(arr []int) int {
	return arr[len(arr)-1]
}

func mustAtoi(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}
