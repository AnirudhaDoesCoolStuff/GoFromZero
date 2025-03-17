package main

import "fmt"

func main() {

	// 1️⃣ Even or Odd
	num := 5

	if num % 2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// 2️⃣ Grade Calculator
	grade := 99

	if grade >= 90 {
		fmt.Println("Grade is A")
	} else if grade >= 80 && grade < 90 {
		fmt.Println("Grade is B")
	} else if grade >= 70 && grade < 80 {
		fmt.Println("Grade is C")
	} else {
		fmt.Println("Fail")
	}

	// 3️⃣ Leap Year Check
	year := 2000
	
	if year % 4 == 0 && year % 100 != 0  || year % 400 == 0 {
		fmt.Println(year, " is a leap year.");
	} else {
		fmt.Println(year, " is not a leap year.");
	}

	// 4️⃣ Login System

	username := "admin"
	password := 12345

	if username == "admin" && password == 12345 {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Invalid Credentials")
	}

}