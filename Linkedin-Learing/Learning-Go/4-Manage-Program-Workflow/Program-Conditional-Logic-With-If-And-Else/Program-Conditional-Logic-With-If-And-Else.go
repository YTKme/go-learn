package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program Conditional Logic With If And Else")

	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less Than Zero"
	} else if theAnswer == 0 {
		result = "Equal To Zero"
	} else {
		result = "Greater Than Zero"
	}

	fmt.Println("The Answer:", theAnswer)
	fmt.Println("Result:", result)
}
