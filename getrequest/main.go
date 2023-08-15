package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	performGetRequest()
}

func performGetRequest() {
	const myurl = "http://"

	response, err := http.Get(myurl)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("content length: ", response.ContentLength)

	//another way of handling http requests

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	byteCount, err := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("responseString is: ", responseString.String())
}
