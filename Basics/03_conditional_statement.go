package main

import "fmt"

func main() {
	age := 20
	// if-else-if statement
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a kid")
	}

	// Nested if statement
	age = 20
	if age >= 18 {
		if age >= 21 {
			fmt.Println("You are eligible to vote and drink")
		} else {
			fmt.Println("You are eligible to vote but not to drink")
		}
	}

	// Short statement
	if a := 10; a > 5 {
		fmt.Println("a is greater than 5")
	}

	// Short statement with else
	if a := 3; a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}

	// Switch-case statement
	day := 3
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wednesday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")
	}

	// Additional switch case example
	dayOfWeek := "Monday"
	switch dayOfWeek {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("End of the work week")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Midweek day")
	}
}
