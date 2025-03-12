package main 

import "fmt"

func main() {
	// 1. Using var (Explicit Type Declaration)
	var name string = "John Doe"
	var age int = 24
	var height float64 = 153.4

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(height)

	// 2. Using var Without Initialization (Default Values)
	var testName string
	var testAge int
	fmt.Println(testName)
	fmt.Println(testAge)

	// 3. Using := (Short Variable Declaration - Implicit Type)
	carName := "BMW"
	carMileage := 15.5

	fmt.Println(carName)
	fmt.Println(carMileage)

	// 4. Multiple Variable Declaration
	var a, b, c int = 1, 2, 3
	x, y, z := 5, 55, 555
	fmt.Println(a, b, c, x, y, z)

	// 5. Declaring Multiple Variables Using a Block
	var (
		firstName string = "John"
		lastName string = "Doe"
	)
	fmt.Println(firstName, lastName)

	// Constants
	const pi = 3.142

	fmt.Println(pi)

	// Edge Cases
	collegeName := "Harvard"
	// collegeName := "MIT" // This will throw the error
	fmt.Println(collegeName)

	const gravity = 9.8
	// gravity = 9.9 // Cannot update the const variable

	


}