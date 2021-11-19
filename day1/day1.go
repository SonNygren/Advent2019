package day1

import (
	"advent2019/reader"
	"fmt"
)

func Run() error {
	return nil
}

func Part1() (int, error) {
	input, err := reader.ParseAsInt("day1/input")
	if err != nil {
		return 0, err
	}
	result := 0
	for _, mass := range input {
		result += (mass / 3) - 2
	}

	fmt.Printf("result: %v\n", result)
	return result, nil
}

func Part2() error {
	input, err := reader.ParseAsInt("day1/input")
	if err != nil {
		return err
	}
	fmt.Printf("len(input): %v\n", len(input))
	output := 0
	for _, mass := range input {

		for {
			mass = (mass / 3) - 2

			if mass <= 0 {
				break
			}
			output += mass
		}
	}
	fmt.Printf("output: %v\n", output)
	return nil
}

func Part3() error {
	input, err := reader.ParseAsInt("day1/input")
	if err != nil {
		return err
	}

	output := 0
	for _, mass := range input {
		output += ropa_på_sig_själv(mass, 0)
	}
	fmt.Printf("result: %v\n", output)
	return nil
}

func ropa_på_sig_själv(mass int, output int) int {
	mass = (mass / 3) - 2
	if mass <= 0 {
		return output
	}
	return ropa_på_sig_själv(mass, output+mass)
}

func ropa_själv_MedPointer(mass int, output *int) {

	//kalla på sig själv
	//en condition för att komma ut

	mass = (mass / 3) - 2

	fmt.Printf("mass: %v\n", mass)
	if mass <= 0 {
		return
	}
	*output += mass
	ropa_själv_MedPointer(mass, output)

}
