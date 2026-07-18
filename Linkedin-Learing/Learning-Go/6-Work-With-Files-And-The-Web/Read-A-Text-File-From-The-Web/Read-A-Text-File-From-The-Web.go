package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Read A Text File From The Web")

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
	fmt.Println("Content: ", content)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
