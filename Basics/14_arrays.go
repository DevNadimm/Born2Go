package main

import "fmt"

func main () {
	// var marks = [6]int{14, 76, 45, 95, 34, 23}
	marks := [6]int {14, 76, 45, 95, 34, 23} // marks := [...]int {14, 76, 45, 95, 34, 23} --> Go count the length automatically

	fmt.Println(marks)

	marks[0] = 99
	fmt.Println(marks)
}