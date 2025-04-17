package main

import "fmt"

func main() {
	// === Pointer Basics ===
	x := 20
	p := &x // p holds the address of x

	fmt.Println("Address of x (p):", p)
	fmt.Println("Value at address p (*p):", *p)

	*p = 30 // Change value via pointer
	fmt.Println("Updated value of x:", x)

	fmt.Println()

	// === Pass by Value ===
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before printNumbers (pass by value):", numbers)
	printNumbers(numbers)

	// === Pass by Reference ===
	printNumbers2(&numbers)
	fmt.Println("After printNumbers2 (pass by reference):", numbers)

	fmt.Println()

	// Show the final state of the numbers array
	printNumbers(numbers)
}

// Pass by value: makes a copy of the array
func printNumbers(numbers [5]int) {
	fmt.Println("Inside printNumbers (value):", numbers)
}

// Pass by reference: modifies the original array
func printNumbers2(numbers *[5]int) {
	fmt.Println("Inside printNumbers2 (reference):", *numbers)
	numbers[0] = 100
}
