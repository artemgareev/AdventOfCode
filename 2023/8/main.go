package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/8/input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(file, []byte("\n"))
	instructions := parts[0]

	network := map[string][2]string{}
	for _, node := range parts[2:] {
		nodeParts := bytes.Split(node, []byte("="))
		start := nodeParts[0][:3]

		nodePaths := bytes.Split(nodeParts[1], []byte(","))
		left := nodePaths[0][len(nodePaths[0])-3:]
		right := nodePaths[1][1:4]

		network[string(start)] = [2]string{string(left), string(right)}
	}

	part1(instructions, network)
	part2(instructions, network)
}

func part1(instructions []byte, network map[string][2]string) {
	currentNode := "AAA"
	pointer := 0
	steps := 0
	for currentNode != "ZZZ" {
		if instructions[pointer] == 'L' {
			currentNode = network[currentNode][0]
		}
		if instructions[pointer] == 'R' {
			currentNode = network[currentNode][1]
		}
		pointer++
		if pointer == len(instructions) {
			pointer = 0
		}
		steps++
	}

	fmt.Println(steps)
}

func part2(instructions []byte, network map[string][2]string) {
	startsWithA := []string{}
	for node := range network {
		if node[2] == 'A' {
			startsWithA = append(startsWithA, node)
		}
	}

	stepsZ := []int{}
	for _, currentNode := range startsWithA {
		pointer := 0
		steps := 0
		for currentNode[2] != 'Z' {
			if instructions[pointer] == 'L' {
				currentNode = network[currentNode][0]
			}
			if instructions[pointer] == 'R' {
				currentNode = network[currentNode][1]
			}
			pointer++
			if pointer == len(instructions) {
				pointer = 0
			}
			steps++
		}
		stepsZ = append(stepsZ, steps)
	}

	fmt.Println(lcm(1, 1, stepsZ...))

}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
