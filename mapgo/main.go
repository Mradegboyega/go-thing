package main

import "fmt"

func main() {
	fmt.Println("maps in Go!")

	languages := make(map[string]string)
	languages["Go"] = "https://golang.org/"
	languages["Rust"] = "https://www.rust-lang.org/"
	languages["Python"] = "https://www.python.org/"
	languages["C#"] = "CSharp"
	languages["Java"] = "https://www.java.com/"
	languages["C++"] = "https://www.cplusplus.com/"
	languages["C"] = "https://www.cprogramming.com/"
	//fmt.Println(languages)
	//fmt.Println(languages["Go"])

	for key, value := range languages {
		fmt.Printf("for key %v: %v\n", key, value)
	}
}
