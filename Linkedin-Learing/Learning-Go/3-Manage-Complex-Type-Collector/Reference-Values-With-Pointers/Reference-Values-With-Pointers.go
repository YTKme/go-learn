package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p *int = &anInt

	fmt.Println("Pointers")

	if p == nil {
		fmt.Println("Value of p: nil")
	} else {
		fmt.Printf("Value of p: %v\n", *p)
	}

	value1 := 42.13
	pointer1 := &value1
	*pointer1 = *pointer1 / 31

	fmt.Printf("Pointer: %v\n", *pointer1)
	fmt.Printf("Value 1: %v\n", value1)
}
