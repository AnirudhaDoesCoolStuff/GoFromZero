package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println(1)
	// }

	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}
	
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // Exits the loop when i == 5
		}
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		if i == 3 {
			continue // Skips iteration when i == 3
		}
		fmt.Println(i)
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
	}

	numbers := []int{10, 20, 30}
	for index, value := range numbers {
		fmt.Println(index, value)
	}

	for _, value := range numbers {
		fmt.Println(value)
	}
	
}