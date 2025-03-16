package main

import "fmt"

func main() {
	// 1️⃣ Basic Constant Declaration
	const name = "Anirudha"
	const country = "India"
	const age = 24
	fmt.Println(name, country, age)

	// 2️⃣ Typed vs. Untyped Constants
	const testNum = 10
	var testInt int = testNum
	var testFloat float64 = testNum
	fmt.Println(testInt, testFloat)

	// 3️⃣ Constant Expression (Rectangle Area)
	const length = 5
	const width = 10
	const area = length * width
	fmt.Println("Rectangle Area:", area)

	// 4️⃣ iota Challenge (Gold, Silver, Bronze) - Fixed
	const (
		Gold = iota + 1 // Now starts at 1
		Silver
		Bronze
	)
	fmt.Println(Gold, Silver, Bronze) // Correct Output: 1 2 3

	// 5️⃣ Bitwise iota Challenge (File Permissions)
	const (
		Read = 1 << iota  // 1 (0001)
		Write             // 2 (0010)
		Execute           // 4 (0100)
	)
	fmt.Println(Read, Write, Execute) // Output: 1 2 4
}
