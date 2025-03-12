package main

import "fmt"

func main() {
	// 1. Basic Operations: Print different types of values (integer, float, string, boolean) in a single Println statement.
	fmt.Println("1. \n Integer: ", 5, "\n Float: ", 55.55, "\n String: ", "Learn GO", "\n Boolean: ", true, "\n" )

	// 2. Formatted Output: Use Printf to display a sentence like: "The price of an item is $49.99" with a floating-point value.
	fmt.Printf("2. The price of an item is $%.2f. \n", 49.99)


	// Escape Sequences: Print a string containing new lines, tabs, and quotes.
	fmt.Println("3. Go is best programming \n \t  language in the world  \"🚀\"")

	// Unicode Challenge: Print a sentence in your native language along with an emoji.
	fmt.Println("4. ನನ್ನ ಹೆಸರು ಅನಿರುದ್ಧ. 🚀")
}