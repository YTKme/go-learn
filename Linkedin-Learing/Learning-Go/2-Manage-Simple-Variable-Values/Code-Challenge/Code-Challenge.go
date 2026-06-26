package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(value1 string, value2 string) float64 {
	// Parse value1 from string to float64
	number1, _ := strconv.ParseFloat(strings.TrimSpace(value1), 64)

	// Parse value2 from string to float64
	number2, _ := strconv.ParseFloat(strings.TrimSpace(value2), 64)

	// Calculate Result
	return number1 + number2
}

func main() {
	fmt.Printf("Code Challenge: Simple Calculator\n")

	value1 := "  10  "
	value2 := "  5.5  "

	result := calculate(value1, value2)
	fmt.Printf("Result: %.2f\n", result)
}
