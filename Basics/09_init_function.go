package main

import "fmt"

func main() {
	fmt.Println("Hello, welcome to Go!")
}

func init() {
	fmt.Println("Hello, I am the init function and I will execute first.")
}
