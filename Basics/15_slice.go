package main

import "fmt"

/*
   ### Slice ###
   A slice is a dynamically-sized, flexible view into the elements of an array.
*/

func main() {
	numbers := []int{1, 2, 3}          // Slice declaration with initial values
	numbers = append(numbers, 4, 5, 6) // Appending new elements to the slice

	fmt.Println(numbers) // Output: [1 2 3 4 5 6]

	// Topic: Array to Slice
	marksArr := [5]int{50, 60, 70, 80, 90}
	marksSlice := marksArr[0:2] // From index 0 up to (but not including) index 2

	fmt.Println(marksSlice) // Output: [50 60]
}
