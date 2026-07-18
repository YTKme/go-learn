package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

func main() {
	fmt.Println("Code Challenge: Read A Shopping Cart From JSON")

	jsonString :=
		`[{"name":"apple","price":2.99,"quantity":3},
		{"name":"orange","price":1.50,"quantity":8},
		{"name":"banana","price":0.49,"quantity":12}]`

	result := getCartFromJson(jsonString)

	for _, cartItem := range result {
		fmt.Printf("Item: %v, Price: $%v, Quantity: %v\n", cartItem.Name, cartItem.Price, cartItem.Quantity)
	}
}

func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem

	decoder := json.NewDecoder(strings.NewReader(jsonString))

	_, err := decoder.Token()
	checkError(err)

	var item CartItem

	for decoder.More() {
		err := decoder.Decode(&item)
		checkError(err)
		cart = append(cart, item)
	}

	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
