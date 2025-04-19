package main

import "fmt"

/*
   ### Slice ###
   A slice is a dynamically-sized, flexible view into the elements of an array.
*/

func main() {
	arr := [6]string{"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // Slice - pointer, length, capacity || pointer = arr[1]th address | length = 3 | capacity = 5
	// s = [is a go]
	printDetails(&s)

	s1 := s[1:3] // Slice || pointer = arr[2]th address | length = 2 | capacity = 4
	// s1 = [a Go]
	printDetails(&s1)

	s2 := []int{1, 2, 4, 8} // Slice Literal || pointer = ? | length = 4 | capacity = 4
	// s2 = [1 2 4 8]
	printDetailsForInt(&s2)

	s3 := make([]int, 5) // Slice with make func || pointer = ? | length = 5 | capacity = 5
	// s3 = [0 0 0 0 0]
	printDetailsForInt(&s3)

	s4 := make([]int, 3, 5) // Slice with make func with cap || pointer = ? | length = 3 | capacity = 5
	// s4 = [0 0 0]
	printDetailsForInt(&s4)
}

func printDetails(arr *[]string) {
	fmt.Println("Slice is:", *arr, "|| Pointer =", &(*arr)[0], "| Length =", len(*arr), "| Capacity =", cap(*arr))
}

func printDetailsForInt(arr *[]int) {
	fmt.Println("Slice is:", *arr, "|| Pointer =", &(*arr)[0], "| Length =", len(*arr), "| Capacity =", cap(*arr))
}
