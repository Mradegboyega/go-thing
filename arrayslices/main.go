package main

import "fmt"

func main() {
	//fmt.Println("array slices in Go!")

	var country = []string{"US", "GB", "NG", "FRA", "GH"}
	//fmt.Println("example of country: ", country)

	country = append(country, "RW", "MX")
	//fmt.Println(country)

	country = append(country[:4])
	//fmt.Println(country)

	highScore := make([]int, 4)
	highScore[0] = 65
	highScore[1] = 76
	highScore[2] = 79
	highScore[3] = 81

	highScore = append(highScore, 73, 80, 89)

	//fmt.Println(highScore)

	//using index to remove values from slice

	var courses = []string{"javascript", "ruby", "java", "python", "C", "rust", "C++"}

	fmt.Println(courses)
	var index int = 4

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
