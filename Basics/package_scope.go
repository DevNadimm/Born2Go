package main

import (
	"fmt"

	"example.com/Basics/mathlib"
)

// go mod init example.com

func main() {
	fmt.Println("Hello, This is Package Scope!")

	// Access the Basics/mathlib/math.go functions
	var result int = mathlib.Add(1, 2)
	fmt.Println("Addition of 1 and 2 is:", result)
}
