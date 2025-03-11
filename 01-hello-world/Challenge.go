package main

import "fmt"

func main() {
	// Basic Printing: Write a program that prints your name and a greeting on separate lines.
	fmt.Println("Hello Anirudha!")
	fmt.Println("Welcome to learn Go🚀")

	// Formatted Output: Use Printf to print your name and age like: "My name is John and I am 25 years old."
	fmt.Printf("My name is %s and I am %d years old. \n", "Anirudha", 25)

	// Multi-line String: Print a multi-line string using a single Println.
	fmt.Println(`This is one way of printing
	on multiline`)

	// Experiment with Unicode: Print a message containing emojis and non-English characters.
	fmt.Print("ನಮ್ಮ ಕರ್ನಾಟಕ 🚀")
}