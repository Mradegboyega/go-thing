package main

import "fmt"

func main() {
	fmt.Println("arrays in Go!")

	var fruit [5]string

	fruit[0] = "apple"
	fruit[1] = "banana"
	fruit[2] = "orange"
	fruit[3] = "mango"
	fruit[4] = "kiwi"

	fmt.Println(fruit)
	fmt.Println(len(fruit))
	fmt.Println(cap(fruit))
	fmt.Println(fruit[0])

	var vegList = [3]string{"ila", "ewedu", "efo"}
	fmt.Println(vegList)
}
