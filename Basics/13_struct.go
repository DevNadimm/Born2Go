package main

import "fmt"

type Student struct {
    Name string
    Age int
    Dob string
}

func NewStudent (name string, age int, dob string) Student {
    return Student{
        Name: name,
        Age: age,
        Dob: dob,
    }
}

func ViewDetails (student Student) {
    fmt.Println("Student Name: ", student.Name)
    fmt.Println("Age: ", student.Age)
    fmt.Println("Date of Birth: ", student.Dob)
}

func main () {
    s1 := NewStudent("Nadim", 21, "2004-05-17")
    ViewDetails(s1);
}