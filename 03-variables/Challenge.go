package main

import "fmt"

func main() {
	// 1. Basic Variable Declaration: Declare and print a name, age, and height using both var and :=
	var nameOne string = "John Doe"
	var ageOne int = 24
	var heightOne float64 = 153.5

	nameTwo := "Jane Doe"
	ageTwo := 25
	heightTwo := 155.55

	fmt.Printf("Person 1: %s, Age: %d, Height: %.2f\n", nameOne, ageOne, heightOne)
	fmt.Printf("Person 2: %s, Age: %d, Height: %.2f\n", nameTwo, ageTwo, heightTwo)

	// 2. Multiple Variables: Declare multiple variables in a single line and print them.
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)

	// 3. Constants: Define a constant pi and use it in a calculation (e.g., area of a circle).
	const pi = 3.14
	area := pi * 3.33 * 4.44
	fmt.Printf("Area: %.2f\n", area)

	// 4. Scope Issue: Try to redeclare a variable inside a function and observe the error.
	test()

	// 5. Unused Variable Fix: Declare a variable but prevent Go from throwing an unused variable error.
	var thisIsUnUsedVar int
	_ = thisIsUnUsedVar
}

func test() {
	testTheVar := 5
	// testTheVar := "John"  // Uncommenting will cause an error
	fmt.Println(testTheVar)
}
