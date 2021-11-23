package day3

import (
	"advent2019/reader"
	"fmt"
	"strconv"
	"strings"
)

func Part1() error {
	input, err := reader.ParseAsString("day3/input")
	if err != nil {
		return err
	}

	wirePaths := [][]string{}
	for _, v := range input {
		wirePaths = append(wirePaths, strings.Split(v, ","))
	}

	wires := []map[int]Point{}
	for _, path := range wirePaths {
		wire, err := TraverseWire(path)
		if err != nil {
			return err
		}
		wires = append(wires, wire)
	}

	res := findCrossCoords(wires[0], wires[1])

	shortest := res[0].LenghtTraveled
	for _, crossCoords := range res {
		if crossCoords.LenghtTraveled < shortest {
			shortest = crossCoords.LenghtTraveled
		}
	}
	fmt.Printf("shortest: %v\n", shortest)

	smallest := turnPositive(res[0].X) + turnPositive(res[0].Y)
	for _, p := range res {
		dist := turnPositive(p.X) + turnPositive(p.Y)
		if dist < smallest {
			smallest = dist
		}
	}

	fmt.Printf("smallest: %v\n", smallest)
	return nil
}

func turnPositive(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func TraverseWire(path []string) (map[int]Point, error) {
	wire := map[int]Point{
		0: {0, 0, 0},
	}

	xPos := 0
	yPos := 0
	traveled := 0

	for index, value := range path {
		direction := string(value[0])
		movement, err := strconv.Atoi(value[1:])
		if err != nil {
			return map[int]Point{}, fmt.Errorf("failed to parse movement Err: %v", err.Error())
		}
		switch direction {
		case R:
			xPos += movement
		case L:
			xPos -= movement
		case U:
			yPos += movement
		case D:
			yPos -= movement
		}

		traveled += movement
		wire[index+1] = Point{xPos, yPos, traveled}
	}
	return wire, nil
}

func findCrossCoords(wireA, wireB map[int]Point) []Point {
	Xings := []Point{}
	wireALines := map[int]Line{}
	wireBLines := map[int]Line{}

	for i := 1; i < len(wireA); i++ {
		wireALines[i-1] = Line{
			Start: wireA[i-1],
			End:   wireA[i],
		}
	}

	for i := 1; i < len(wireB); i++ {
		wireBLines[i-1] = Line{
			Start: wireB[i-1],
			End:   wireB[i],
		}

	}

	for i := 0; i < len(wireALines); i++ {
		for j := 0; j < len(wireBLines); j++ {
			Xings = append(Xings, append(
				findCrossing(wireALines[i], wireBLines[j]),
				findCrossing(wireBLines[j], wireALines[i])...,
			)...)
		}
	}

	return Xings
}

//     .          bLastB
// -----------    aCurrentA
//     .          bCurrentB
func findCrossing(lineA, lineB Line) []Point {
	if areParalele(lineA, lineB) {
		return []Point{}
	}

	aStart, aEnd := findBorders(lineA.Start.Y, lineA.End.Y)
	if !isWithinArea(aStart, aEnd, lineB.Start.Y) {
		return []Point{}
	}

	bStart, bEnd := findBorders(lineB.Start.X, lineB.End.X)
	if !isWithinArea(bStart, bEnd, lineA.Start.X) {
		return []Point{}
	}

	smallestA := getSmallest(lineA.Start, lineA.End)
	smallestB := getSmallest(lineB.Start, lineB.End)
	littleA, bigA := findBorders(smallestA.Y, lineB.Start.Y)
	littleB, bigB := findBorders(smallestB.X, lineA.Start.X)
	a := smallestA.LenghtTraveled + bigA - littleA
	b := smallestB.LenghtTraveled + bigB - littleB
	totalLenghtToXing := a + b

	return []Point{{
		LenghtTraveled: totalLenghtToXing,
		X:              lineA.Start.X,
		Y:              lineB.Start.Y,
	}}
}

func getSmallest(start, end Point) Point {
	if start.LenghtTraveled < end.LenghtTraveled {
		return start
	}
	return end
}

// check for crossing in starting-point which shouldnt be taken into account
func areParalele(lineA, lineB Line) bool {
	return isVertical(lineA) && isVertical(lineB) || !isVertical(lineA) && !isVertical(lineB)
}

func isVertical(line Line) bool {
	return line.Start.Y == line.End.Y
}

// returns true if cord is bigger then start and smaller then end
func isWithinArea(start, end, cord int) bool {
	return start < cord && cord < end
}

func findBorders(borderA, borderB int) (int, int) {
	if borderB < borderA {
		return borderB, borderA
	}
	return borderA, borderB
}
