package main

import "fmt"

// -----------------------------
// 💡 Function: greet
// This function takes a parameter 'name' and prints a greeting.
// 'name' is the parameter.
// When we call greet("Nadim"), "Nadim" is the argument.
// -----------------------------
func greet(name string) {
	fmt.Println("Hello,", name)
}

// -----------------------------
// 🔹 First-Order Functions
// These functions do NOT take or return other functions.
// -----------------------------

// Adds two integers
func add(a, b int) int {
	return a + b
}

// Multiplies two integers
func mul(a, b int) int {
	return a * b
}

// -----------------------------
// 🔸 Higher-Order Function
// A function that takes another function as a parameter.
// Here, 'op' is a function that takes two ints and returns an int.
// -----------------------------
func operation(x int, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	// 🔹 Parameter vs Argument
	// "Nadim" is the argument passed to the greet function
	greet("Nadim")

	// 🔸 Using Higher-Order Function with 'add'
	sum := operation(2, 4, add)
	fmt.Println("Sum:", sum) // Output: Sum: 6

	// 🔸 Using Higher-Order Function with 'mul'
	product := operation(2, 4, mul)
	fmt.Println("Product:", product) // Output: Product: 8
}
