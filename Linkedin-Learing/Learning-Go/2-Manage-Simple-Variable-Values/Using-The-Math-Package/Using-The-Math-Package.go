package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100

	fmt.Printf("The Sum Is Now: %v\n\n", floatSum)

	fmt.Println("The Value of PI:", math.Pi)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)
}
