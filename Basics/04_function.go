package main

import "fmt"

func main() {
	a := 10
	b := 20
	year := 2025

	checkLeapYear(year)
	printSum(a, b)
	var sum, difference int = multiReturn(10, 20) // Alternative - sum, difference := multiReturn(10, 20)
	fmt.Printf("Sum is: %d, and Difference is: %d\n",sum, difference)
}

func checkLeapYear(year int) {
	if isLeapYear(year) {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is not a leap year\n", year)
	}
}

func printSum(a, b int) {
	result := add(a, b)
	fmt.Println("Sum:", result)
}

func add(num1, num2 int) int {
	return num1 + num2
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// Multi return function
func multiReturn(a, b int) (int, int) {
	return a+b, a-b
}
