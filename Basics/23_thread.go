package main

import (
	"fmt"
	"time"
)

func work1() {
	for i := 1; i <= 3; i++ {
		fmt.Println("🧵 Thread 1 working:", i)
		time.Sleep(1 * time.Second)
	}
}

func work2() {
	for i := 1; i <= 3; i++ {
		fmt.Println("🧵 Thread 2 working:", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go work1() // Goroutine 1 (Thread 1)
	go work2() // Goroutine 2 (Thread 2)
	time.Sleep(4 * time.Second)
}
