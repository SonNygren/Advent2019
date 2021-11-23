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

	wires := []Wire{}
	for _, path := range wirePaths {
		wire, err := TraverseWire(path)
		if err != nil {
			return err
		}
		wires = append(wires, wire)
	}

	res := GetPositions(wires[0], wires[1])

	shortest := res[0].LenghtTraveled

	for _, crossCoords := range res {
		fmt.Printf("crossCoords: %v\n", crossCoords)
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

func TraverseWire(path []string) (Wire, error) {
	wire := Wire{
		Travaled: map[Point]int{},
		Points:   map[int]Point{},
	}
	startPoint := Point{0, 0, 0}
	wire.Points[0] = startPoint

	xPos := 0
	yPos := 0
	traveled := 0

	for index, value := range path {
		direction := string(value[0])
		movement, err := strconv.Atoi(value[1:])
		if err != nil {
			return Wire{}, fmt.Errorf("failed to parse movement Err: %v", err.Error())
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
		point := Point{xPos, yPos, traveled}
		wire.Travaled[point] = traveled
		wire.Points[index+1] = point
	}
	return wire, nil
}

func GetPositions(wireA, wireB Wire) []Point {
	Xings := []Point{}
	for i := 1; i < len(wireA.Points); i++ {
		lineA := Line{
			Travaled: wireA.Travaled,
			Start:    wireA.Points[i-1],
			End:      wireA.Points[i],
		}

		for j := 1; j < len(wireB.Points); j++ {
			lineB := Line{
				Travaled: wireB.Travaled,
				Start:    wireB.Points[j-1],
				End:      wireB.Points[j],
			}

			Xings = append(Xings, append(
				findCrossing(lineA, lineB),
				findCrossing(lineB, lineA)...,
			)...)
		}
	}
	return Xings
}

//     .          bLastB
// -----------    aCurrentA
//     .          bCurrentB
func findCrossing(lineA, lineB Line) []Point {
	if hasSameAngle(lineA, lineB) {
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

	smallestA := getSmallest(lineA)
	smallestB := getSmallest(lineB)
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

func getSmallest(line Line) Point {
	if line.Travaled[line.Start] < line.Travaled[line.End] {
		return line.Start
	}
	return line.End

}

// check for crossing in starting-point which shouldnt be taken into account
func hasSameAngle(lineA, lineB Line) bool {
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
