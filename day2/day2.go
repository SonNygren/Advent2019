package day5

import (
	"advent2019/reader"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() error {
	input, err := reader.ParseAsInt("input")
	if err != nil {
		return err
	}
	run(input)
	return nil
}

func run(input []int) {

	program := map[int]int{}
	for i, v := range input {
		program[i] = v
	}

	for adress := 0; adress < len(input); {
		checkNumbers, _ := modes(adress, program)
		switch checkNumbers {
		case ADD:
			add(adress, program)
			//	if checkRes(add(adress, program)) {
			//		fmt.Printf("noun: %v\n", noun)
			//		fmt.Printf("verb: %v\n\n", verb)
			//		fmt.Printf("(100*noun + verb): %v\n", (100*noun + verb))
			//	}
			adress += 2
		case MULTI:
			multiply(adress, program)
			//	if checkRes(multiply(adress, program)) {
			//		fmt.Printf("noun: %v\n", noun)
			//		fmt.Printf("verb: %v\n\n", verb)
			//		fmt.Printf("(100*noun + verb): %v\n", (100*noun + verb))
			//}
			adress += 2
		case THREE:
			output := _three(userInput(), adress, program)
			fmt.Printf("output_THREE: %v\n", output)

			adress += 2
		case FOUR:
			output := _four(adress, program)
			fmt.Printf("output_FOUR: %v\n", output)
			adress += 2
		case QUIT:
			//fmt.Println(program[0])
			return
		}
	}
}

func userInput() int {
	fmt.Println("Please enter a number:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result, _ := strconv.Atoi(scanner.Text())

	return result
}

const (
	ADD   = 1
	MULTI = 2
	THREE = 3
	FOUR  = 4
	QUIT  = 99
)

func modes(adress int, program map[int]int) (int, error) {

	strconvert := strconv.Itoa(program[adress])
	strSplit := strings.SplitAfter(strconvert, "")
	//fmt.Printf("strSplit: %v\n", strSplit)

	for i := 0; i < len(strSplit); i++ {

		if len(strSplit) <= 4 { // titta på sista siffran
			return strconv.Atoi((strSplit[len(strSplit)-1]))
		}

	}
	return program[adress], nil
}

func _three(input int, adress int, program map[int]int) int {

	valueOfIndex := program[adress+1]
	program[valueOfIndex] = input

	return program[valueOfIndex]
}

func _four(adress int, program map[int]int) int {

	valueOfIndex := program[adress+1]
	result := program[valueOfIndex]
	return result

}

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

func checkRes(res, a, b int) bool {
	if res == 19690720 {
		fmt.Printf("a: %v\n", a)
		fmt.Printf("b: %v\n", b)
		return true
	}
	return false
}

//
//func Part2() error {
//	input, err := reader.ParseAsInt("day2/input")
//	if err != nil {
//		return err
//	}
//
//	for noun := 0; noun <= 99; noun++ {
//		for verb := 0; verb <= 99; verb++ {
//			run(noun, verb, input)
//		}
//	}
//	return err
//}
//
