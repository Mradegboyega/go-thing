package main

import (
	"fmt"
)

func main() {
	fmt.Println("struct foundation in Go")
	// no inheritance in Go; No super or parent

	oluwatoyin := User{"oluwatoyin", "oluwatoyin@domain.com", true, 27}
	fmt.Println(oluwatoyin)
	fmt.Printf("Oluwatoyin details are; %+v", oluwatoyin)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
