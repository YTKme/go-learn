package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

type Tour struct {
	Name  string
	Price string
}

func main() {
	fmt.Println("Parse JSON Formatted Text")

	content := readHttpContent()
	// fmt.Println("Content: ", content)

	tours := tourFromJson(content)

	for _, tour := range tours {
		price, _ := strconv.ParseFloat(tour.Price, 16)
		fmt.Printf("Tour: %s, Price: $%v\n", tour.Name, price)
	}

}

func tourFromJson(content string) []Tour {
	tourSlice := make([]Tour, 0)
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	checkError(err)

	var tour Tour

	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tourSlice = append(tourSlice, tour)
	}

	return tourSlice
}

func readHttpContent() string {
	fmt.Println("Network Request")

	httpClient := http.Client{}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	checkError(err)
	// Need `User-Agent`
	request.Header.Set("User-Agent", "Learning-Go-Client")

	response, err := httpClient.Do(request)
	checkError(err)

	defer response.Body.Close()

	fmt.Printf("Response Type: %T\n", response)

	byteArray, err := io.ReadAll(response.Body)
	checkError(err)
	content := string(byteArray)
	// fmt.Println("Content: ", content)

	return content
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
