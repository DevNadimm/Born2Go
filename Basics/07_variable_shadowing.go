package main

import "fmt"

var a int = 10;  // Global variable

func main() {
    age := 10
    if age >= 10 {
        a := 20  // This 'a' is a new local variable (shadowing global 'a')
        fmt.Println(a)  // Prints 20 (local 'a')
    }
    fmt.Println(a)  // Prints 10 (global 'a' remains unchanged)
}
