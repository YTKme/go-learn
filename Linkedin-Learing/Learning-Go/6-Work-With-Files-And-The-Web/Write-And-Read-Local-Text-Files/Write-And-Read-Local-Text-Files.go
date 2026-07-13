package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Write And Read Local Text Files")

	filename := "./fromString.txt"
	fileAlpha, err := os.Create(filename)
	checkError(err)
	defer fileAlpha.Close()

	length, err := io.WriteString(fileAlpha, "Hello From Go!")
	fmt.Printf("Wrote File With %v Character.\n", length)

	readFile(filename)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	checkError(err)
	fmt.Println("Text Read From File:", string(data))
}
