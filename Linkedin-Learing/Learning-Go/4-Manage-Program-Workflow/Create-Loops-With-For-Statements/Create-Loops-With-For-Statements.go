package main

import (
	"fmt"
)

func main() {
	fmt.Println("Create Loops With For Statements")

	colorSlice := []string{"Red", "Green", "Blue"}

	// for i := 0; i < len(colorSlice); i++ {
	// 	fmt.Println("Color:", colorSlice[i])
	// }

	for i := range colorSlice {
		fmt.Println("Color Range:", colorSlice[i])
	}

	for _, color := range colorSlice {
		fmt.Println("Color Value:", color)
	}

	stateSlice := make(map[string]string)
	stateSlice["WA"] = "Washington"
	stateSlice["OR"] = "Oregon"
	stateSlice["CA"] = "California"

	for stateCode, _ := range stateSlice {
		fmt.Println("State Code:", stateSlice[stateCode])
	}

	value := 0
	sum := 0

	for value < 5 {
		sum += value
		fmt.Printf("Value: %v\n", value)
		fmt.Printf("Sum: %v\n", sum)
		value++
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	println("End of Program")
	fmt.Printf("Sum: %v\n", sum)
}
