package main

import "fmt"

// Global Variable
var appName = "Go Learning App"

func showAppName() {
    fmt.Println("Application Name:", appName)
}

func main() {
    // Using Global Variable
    showAppName()
    fmt.Println("Main Function:", appName)

    // Local Scope Example
    sayHello()

    // ❌ Uncommenting below line will cause an error 
    // because `message` is only accessible inside `sayHello()`
    // fmt.Println(message)

    // Same Name: Local & Global Variable
    printMessage()
    fmt.Println("Inside Main:", message) // Global Variable
}

// Local Scope Example
func sayHello() {
    message := "Hello, Go!" // Local Variable
    fmt.Println("Inside sayHello():", message)
}

// Global Variable
var message = "Global Message"

func printMessage() {
    message := "Local Message" // Local Variable
    fmt.Println("Inside Function:", message)
}
