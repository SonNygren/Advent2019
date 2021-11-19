package reader

import (
	"bufio"
	"os"
	"strconv"
)

func ParseAsString(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input, nil
}

func ParseAsInt(path string) ([]int, error) {
	input, err := ParseAsString(path)
	if err != nil {
		return []int{}, err
	}
	outputArray := []int{}
	for i := range input {
		output, err := strconv.Atoi(input[i])
		if err != nil {
			return []int{}, err
		}
		outputArray = append(outputArray, output)
	}

	return outputArray, nil
}
