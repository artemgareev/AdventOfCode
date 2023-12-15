package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type BoxLens struct {
	Label       string
	FocalLength int
}

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/15/input.txt")
	if err != nil {
		panic(err)
	}

	// rn=1 becomes 30.
	//Determine the ASCII code for the current character of the string.
	//Increase the current value by the ASCII code you just determined.
	//Set the current value to itself multiplied by 17.
	//Set the current value to the remainder of dividing itself by 256.
	boxes := map[int][]*BoxLens{}
	p1 := 0
	for _, cmd := range bytes.Split(file, []byte(",")) {
		p1 += hasher(0, cmd)
	}
	for _, cmd := range bytes.Split(file, []byte(",")) {
		if cmd[len(cmd)-2] == '=' {
			boxNum := hasher(0, cmd[:len(cmd)-2])
			focalLen := mustAtoi(string(cmd[len(cmd)-1:]))
			lensLabel := string(cmd[:len(cmd)-2])

			found := false
			if _, ok := boxes[boxNum]; ok {
				for i := 0; i < len(boxes[boxNum]); i++ {
					if boxes[boxNum][i].Label == lensLabel {
						boxes[boxNum][i].FocalLength = focalLen
						found = true
						break
					}
				}
			}
			if !found {
				boxes[boxNum] = append(boxes[boxNum], &BoxLens{lensLabel, focalLen})
			}
		}
		if cmd[len(cmd)-1] == '-' {
			boxNum := hasher(0, cmd[:len(cmd)-1])
			lensLabel := string(cmd[:len(cmd)-1])

			boxElems := boxes[boxNum]
			for i := 0; i < len(boxElems); i++ {
				if boxElems[i].Label == lensLabel {
					boxes[boxNum] = remove(boxElems, i)
					break
				}

			}
		}
	}
	p2 := 0
	for i := 0; i < 256; i++ {
		if boxVals, ok := boxes[i]; ok {
			for idx, boxElem := range boxVals {
				p2 += (i + 1) * (idx + 1) * boxElem.FocalLength
			}
		}
	}
	fmt.Println(p1)
	fmt.Println(p2)
}

func hasher(prevVal int, cmd []byte) int {
	currentValue := prevVal
	for _, ch := range cmd {
		currentValue = (currentValue + int(ch)) * 17 % 256
	}
	return currentValue
}

func remove(slice []*BoxLens, s int) []*BoxLens {
	return append(slice[:s], slice[s+1:]...)
}

func mustAtoi(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}
