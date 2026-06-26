package main

import (
	"fmt"
)

func main() {
	fmt.Println("Group Related Values In Structs")

	poodle := Dog{
		Breed:  "Poodle",
		Weight: 34,
	}
	fmt.Println("Poodle:", poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)

	poodle.Weight = 29
	fmt.Println("Poodle:", poodle)
}

type Dog struct {
	Breed  string
	Weight int
}
