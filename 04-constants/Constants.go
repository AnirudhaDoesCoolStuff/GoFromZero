package main

import "fmt"

func main() {
	const pi = 3.142
	const appName = "GoMaster"

	fmt.Println(pi)
	fmt.Println(appName)

	const (
		firstName = "John"
		lastName = "Doe"
	)

	fmt.Println(firstName, lastName)
	
	// You cannot assign y (float64) to an int variable without explicit conversion.

	// const num int = 5
	// const y float64 = float64(num)
	const num int = 5
	const y float64 = float64(num)

	fmt.Println(num)
	fmt.Println(y)

	// An untyped constant does not have a fixed type until it is used.
	const z = 10 // Can be used as int or float

	var a int = z
	var b float64 = z

	fmt.Println(a, b)

	const (
		A = iota
		B
		C
	)

	fmt.Println(A, B, C)

	const (
		Read = 1 << iota 
		Write 
		Execute
	)

	fmt.Println(Read, Write, Execute)

	// Edge Cases & Rules

	// a)❌ Constants Cannot Change

	const testNum = 5
	// testNum = 6  // Uncomenting this will throw an error

	fmt.Println(testNum)

	// ❌ Cannot Use Variables to Define Constants

	var x = 32;
	// const number = x;
	fmt.Println(x)

	// ✔ Allowed:

	const x = 10
	const y = x * 2 // Works! (Compile-time computation)






}