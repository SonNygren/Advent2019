package day3

import (
	"advent2019/reader"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() error {

	input, err := reader.ParseAsString("day3/input")
	if err != nil {
		return err
	}

	wirePath := [][]string{}
	for _, v := range input {
		wirePath = append(wirePath, strings.Split(v, ","))
	}

	//for _, item := range wirePath[0] {
	//	fmt.Printf("item[0]: %v\n", string(item[0]))
	//	fmt.Printf("item[1:]: %v\n", string(item[1:]))
	//}
	//getPositions()
	return nil
}

func visualize(wirePath [][]string) error {
	//wires := []byte(wirePath)
	err := os.WriteFile("day3/output", []byte{}, 5)
	if err != nil {
		return err
	}

	return nil
}

type Point struct {
	X int
	Y int
}

type Wire struct {
	Path   []string
	Plot   map[string][]int
	Points []Point
}

func TraverseWire() error {

	wires := []Wire{
		{
			Path:   []string{"R5", "U4"},
			Plot:   map[string][]int{},
			Points: []Point{},
		},
		{
			Path:   []string{"U2", "R10"},
			Plot:   map[string][]int{},
			Points: []Point{},
		},
	}
	count := 0

	pons := map[int]Point{}

	for _, wire := range wires {

		xPos := 0
		yPos := 0
		wire.Plot["0,0"] = append(wire.Plot["0,0"], 0)
		pons[count] = Point{X: 0, Y: 0}
		for index, value := range wire.Path {
			count++

			switch string(value[0]) {
			case R:
				v, _ := strconv.Atoi(value[1:])
				xPos, _ = routeWire(wire.Plot, index, xPos+v, yPos)
				pons[count] = Point{
					X: xPos,
					Y: yPos,
				}
			case L:
				v, _ := strconv.Atoi(value[1:])
				xPos, _ = routeWire(wire.Plot, index, xPos-v, yPos)
				pons[count] = Point{
					X: xPos,
					Y: yPos,
				}

			case U:
				v, _ := strconv.Atoi(value[1:])
				_, yPos = routeWire(wire.Plot, index, xPos, yPos+v)
				pons[count] = Point{
					X: xPos,
					Y: yPos,
				}

			case D:
				v, _ := strconv.Atoi(value[1:])
				_, yPos = routeWire(wire.Plot, index, xPos, yPos-v)
				pons[count] = Point{
					X: xPos,
					Y: yPos,
				}
			}
		}
	}
	fmt.Printf("pons3: %v\n", pons)

	f := []Point{}
	q := []Point{}

	for i := 0; i <= 2; i++ {
		f = append(f, pons[i])
	}
	for i := 3; i <= 5; i++ {
		q = append(q, pons[i])
	}
	fmt.Printf("f: %v\n", f)
	fmt.Printf("q: %v\n", q)
	xs := getPositions(f, q)

	fmt.Printf("xs: %v\n", xs)
	return nil
}

func getBorders(borderA, borderB int) (int, int) {
	if borderB < borderA {
		temp := borderA
		borderA = borderB
		borderB = temp
	}
	return borderA, borderB
}

func getPositions(pointsA, pointsB []Point) []Point {
	Xings := []Point{}
	for i := 1; i < len(pointsA); i++ {

		xled := false
		start := 0
		end := 0
		wireACurrentX := pointsA[i].X
		wireAMinusOne := pointsA[i-1].X
		if wireACurrentX == wireAMinusOne { // vertical ram
			start, end = getBorders(pointsA[i].Y, pointsA[i-1].Y) // om wireA inte ändras på X axeln, titta på vilket värde som är större eller midnre
			xled = true
		} else { // horizontel ram
			start, end = getBorders(pointsA[i].X, pointsA[i-1].X)
			xled = false
		}

		for j := 1; j < len(pointsB); j++ {
			lastX := pointsB[j-1].X
			lastY := pointsB[j-1].Y

			currentX := pointsB[j].X
			currentY := pointsB[j].Y

			if xled {
				if end < pointsB[j].X || pointsB[j].X < start { // if its outside the borders
					continue
				}

				if currentX != lastX { // check for crossing in startingpoint which shouldnt be taken into account
					continue
				} else {
					//     .          last
					// -----------    pointsA[i].y
					//     .          current
					if currentY < lastY && pointsA[i].Y < lastY && currentY < pointsA[i].Y {
						Xings = append(Xings, Point{
							X: currentX,
							Y: pointsA[i].Y,
						})
					} else if lastY < currentY && pointsA[i].Y < currentY && lastY < pointsA[i].Y {
						Xings = append(Xings, Point{
							X: currentX,
							Y: pointsA[i].Y,
						})
					}
				}

			} else {
				if end < pointsB[j].Y || pointsB[j].Y < start {
					continue
				}
				if currentY != lastY {
					continue
				} else {
					if currentX < lastX && wireACurrentX < lastX && currentX < wireACurrentX {
						Xings = append(Xings, Point{
							X: wireACurrentX,
							Y: currentY,
						})
					} else if lastX < currentX && wireACurrentX < currentX && lastX < wireACurrentX {
						Xings = append(Xings, Point{
							X: wireACurrentX,
							Y: currentY,
						})
					}
				}
			}
		}
	}
	return Xings
}

func routeWire(wire map[string][]int, index, xPos, yPos int) (int, int) { // wire = map to hold plot of the wireA and wireB
	cords := fmt.Sprintf("%v,%v", xPos, yPos)
	wire[cords] = append(wire[cords], index+1)
	return xPos, yPos
}

const (
	R = "R"
	L = "L"
	D = "D"
	U = "U"
)
