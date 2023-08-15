package main

import (
	"fmt"
)

func main() {
	fmt.Println("struct foundation in Go")
	// no inheritance in Go; No super or parent

	oluwatoyin := User{"oluwatoyin", "oluwatoyin@domain.com", true, 27}
	fmt.Println(oluwatoyin)
	fmt.Printf("Oluwatoyin details are; %+v\n", oluwatoyin)

	oluwatoyin.NewMail()
	oluwatoyin.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "oluwatoyin@domain.dev"
	fmt.Println("User email is: ", u.Email)
}
