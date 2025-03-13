package main

import "fmt"

func main() {
	var studentId int = 123942 
	var studentName string = "Nadim Chowdhury"
	var studentMajor string = "CSE"
	var studentAge int = 21

	fmt.Println("\n")
	fmt.Println("Student ID:", studentId)
	fmt.Println("Student Name:", studentName)
	fmt.Println("Student Major:", studentMajor)
	fmt.Println("Student Age:", studentAge)

	// Alternative
	sId := 200384
	sName := "Nadim"

	fmt.Printf("\n\n")
	fmt.Printf("Student ID: %d\n", sId)
	fmt.Printf("Student Name: %s\n", sName)
	
	// Variable update
	sName = "Md. Nadimul Haque Chowdhury"
	
	fmt.Printf("\n\n")
	fmt.Println("Unpdated Name:", sName)
	fmt.Printf("\n\n")
}
