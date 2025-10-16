package main

import "fmt"

func changeSlice(arr []int) []int {
	arr[0] = 50
	arr = append(arr, 100)
	return arr
}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	arr = append(arr, 6)
	arr = append(arr, 7)

	x := arr[4:]
	y := changeSlice(x)

	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(len(y))
	fmt.Println(cap(y))
}