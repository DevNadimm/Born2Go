package main

import "fmt"

/*
   ### Slice ###
   A slice is a dynamically-sized, flexible view into the elements of an array.
*/

func main() {
	arr := [6]string {"This", "is", "a", "Go", "interview", "question"}
	fmt.Println(arr)

	s := arr[1:4] // Slice - pointer, length, capacity || pointer = arr[1]th address | length = 3 | capacity = 5
	// s = [is a go]
	fmt.Println(s)

	s1 := s[1:3] // pointer = arr[2]th address | length = 2 | capacity = 4
	// s1 = [a Go]
	fmt.Println(s1)
}
