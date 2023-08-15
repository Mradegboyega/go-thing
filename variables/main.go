package main

import "fmt"

var LoginToken string = "jeejkfwetyhewnmwyff"

func main() {
	var username string = "mradegboyega"
	fmt.Println("username")
	fmt.Println("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println("isLogggedIn")
	fmt.Println("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 254
	fmt.Println("smallVal")
	fmt.Println("smallVal is of type: %T \n", smallVal)

	var smallFloat float64 = 255.456753786
	fmt.Println("smallFloat")
	fmt.Println("smallFloat is of type: %T \n", smallFloat)

	fmt.Println(LoginToken)

	anotherUsername := "Oluwatoyin"
	fmt.Println(anotherUsername)
}
