package main

import "fmt"

func add(a int, b int) int {
	var res int
	res = a + b
	return res
}

func main() {
	a := 10;
	res := add(a, 10)
	fmt.Println(res)
}