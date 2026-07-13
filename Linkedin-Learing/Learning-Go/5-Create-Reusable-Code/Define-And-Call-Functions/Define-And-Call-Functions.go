package main

import (
	"fmt"
)

func main() {
	fmt.Println("Define And Call Functions")

	doSomething()

	value1 := 5
	value2 := 10
	value3 := 56
	sum := addValue(value1, value2)
	fmt.Printf("The sum of %v and %v is %v.\n", value1, value2, sum)

	sumAll := addAllValue(value1, value2, value3)
	fmt.Printf("The sum of all value is %v.\n", sumAll)

	sum, count, average := addMoreValue(value1, value2, value3)
	fmt.Printf("The sum of all value is %v, count is %v and average is %v.\n", sum, count, average)
}

func doSomething() {
	fmt.Println("Doing Something")
}
func addValue(value1, value2 int) int {
	return value1 + value2
}

func addAllValue(valueSlice ...int) int {
	sum := 0

	for _, value := range valueSlice {
		sum += value
	}

	return sum
}

func addMoreValue(valueSlice ...int) (int, int, float64) {
	sum := 0

	for _, value := range valueSlice {
		sum += value
	}

	count := len(valueSlice)
	average := float64(sum) / float64(count)

	return sum, count, average
}
