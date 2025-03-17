package main

import "fmt"

func main() {
	// 1️⃣ Basic Switch Statement
	// Write a program that checks a number and prints:

	// "One" if it's 1
	// "Two" if it's 2
	// "Three" if it's 3
	// "Unknown Number" otherwise

	num := 3

	switch num {
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	default: fmt.Println("Unknown Number")
	}

	// 2️⃣ Multiple Case Values
	// Write a switch statement that checks if a given day is a weekday or a weekend.

	switch day := "Sunday"; day {
	case "Saturday", "Sunday": fmt.Println("Weekend🚀")
	default: fmt.Println("Weekday😑")
	}

	// 3️⃣ Grade Calculator (Switch Without Expression)
	// Given a score, print the grade:

	// 90+ → "A"
	// 80-89 → "B"
	// 70-79 → "C"
	// Otherwise → "F"

	grade := 92

	switch {
	case grade >= 90: fmt.Println("Grade A")
	case grade >= 80: fmt.Println("Grade B")
	case grade >= 70: fmt.Println("Grade C")
	default: fmt.Println("Grade F")
	}

	// 4️⃣ Fallthrough Challenge
	// Write a switch statement where:

	// Case 1 prints "One" and falls through.
	// Case 2 prints "Two" and stops.
	// Default prints "Not One or Two".

	number := 1

	switch number {
	case 1: 
		fmt.Println("One")
		fallthrough
	case 2: 
		fmt.Println("Two")
	default:
		fmt.Println("Not One or Two")
	}
	// 5️⃣ Type Switch
	// Create a switch that takes an interface{} value and checks if it's an int, string, or bool.
}