package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/17/input.txt")
	if err != nil {
		panic(err)
	}

	heatmap := bytes.Split(file, []byte("\n"))

	fmt.Println(bfs(heatmap))
	fmt.Println(bfs2(heatmap))
}

func bfs(heatMap [][]byte) int {
	visited := map[[5]int]bool{}
	queue := MinHeap{&Elem{0, 0, 0, 0, 0, 0}}

	for len(queue) > 0 {
		el := heap.Pop(&queue).(*Elem)
		hl, i, j, dI, dJ, n := el.hl, el.i, el.j, el.dI, el.dJ, el.n

		if i == len(heatMap)-1 && j == len(heatMap[0])-1 {
			return hl
		}

		if visited[[5]int{i, j, dI, dJ, n}] {
			continue
		}
		visited[[5]int{i, j, dI, dJ, n}] = true

		// keep going same direction
		if n < 3 && [2]int{dI, dJ} != [2]int{0, 0} {
			iI := i + dI
			jJ := j + dJ
			if iI >= 0 && jJ >= 0 && iI < len(heatMap) && jJ < len(heatMap[0]) {
				cost := hl + mustAtoi(heatMap[iI][jJ])
				heap.Push(&queue, &Elem{cost, iI, jJ, dI, dJ, n + 1})
			}
		}

		// make turn
		for _, move := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if move == [2]int{dI, dJ} || move == [2]int{-dI, -dJ} {
				continue
			}
			iI := i + move[0]
			jJ := j + move[1]
			if iI >= 0 && jJ >= 0 && iI < len(heatMap) && jJ < len(heatMap[0]) {
				cost := hl + mustAtoi(heatMap[iI][jJ])
				heap.Push(&queue, &Elem{cost, iI, jJ, move[0], move[1], 1})

			}
		}
	}

	return 0
}

func bfs2(heatMap [][]byte) int {
	visited := map[[5]int]bool{}
	queue := MinHeap{&Elem{0, 0, 0, 0, 0, 0}}

	for len(queue) > 0 {
		el := heap.Pop(&queue).(*Elem)
		hl, i, j, dI, dJ, n := el.hl, el.i, el.j, el.dI, el.dJ, el.n

		if i == len(heatMap)-1 && j == len(heatMap[0])-1 && n >= 4 {
			return hl
		}

		if visited[[5]int{i, j, dI, dJ, n}] {
			continue
		}
		visited[[5]int{i, j, dI, dJ, n}] = true

		// keep going same direction
		if n < 10 && [2]int{dI, dJ} != [2]int{0, 0} {
			iI := i + dI
			jJ := j + dJ
			if iI >= 0 && jJ >= 0 && iI < len(heatMap) && jJ < len(heatMap[0]) {
				cost := hl + mustAtoi(heatMap[iI][jJ])
				heap.Push(&queue, &Elem{cost, iI, jJ, dI, dJ, n + 1})
			}
		}

		// make turn
		if n >= 4 || [2]int{dI, dJ} == [2]int{0, 0} {
			for _, move := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
				if move == [2]int{dI, dJ} || move == [2]int{-dI, -dJ} {
					continue
				}
				iI := i + move[0]
				jJ := j + move[1]
				if iI >= 0 && jJ >= 0 && iI < len(heatMap) && jJ < len(heatMap[0]) {
					cost := hl + mustAtoi(heatMap[iI][jJ])
					heap.Push(&queue, &Elem{cost, iI, jJ, move[0], move[1], 1})

				}
			}
		}
	}

	return 0
}

func mustAtoi(str byte) int {
	val, _ := strconv.Atoi(string(str))
	return val
}

type Elem struct {
	hl, i, j, dI, dJ, n int
}

type MinHeap []*Elem

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].hl < h[j].hl }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(*Elem)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
