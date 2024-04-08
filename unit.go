package main

import "fmt"

func main() {
	fmt.Println(Max([]int{4, 45, 48, 98, 5}))
}

func Max(numbers []int) int {
	var max int

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func IndexMax(numbers []int) int {
	var max int
	var index int

	for i, number := range numbers {
		if number > max {
			max = number
			index = i
		}
	}
	return index
}
