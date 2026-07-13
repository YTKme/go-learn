package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Run A Function As A Goroutine")

	go saySomething("Hello From The Goroutine!")
	fmt.Println("Hello From The Main!")

	go func(message string) {
		fmt.Println("Anonymous Message: ", message)
	}("Hello From The Anonymous Function!")

	time.Sleep(2 * time.Second)
	fmt.Println("All Done!")
}

func saySomething(message string) {
	time.Sleep(1 * time.Second)

	fmt.Println("Message: ", message)
}
