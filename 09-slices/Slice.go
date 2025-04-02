package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(numbers)

	fruits := []string{"apple", "banana", "orange"}
	fmt.Println(fruits)

	nums := make([]int, 3)
	fmt.Println(nums)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:4]

	fmt.Println(slice)

	// Modifying a Slice
	// Since slices reference an underlying array, modifying a slice affects the original array.

	oddNums := [5]int{1, 3, 5, 7, 9}
	sliced := oddNums[0:3]

	sliced[0] = 77
	fmt.Println(sliced)
	fmt.Println(oddNums)

	evenNums := []int{2, 4, 6, 8}
	evenNums = append(evenNums, 10, 12, 14, 16, 18, 20)

	fmt.Println(evenNums)

	// Copy
	src := []int{1, 2, 3}
	dest := make([]int, len(src))

	copy(dest, src)
	fmt.Println(dest)
 }