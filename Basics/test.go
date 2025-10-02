package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main () {
    c1 := counter()
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2
}

func init () {
    fmt.Println("Hello GO!")
}