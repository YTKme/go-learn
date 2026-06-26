package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	// var colorArray = [3]string{"Red", "Green", "Blue"}
	// This is a slice
	// var colorArray = []string{"Red", "Green", "Blue"}
	//
	var colorArray = make([]string, 0, 3)

	fmt.Println("Slice")

	colorArray = append(colorArray, "Red", "Green", "Blue")
	fmt.Println("Color Array:", colorArray)

	colorArray = append(colorArray, "Purple", "Fuchsia")
	fmt.Println("Color Array:", colorArray)

	colorArray = remove(colorArray, 1)
	fmt.Println("Color Array After Removal:", colorArray)

	sort.Strings(colorArray)
	fmt.Println("Color Array After Sorting:", colorArray)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
