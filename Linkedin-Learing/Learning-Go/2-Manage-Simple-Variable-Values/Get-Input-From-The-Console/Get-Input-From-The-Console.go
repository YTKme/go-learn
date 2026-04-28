package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("String:", text)

	fmt.Print("Enter Number: ")
	str, _ := reader.ReadString('\n')
	number, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Value of Number:", number)
	}
}
