package main

import (
	"fmt"
)

func algorithm(input int) []int {
	var result []int

	for input > 1 {
		result = append(result, input)
		input /= 2
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
	input := 9
	output := algorithm(input)

	fmt.Println("Input:")
	fmt.Println(input)
	fmt.Println("Output:")
	for _, value := range output {
		fmt.Println(value)
	}
}
