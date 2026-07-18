package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	var colorArray [3]string
	colorArray[0] = "Red"
	colorArray[1] = "Green"
	colorArray[2] = "Blue"

	fmt.Println("Arrays")

	fmt.Println("Color Array:", colorArray)
	fmt.Println("Color 1:", colorArray[0])

	var numberArray = [5]int{5, 3, 1, 2, 4}
	fmt.Println("Number Array:", numberArray)

	fmt.Println("Number of Color:", len(colorArray))
	fmt.Println("Number of Number:", len(numberArray))
}
