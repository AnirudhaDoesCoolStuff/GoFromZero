package main

import "fmt"

func main() {
	// 1️⃣ Basic Loop: Print numbers from 1 to 10 using a for loop.
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 2️⃣ Even Numbers Only: Print even numbers from 1 to 20.
	for i := 2; i <= 20 ; i += 2 {
		fmt.Println(i)
	}

	// 3️⃣ Sum of First n Natural Numbers: Compute the sum of the first n natural numbers (n = 10).
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i;
	}
	fmt.Println(sum)

	// 4️⃣ Multiplication Table: Print the multiplication table for 5.
	for i := 1; i <= 10; i++ {
		fmt.Println(5 * i)
	}

	// 5️⃣ Reverse Loop: Print numbers from 10 to 1 using a for loop.
	for i := 10; i >= 1; i-- {
		fmt.Println(i)
	}

	// 6️⃣ Skip Multiples of 3: Print numbers from 1 to 20, but skip multiples of 3 using continue.
	for i := 1; i <= 20; i++ {
		if(i%3 == 0){
			continue
		}
		fmt.Println(i)
	}

	// 7️⃣ Loop Over a Slice: Given fruits := []string{"apple", "banana", "cherry"}, loop over it and print each fruit.
	fruits := []string{"apple", "banana", "cherry"}
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// 8️⃣ Find Prime Numbers: Print prime numbers between 1 and 50.
	
	for i := 2; i <= 50; i++ {
		if i == 2 {
			fmt.Println("2 is a prime number")
			continue
		}
		printPrimeNumber(i)
	}

	
}

func printPrimeNumber(number int) {
	var count int

	if number % 2 == 0 {
		return 
	}

	for i := 1; i <= number; i++ {
		if number % i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Printf("%d is a prime number \n", number)
	}
	
	
}