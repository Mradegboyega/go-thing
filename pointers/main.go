package main

import "fmt"

func main() {
	fmt.Println("Trying out pointers in Go...")

	myNumber := 24

	var ptr = &myNumber

	fmt.Println("Value of pointer: ", ptr)
	fmt.Println("Value of pointer: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value: ", *ptr)

}
