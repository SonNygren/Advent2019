package day3

type Point struct {
	X              int
	Y              int
	LenghtTraveled int
}

type Line struct {
	Travaled map[Point]int
	Start    Point
	End      Point
}

const (
	R = "R"
	L = "L"
	D = "D"
	U = "U"
)
