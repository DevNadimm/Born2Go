package main

import "fmt"

// Topic - Anonymous function and IIFE (Immediately Invoked Function Expression)

func main() {

	// Anonymous function - A function that has no name.
	func(a, b int) {
		sum := a + b
		fmt.Println(sum)
	}(10, 20) // IIFE

	/*
	   NOTE -
	   An anonymous function returns an error unless it is invoked immediately.
	*/
}
