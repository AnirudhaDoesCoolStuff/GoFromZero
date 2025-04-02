package main

import "fmt"

func main() {
	// 1️⃣ Create and Print a Slice
	// Create a slice of 5 strings, assign values, and print them.
	names := []string{"john", "jane", "jin", "jim", "kim"}
	fmt.Println(names)

	// 2️⃣ Modify a Slice and Show the Effect on the Original Array
	// Create a slice from an array, modify it, and print both to see the effect.
	testArr := [5]int{1, 2, 3, 4, 5}
	testSlice := testArr[0:3]

	testSlice[0] = 44
	fmt.Println(testArr)
	fmt.Println(testSlice)

	// 3️⃣ Find the Maximum Value in a Slice
	// Write the code to find and print the maximum value in a slice.

	randomNums := []int{1, 334, 534, 123, 234, 32423, 234, 23 }

	big := randomNums[0]

	for i := 1; i < len(randomNums); i++ {
		if randomNums[i] > big {
			big = randomNums[i]
		}
	} 

	fmt.Println(big)



	// 4️⃣ Reverse a Slice In-Place
	// Reverse the order of elements in a slice without creating a new slice.

	s := []int{1, 2, 3, 4, 5}

	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}

	fmt.Println(s)

	// 5️⃣ Remove an Element from a Slice
	// Write a function that removes an element at index i from a slice without leaving an empty space.

	s2 := []int{1, 2, 3, 4, 5}
	i := 2
	s2 = append(s2[:i], s2[i+1:]...)

	fmt.Println(s2)
	
}