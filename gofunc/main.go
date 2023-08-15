package main

import "fmt"

func main() {
	fmt.Println("here is functions in Go  !")

	greet()
	goodbye()

	result := add(4, 7)
	fmt.Println("Result is: ", result)

	proRes := proAdd(3, 5, 8, 1)
	fmt.Println("Result is: ", proRes)
}

func greet() {
	fmt.Println("Hello User!")
}

func goodbye() {
	fmt.Println("Goodbye User!")
}

func add(valOne int, valTwo int) int {
	return valOne + valTwo

}

func proAdd(values ...int) int {
	total := 0

	for _, value := range values {
		total += value
	}
	return total
}
