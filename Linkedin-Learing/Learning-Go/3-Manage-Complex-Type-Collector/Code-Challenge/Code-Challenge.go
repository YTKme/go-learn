package main

import (
	"fmt"
)

const showExpectedResult = false
const showHint = false

func main() {
	fmt.Println("Code Challenge: Convert Slices of Simple Values to a Slice of Objects")

	colorNameSlice := []string{"Red", "Green", "Blue"}
	hexValueSlice := []int{0xFF0000, 0x00FF00, 0x0000FF}
	colorSlice := sliceToObject(colorNameSlice, hexValueSlice)

	fmt.Println("Color:", colorSlice)
}

func sliceToObject(colorNameSlice []string, hexValueSlice []int) []Color {
	colorSlice := make([]Color, 0, len(colorNameSlice))
	for i := 0; i < len(colorNameSlice) && i < len(hexValueSlice); i++ {
		colorSlice = append(
			colorSlice,
			Color{
				Name: colorNameSlice[i],
				Hex:  hexValueSlice[i],
			},
		)
	}

	return colorSlice
}

type Color struct {
	Name string
	Hex  int
}
