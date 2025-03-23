package main

import "fmt"

func main() {
	// Declaring an array
	var numbers [5]int;
	fmt.Println(numbers)

	// Initializing an array
	var primes = [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)

	// shorter version
	fruits := [3]string{"apple", "banana", "mango"}
	fmt.Println(fruits)

	// Let go determine the number
	evens := [...]int{2, 4, 6, 8, 10}
	fmt.Println(evens)

	// Accessing and Modifying Elements
	nums := [3]int{1, 2, 3}
	nums[2] = 1

	fmt.Println(nums)

	// Looping Through an Array

	// Using 'for'
	oddNums := [4]int{1, 3, 5, 7}

	for i := 1; i < len(oddNums); i++ {
		fmt.Println(oddNums[i])
	}

	for index, value := range oddNums {
		fmt.Println("Index: ", index, " Odd number: ", value)
	}

	// Copying an Array (By Value)
	evenNums := [3]int{2, 4, 6}
	evenNumsCopied := evenNums

	fmt.Println(evenNums)
	fmt.Println(evenNumsCopied)



	
}


func mofigyArray(arr int[])