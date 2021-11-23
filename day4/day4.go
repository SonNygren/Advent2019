package day4

import "fmt"

func Part1() error {

	//	startInput := 138241
	//	endInput := 674034
	//sliceOfInputs := []string{}

	counterForBoth := 0
	counterForOrder := 0
	for i := 138241; i < 674034; i++ {

		in := fmt.Sprintf("%v", i)
		if increasingOrder(in) {
			counterForOrder++
			//fmt.Printf("i: %v\n", i)
			if isPair(in) {
				counterForBoth++
			}

		}
	}
	fmt.Printf("counterForOrder: %v\n", counterForOrder)
	fmt.Printf("counterForAdjacent: %v\n", counterForBoth)
	return nil
}

func isPair(password string) bool {
	counter := map[int]int{}

	for i := 1; i < len(password); i++ {
		same := int(password[i]) == int(password[i-1])
		if same {
			counter[int(password[i])] += 2
		}
	}

	for _, v := range counter {
		if v == 2 {
			return true
		}
	}
	return false
}

func increasingOrder(password string) bool {
	for i := 1; i < len(password); i++ {
		if int(password[i]) < int(password[i-1]) {
			return false
		}
	}
	return true
}
