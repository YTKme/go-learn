package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Code Challenge: Create a More Advanced Calculator App")

	value1 := "10"
	value2 := "5.5"
	operation := "+"
	result := calculate(value1, value2, operation)
	fmt.Printf("Result: %v\n", result)
}

func calculate(input1 string, input2 string, operation string) float64 {
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "+":
		return addValues(value1, value2)
	case "-":
		return subtractValues(value1, value2)
	case "*":
		return multiplyValues(value1, value2)
	default:
		fmt.Println("Invalid Operation")
		return 0
	}
}

func convertInputToValue(input string) float64 {
	inputValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error: ", err)
		return 0
	}
	return inputValue
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}
