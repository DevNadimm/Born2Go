package main 

import "fmt"

func main () {
	var counter = createCounter()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func createCounter() func() int {
	count := 0
	return func() int{
		count++
		return count
	}
}