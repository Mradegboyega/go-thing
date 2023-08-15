package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time learning in Go")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2023 15:04:05 UTC"))

	createdDate := time.Date(2023, time.February, 8, 7, 0.+time)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2023 Monday"))

}
