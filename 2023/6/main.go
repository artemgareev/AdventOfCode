package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/5/input.txt")
	if err != nil {
		panic(err)
	}

	parts := bytes.Split(file, []byte("\n"))
	times := bytes.Fields(bytes.Split(parts[0], []byte(":"))[1])
	distances := bytes.Fields(bytes.Split(parts[1], []byte(":"))[1])

	bigTime := bytes.Join(times, []byte(""))
	bigBestDistanceTime := bytes.Join(distances, []byte(""))

	calc := func(times, distances [][]byte) int {
		result := 1
		for i := 0; i < len(times); i++ {
			givenTime := mustAtoi(string(times[i]))
			distanceBestTime := mustAtoi(string(distances[i]))

			winningTimes := 0
			for j := 0; j < givenTime; j++ {
				if (givenTime-j)*j > distanceBestTime {
					winningTimes++
				}
			}
			result *= winningTimes
		}
		return result
	}

	fmt.Println(calc(times, distances))
	fmt.Println(calc([][]byte{bigTime}, [][]byte{bigBestDistanceTime}))
}

func mustAtoi(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}
