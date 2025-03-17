package main

import "fmt"

func main() {
	num := 2

	switch num {
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		default: fmt.Println("Unknown")
	}

	switch day := "Monday"; day {
	case "Saturday", "Sunday": fmt.Println("Weekend")
	default: fmt.Println("Weekday")
	}

	score := 85 

	switch {
	case score >= 90:
		fmt.Println("A") 
	case score >= 80:
		fmt.Println("B") 
	case score >= 70:
		fmt.Println("C") 
	default: fmt.Println("Fail")
	}

	// fall through

	number := 1

	switch number {
	case 1: 
		fmt.Println("One")
		fallthrough
	case 2: 
		fmt.Println("Two")
	default: 
		fmt.Println("Unknown")
	}
}