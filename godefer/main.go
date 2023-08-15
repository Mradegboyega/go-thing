package main

import "fmt"

func main() {
	defer fmt.Println("defer sequence in Go!")
	defer fmt.Println("I will make it")
	defer fmt.Println("Last in first out")
	fmt.Println("loving Go")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
