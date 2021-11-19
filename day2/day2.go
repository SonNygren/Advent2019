package day2

import (
	"advent2019/reader"
	"fmt"
)

func Part1() error {
	input, err := reader.ParseAsInt("day2/input")
	if err != nil {
		return err
	}
	run(12, 1, input)
	return nil
}

func run(noun int, verb int, input []int) {
	input[1] = noun
	input[2] = verb

	program := map[int]int{}
	for i, v := range input {
		program[i] = v
	}

	for adress := 0; adress < len(input); {
		switch program[adress] {
		case ADD:
			if checkRes(add(adress, program)) {
				fmt.Printf("noun: %v\n", noun)
				fmt.Printf("verb: %v\n\n", verb)
				fmt.Printf("(100*noun + verb): %v\n", (100*noun + verb))
			}
			adress += 4
		case MULTI:
			if checkRes(multiply(adress, program)) {
				fmt.Printf("noun: %v\n", noun)
				fmt.Printf("verb: %v\n\n", verb)
				fmt.Printf("(100*noun + verb): %v\n", (100*noun + verb))
			}
			adress += 4
		case QUIT:
			//fmt.Println(program[0])
			return
		}
	}
}

func checkRes(res, a, b int) bool {
	if res == 19690720 {
		fmt.Printf("a: %v\n", a)
		fmt.Printf("b: %v\n", b)
		return true
	}
	return false
}

const (
	ADD   = 1
	MULTI = 2
	QUIT  = 99
)

func add(adress int, program map[int]int) (int, int, int) {
	a := program[program[adress+1]]
	b := program[program[adress+2]]
	res := a + b
	program[program[adress+3]] = res
	return res, a, b
}

func multiply(adress int, program map[int]int) (int, int, int) {

	indexA := program[adress+1]
	indexB := program[adress+2]

	a := program[indexA]
	b := program[indexB]

	res := a * b
	indexC := program[adress+3]
	program[indexC] = res
	return res, a, b
}

func Part2() error {
	input, err := reader.ParseAsInt("day2/input")
	if err != nil {
		return err
	}

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			run(noun, verb, input)
		}
	}
	return err
}
