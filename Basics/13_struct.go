package main

import "fmt"

type User struct {
	name string
	age  int
	role string
}

// Method for the User struct (outside the struct definition) -- called as receiver function
func (u User) greet() {
	fmt.Printf("Hello, my name is %s and I'm an %s.\n", u.name, u.role)
}

func main() {
	user := User{
		name: "Nadim Chy",
		age:  20,
		role: "admin",
	}

	fmt.Println("User Name:", user.name, "\nAge:", user.age, "\nRole:", user.role)

	user.greet()
}
