package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum)

	total := float64(i1) + f2
	fmt.Println("Total:", total)

	difference := float64(i1) - f2
	fmt.Println("Difference:", difference)

	product := float64(i1) * f2
	fmt.Println("Product:", product)

	quotient := float64(i1) / f2
	fmt.Println("Quotient:", quotient)

	remainder := i1 % 5
	fmt.Println("Remainder:", remainder)
}
