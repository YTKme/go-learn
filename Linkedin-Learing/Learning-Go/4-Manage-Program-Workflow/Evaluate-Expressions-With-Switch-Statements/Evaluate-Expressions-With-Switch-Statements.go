package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Evaluate Expressions With Switch Statements")

	weekday := time.Now().Weekday()
	fmt.Println("Today:", weekday)

	dayNumber := int(weekday)
	fmt.Println("Day Number:", dayNumber)

	var result string

	switch dayNumber = 6; dayNumber {
	case 1:
		result = "It's a Monday"
	case 2:
		result = "It's a Tuesday"
	case 3:
		result = "It's a Wednesday"
	case 4:
		result = "It's a Thursday"
	case 5:
		result = "It's a Friday"
	default:
		result = "It's the Weekend"
	}

	fmt.Println("Result:", result)

	x := 42
	// x := 0
	switch {
	case x < 0:
		result = "Less Than Zero"
		// fallthrough
	case x == 0:
		result = "Equal To Zero"
		// fallthrough
	default:
		result = "Greater Than Zero"
	}

	fmt.Println("X:", x)
	fmt.Println("Result:", result)
}
