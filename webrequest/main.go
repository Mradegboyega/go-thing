package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Web request in Go!")

	response, err := http.Get(url)

	//error catch

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	//defer response.Body.Close() which will close the connection after the program has finished

	defer response.Body.Close()

	// read response body

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
