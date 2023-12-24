package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Hailstone struct {
	pX, pY, pZ float64
	vX, vY, vZ float64
	a, b, c    float64
}

// NewHailstone
// Ax+By+Cz+D=0
//
// Cartesian equation of a line:
// (x-x0)/pX=(y-y0)/pY
// x*pY-x0*pY=y*pX-y0*pX
// x*pY-y*pX = x0*pY-y0*pX
// a=pY, b=-pX, c= x0*pY-y0*pX
func NewHailstone(pX, pY, pZ, vX, vY, vZ float64) Hailstone {
	return Hailstone{
		pX: pX, pY: pY, pZ: pZ,
		vX: vX, vY: vY, vZ: vZ,
		a: vY,
		b: -vX,
		c: pX*vY - pY*vX,
	}
}

// a1/b1 = a2/b2
// a1*b2 = a2*b1
func isCollinear(b1, b2 Hailstone) bool {
	return b1.a*b2.b == b2.a*b1.b
}
func main() {
	dir, _ := os.Getwd()
	file, err := os.ReadFile(dir + "/2023/24/input.txt")
	if err != nil {
		panic(err)
	}

	hailstones := []Hailstone{}
	for _, line := range bytes.Split(file, []byte("\n")) {
		p := bytes.Split(line, []byte("@"))

		point := bytes.Split(bytes.Join(bytes.Fields(p[0]), nil), []byte(","))
		vector := bytes.Split(bytes.Join(bytes.Fields(p[1]), nil), []byte(","))

		h := NewHailstone(
			mustAtoi(point[0]), mustAtoi(point[1]), mustAtoi(point[2]),
			mustAtoi(vector[0]), mustAtoi(vector[1]), mustAtoi(vector[2]),
		)
		hailstones = append(hailstones, h)

	}

	total := 0
	for i, hr1 := range hailstones {
		for _, hr2 := range hailstones[:i] {
			if isCollinear(hr1, hr2) {
				continue
			}

			a1, b1, c1 := hr1.a, hr1.b, hr1.c
			a2, b2, c2 := hr2.a, hr2.b, hr2.c
			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c1*a2 - c2*a1) / (b1*a2 - b2*a1)
			// 			if 7 <= x && x <= 27 && 7 <= y && y <= 27 {
			if 200000000000000 <= x && x <= 400000000000000 && 200000000000000 <= y && y <= 400000000000000 {
				if ((x-hr1.pX)*hr1.vX > 0 && (y-hr1.pY)*hr1.vY > 0) && ((x-hr2.pX)*hr2.vX > 0 && (y-hr2.pY)*hr2.vY > 0) {
					total++
				}
			}
		}
	}
	fmt.Println(total)
}

func mustAtoi(str []byte) float64 {
	val, _ := strconv.Atoi(string(str))
	return float64(val)
}
