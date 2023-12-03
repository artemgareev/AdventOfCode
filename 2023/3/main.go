package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/2023/3/input.txt")
	if err != nil {
		panic(err)
	}

	lines := bytes.Split(file, []byte("\n"))
	m := len(lines)
	n := len(lines[0])

	sum := 0
	gearsSet := map[[2]int][]int{}
	for i := 0; i < m; i++ {
		num := 0
		hasAdjacentSymbol := false
		gears := map[[2]int]struct{}{}
		for j := 0; j < n+1; j++ {
			if j < n && isDigit(i, j, lines) {
				val, _ := strconv.Atoi(string(lines[i][j]))
				num = 10*num + val
				for _, ii := range []int{-1, 0, 1} {
					for _, jj := range []int{-1, 0, 1} {
						II, JJ := i+ii, j+jj
						if II >= 0 && II < m && JJ >= 0 && JJ < n {
							ch := lines[II][JJ]
							if notNumOrPeriod(ch) {
								hasAdjacentSymbol = true
							}
							if ch == '*' {
								gears[[2]int{II, JJ}] = struct{}{}
							}
						}

					}
				}
			} else if num > 0 {
				for gear := range gears {
					gearsSet[gear] = append(gearsSet[gear], num)
				}
				if hasAdjacentSymbol {
					sum += num
				}
				hasAdjacentSymbol = false
				num = 0
				gears = map[[2]int]struct{}{}
			}
		}
	}

	fmt.Println(sum)
	sum2 := 0
	for _, nums := range gearsSet {
		if len(nums) == 2 {
			sum2 += nums[0] * nums[1]
		}
	}
	fmt.Println(sum2)
}

func notNumOrPeriod(num byte) bool {
	if (num >= '0' && num <= '9') || num == '.' {
		return false
	}

	return true
}

func isDigit(i, j int, lines [][]byte) bool {
	return lines[i][j] >= '0' && lines[i][j] <= '9'
}
