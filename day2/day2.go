package day5

import (
	"advent2019/reader"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Part1() error {
	input, err := reader.ParseAsInt("day2/inputday5")
	if err != nil {
		return err
	}

	return run(input)
}

func run(input []int) error {

	program := map[int]int{}
	for i, v := range input {
		program[i] = v
	}

	for adress := 0; adress < len(input); {

		codes := modes(fmt.Sprintf("%v", program[adress]))
		code, _ := strconv.Atoi(codes[len(codes)-2:])
		mode1, _ := strconv.Atoi(codes[len(codes)-3 : len(codes)-2])
		mode2, _ := strconv.Atoi(codes[len(codes)-4 : len(codes)-3])

		modes := map[int]int{
			1: mode1,
			2: mode2,
		}

		switch code {
		case ADD:
			add(adress, program, modes)
			adress += 4
		case MULTI:
			multiply(adress, program, modes)
			adress += 4
		case USERINPUT:
			err := userInput(adress, program)
			if err != nil {
				return err
			}
			adress += 2
		case OUTPUT:
			printValue(adress, program)
			adress += 2
		case QUIT:
			return nil
		}
	}
	return nil
}

func userInput(adress int, program map[int]int) error {
	fmt.Println("Please enter a number:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	result, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return err
	}

	valueOfIndex := program[adress+1]
	program[valueOfIndex] = result

	return nil
}

const (
	ADD       = 1
	MULTI     = 2
	USERINPUT = 3
	OUTPUT    = 4
	QUIT      = 99
)

const (
	INDIRECT = 0
	DIRECT   = 1
)

func getValue(mode, adress int, program map[int]int) int {
	switch mode {
	case INDIRECT:
		return program[program[adress]]
	case DIRECT:
		return program[adress]
	default:
		return 0
	}
}

func modes(digits string) string {
	if len(digits) < 4 {
		return modes("0" + digits)
	}
	return digits
}

func printValue(adress int, program map[int]int) {
	valueOfIndex := program[adress+1]
	result := program[valueOfIndex]
	fmt.Printf("output: %v\n", result)
}

func add(adress int, program, modes map[int]int) (int, int, int) {
	a := getValue(modes[1], adress+1, program)
	b := getValue(modes[2], adress+2, program)
	res := a + b
	program[program[adress+3]] = res
	return res, a, b
}

func multiply(adress int, program, modes map[int]int) (int, int, int) {

	a := getValue(modes[1], adress+1, program)
	b := getValue(modes[2], adress+2, program)

	res := a * b
	indexC := program[adress+3]
	program[indexC] = res
	return res, a, b
}
