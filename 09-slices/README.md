Great! The next topic is **Slices**. Here's your structured learning plan:

---

## **📌 Topic 9: Slices in Go**

Slices are **dynamic, flexible**, and more commonly used than arrays in Go.

### **1️⃣ What is a Slice?**

- A **slice** is a **resizable** version of an array.
- Unlike arrays, slices **don’t have a fixed size**.
- They are just **references** to an **underlying array**.

```go
numbers := []int{10, 20, 30, 40, 50}
fmt.Println(numbers) // Output: [10 20 30 40 50]
```

- **No need to specify size** like arrays.

### **2️⃣ Creating a Slice**

#### **Using a Literal**

```go
fruits := []string{"apple", "banana", "cherry"}
fmt.Println(fruits) // Output: [apple banana cherry]
```

#### **Using `make()` (for dynamic slices)**

```go
numbers := make([]int, 3) // Creates a slice with length 3
fmt.Println(numbers)      // Output: [0 0 0]
```

- `make([]Type, length, capacity)` lets you set **capacity** explicitly.
- If capacity is omitted, it is equal to length.

#### **From an Array**

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // Creates a slice from index 1 to 3 (excluding 4)
fmt.Println(slice) // Output: [2 3 4]
```

- **Slicing syntax:** `arr[start:end]` (includes `start`, excludes `end`).

### **3️⃣ Modifying a Slice**

Since slices reference an **underlying array**, modifying a slice **affects the original array**.

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4] // [2 3 4]

slice[0] = 99
fmt.Println(arr)   // Output: [1 99 3 4 5] (Original array modified)
```

### **4️⃣ Appending Elements**

Use `append()` to add elements dynamically.

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5) // Adds 4 and 5
fmt.Println(nums)         // Output: [1 2 3 4 5]
```

- `append()` creates **a new slice if capacity is exceeded**.

### **5️⃣ Copying Slices**

Use `copy(dest, src)` to copy elements from one slice to another.

```go
src := []int{1, 2, 3}
dest := make([]int, len(src))

copy(dest, src)
fmt.Println(dest) // Output: [1 2 3]
```

- Copying is useful when you **don’t want to modify the original slice**.

### **6️⃣ Iterating Over a Slice**

Using **`for`** loop:

```go
nums := []int{10, 20, 30}
for i := 0; i < len(nums); i++ {
    fmt.Println(nums[i])
}
```

Using **`range`** loop:

```go
for index, value := range nums {
    fmt.Println("Index:", index, "Value:", value)
}
```

### **7️⃣ Edge Cases**

✅ **Slicing beyond capacity leads to errors**

```go
nums := []int{1, 2, 3}
sub := nums[:4] // ERROR: index out of range
```

✅ **Appending a large number of elements increases capacity dynamically**

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5, 6, 7, 8, 9) // Creates a new slice with increased capacity
fmt.Println(nums)
```

---

## **📝 Challenges**

### **1️⃣ Create and Print a Slice**

Create a slice of **5 strings**, assign values, and print them.

### **2️⃣ Modify a Slice and Show the Effect on the Original Array**

Create a slice from an **array**, modify it, and print both to see the effect.

### **3️⃣ Find the Maximum Value in a Slice**

Write a function to find and print the **maximum value** in a slice.

### **4️⃣ Reverse a Slice In-Place**

Reverse the order of elements in a slice **without creating a new slice**.

### **5️⃣ Remove an Element from a Slice**

Write a function that removes an element at **index `i`** from a slice without leaving an empty space.

---

Try solving these challenges and send me your solutions! 🚀
