package main

import "fmt"

func main(){
	// Numeric Values, String Values, Boolean Values

	fmt.Println(18)
	fmt.Println(5.5)

	fmt.Println("Weather is awesome!")
	fmt.Println("Go is double awesome!")

	fmt.Println(true)
	fmt.Println(false)

	// Edge case 

	// 1. Integer Overflow

	var num uint8 = 255;
	num++;
	fmt.Println(num)

	// 2. Floating-Point Precision Issues

	fmt.Println(0.1 + 0.2)

	// 3. Empty String and Zero Values

	var test string
	var number int

	fmt.Println(test)
	fmt.Println(number)
}