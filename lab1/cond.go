package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Write num: ")
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Positive")
	} else if num < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Your num is zero")
	}
}
