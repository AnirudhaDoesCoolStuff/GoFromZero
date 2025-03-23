## **📌 Topic 8: Arrays in Go**

Arrays are **fixed-size** collections of elements of the **same type**.

### **1️⃣ Declaring an Array**

```go
var numbers [5]int  // Array of 5 integers
fmt.Println(numbers) // Output: [0 0 0 0 0] (default values)
```

- The **size of an array is fixed** at declaration and cannot be changed.
- The default value for an **int array** is `0`.

### **2️⃣ Initializing an Array**

```go
// Explicit initialization
var primes = [5]int{2, 3, 5, 7, 11}
fmt.Println(primes) // Output: [2 3 5 7 11]

// Shorter syntax
fruits := [3]string{"apple", "banana", "cherry"}
fmt.Println(fruits) // Output: [apple banana cherry]

// Let Go determine the size
numbers := [...]int{10, 20, 30}
fmt.Println(numbers) // Output: [10 20 30]
```

- The **`...` syntax** lets Go automatically set the array size based on the provided elements.

### **3️⃣ Accessing and Modifying Elements**

```go
numbers := [3]int{10, 20, 30}

fmt.Println(numbers[1]) // Output: 20

numbers[2] = 99
fmt.Println(numbers) // Output: [10 20 99]
```

- Arrays use **zero-based indexing** (`numbers[0]` is the first element).
- Trying to access an **out-of-bounds index** will cause a **runtime error**.

### **4️⃣ Looping Through an Array**

```go
numbers := [4]int{10, 20, 30, 40}

// Using `for`
for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}

// Using `range`
for index, value := range numbers {
    fmt.Println("Index:", index, "Value:", value)
}
```

- `len(arrayName)` gives the length of an array.
- `range` iterates over elements **without needing an explicit index variable**.

### **5️⃣ Copying an Array (By Value)**

```go
original := [3]int{1, 2, 3}
copyArr := original  // Copying an array

copyArr[0] = 99
fmt.Println(original) // Output: [1 2 3]
fmt.Println(copyArr)  // Output: [99 2 3]
```

- Arrays in Go are **copied by value**, meaning modifications to `copyArr` **won't affect** `original`.

### **6️⃣ Edge Cases**

✅ **Modifying an array inside a function won’t change the original unless passed as a pointer**

```go
func modifyArray(arr [3]int) {
    arr[0] = 100
}

numbers := [3]int{1, 2, 3}
modifyArray(numbers)
fmt.Println(numbers) // Output: [1 2 3] (Unchanged)
```

✅ **Using pointers to modify an array**

```go
func modifyArray(arr *[3]int) {
    arr[0] = 100
}

numbers := [3]int{1, 2, 3}
modifyArray(&numbers)
fmt.Println(numbers) // Output: [100 2 3]
```

---

## **📝 Challenges**

### **1️⃣ Basic Array Declaration**

Declare an array of **5 integers**, assign values, and print them.

### **2️⃣ Find the Maximum Value**

Given an array of integers, find and print the **maximum value**.

### **3️⃣ Reverse an Array**

Reverse the elements of a given array and print the reversed array.

### **4️⃣ Check if an Element Exists**

Write a program that checks if a given number exists in an array.

### **5️⃣ Sum of Elements**

Calculate and print the **sum of all elements** in an array.

---

Once you solve these challenges, send me your solutions, and I'll review them! 🚀
