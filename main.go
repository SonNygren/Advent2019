package main

import (
	"advent2019/day4"
	"fmt"
)

func main() {

	if err := day4.Part1(); err != nil {
		fmt.Printf("err: %v\n", err.Error())
	}

}

//err := day3.Part1()
//	if err != nil {
//		fmt.Printf("err: %v\n", err.Error())
//	}
