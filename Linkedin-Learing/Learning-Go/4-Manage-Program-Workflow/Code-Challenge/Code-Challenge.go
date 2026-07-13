package main

import (
	"fmt"
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func calculateTotal(cart []CartItem) float64 {
	var total float64 = 0.0

	for _, item := range cart {
		total += item.Price * float64(item.Quantity)
	}

	return total
}

func main() {
	fmt.Println("Code Challenge: Calculate the Value of a Shopping Cart")

	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Printf("Total: %v\n", result)
}
