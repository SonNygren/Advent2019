package day3

type Point struct {
	X              int
	Y              int
	LenghtTraveled int
}
type Wire struct {
	Travaled map[Point]int
	Points   map[int]Point
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
