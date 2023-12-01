package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

var leftReg1, _ = regexp.Compile("1|2|3|4|5|6|7|8|9")
var rightReg1, _ = regexp.Compile(".*(1|2|3|4|5|6|7|8|9)")

var leftReg2, _ = regexp.Compile("one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9")
var rightReg2, _ = regexp.Compile(".*(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)")

var strToDigit = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// cat input.txt | go run main.go
func main() {
	file, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	sum1 := 0
	sum2 := 0
	for _, line := range bytes.Split(file, []byte("\n")) {
		sum1 += strToDigit[leftReg1.FindString(string(line))]*10 +
			strToDigit[rightReg1.FindStringSubmatch(string(line))[1]]
		sum2 += strToDigit[leftReg2.FindString(string(line))]*10 +
			strToDigit[rightReg2.FindStringSubmatch(string(line))[1]]

	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
