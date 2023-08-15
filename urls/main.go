package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=uetghjdydwnbdw"

func main() {
	fmt.Println("handling url in Go!")
	fmt.Println(myurl)

	//parsing url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Println(qparams.Get("paymentid"))
	fmt.Println(qparams.Get("coursename"))

	//url anatomy

	partsOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev:3000",
		Path:    "/learn",
		RawPath: "coursename=reactjs&paymentid=uetghjdydwnbdw",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
