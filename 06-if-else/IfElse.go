package main

import "fmt"

func main() {
	num := 10

	if num > 0 {
		fmt.Println("Positive Number")
	} else if num < 0 {
		fmt.Println("Negative Number")
	} else {
		fmt.Println("Zero")
	}

	// Short Variable Declaration in If

	if value := 5; value > 2 {
		fmt.Println("Number is greater than 2")
	}
	

	someVar := true
	if someVar {  // ✅ Works
		fmt.Println("Valid condition")
	}

	// num := 10
	// if num {  // ❌ ERROR: Non-boolean condition
	// 	fmt.Println("Invalid in Go")
	// }
}