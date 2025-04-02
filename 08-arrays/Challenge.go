package main

import "fmt"

func main() {
	// 	1️⃣ Basic Array Declaration
	// Declare an array of 5 integers, assign values, and print them.
	var numbers [5]int;
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	fmt.Println(numbers)


	// 2️⃣ Find the Maximum Value
	// Given an array of integers, find and print the maximum value.
	

	allNumbers := [5]int{112 , 175, 123, 523, 32}

	big := allNumbers[0]

	for i := 1; i < len(allNumbers); i++ {
		if allNumbers[i] > big {
			big = allNumbers[i]
		}
	}

	fmt.Println(big)

	// 3️⃣ Reverse an Array
	// Reverse the elements of a given array and print the reversed array.

	var revArr [5]int
	var count int;

	for j := len(allNumbers)-1; j >= 0; j-- {
		revArr[count] = allNumbers[j]
		count++
	}

	fmt.Println(revArr)

	// 4️⃣ Check if an Element Exists
	// Write a program that checks if a given number exists in an array.
	var searchEle int = 175;
	eleFound := false
	
	for _ , value := range allNumbers {
		if value == searchEle {
			eleFound = true
			break
		}
	}

	if eleFound == true {
		fmt.Println("Element found")
	} else {
		fmt.Println("Element Not found!")
	}

	// 5️⃣ Sum of Elements
	// Calculate and print the sum of all elements in an array.

	var sum int;

	for _, value := range allNumbers {
		sum += value;
	}

	fmt.Println(sum)
}