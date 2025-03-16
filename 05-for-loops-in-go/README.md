### **Topic 5: `for` Loops in Go**

Go only has **one looping construct**: the `for` loop.  
Unlike other languages (which have `while` and `do-while`), Go **only** uses `for` for all looping needs.

---

## **1️⃣ Basic `for` Loop**

A standard `for` loop consists of:

```go
for initialization; condition; post-action {
    // loop body
}
```

Example:

```go
for i := 1; i <= 5; i++ {
    fmt.Println("Iteration:", i)
}
```

✅ **Breakdown:**

- `i := 1` → Initializes `i`
- `i <= 5` → Runs the loop **while the condition is true**
- `i++` → Increments `i` after each iteration

---

## **2️⃣ Infinite `for` Loop (like `while` in other languages)**

If you omit all statements, `for` creates an **infinite loop**:

```go
for {
    fmt.Println("This will run forever!")
}
```

🔹 Equivalent to `while(true)` in other languages.  
❗ **Use `break` to prevent infinite execution!**

---

## **3️⃣ `for` as a `while` Loop**

A `for` loop **without initialization and post-action** acts like a `while` loop:

```go
i := 1
for i <= 5 {
    fmt.Println(i)
    i++ // Manual increment
}
```

✅ **Same as**:

```python
i = 1
while i <= 5:
    print(i)
    i += 1
```

---

## **4️⃣ Breaking and Continuing in Loops**

### **🔹 Using `break` (Exit Loop Early)**

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break // Exits the loop when i == 5
    }
    fmt.Println(i)
}
```

**Output:**

```
1
2
3
4
```

(Stops at `5` and exits)

---

### **🔹 Using `continue` (Skip an Iteration)**

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // Skips iteration when i == 3
    }
    fmt.Println(i)
}
```

**Output:**

```
1
2
4
5
```

(3 is skipped)

---

## **5️⃣ Nested `for` Loops (Multiplication Table Example)**

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("%d x %d = %d\n", i, j, i*j)
    }
}
```

✅ **Useful for:**  
✔ **Tables**  
✔ **2D matrix operations**  
✔ **Grid-based problems**

---

## **6️⃣ Looping Over Collections (Slices, Arrays, Maps)**

We use **`for` + `range`** to iterate over slices, arrays, and maps.

### **🔹 Iterating Over a Slice**

```go
numbers := []int{10, 20, 30}
for index, value := range numbers {
    fmt.Println("Index:", index, "Value:", value)
}
```

✔ If you **don't need the index**, use `_`:

```go
for _, value := range numbers {
    fmt.Println("Value:", value)
}
```

---

## **7️⃣ Edge Cases & Notes**

❌ **Incorrect Loop Condition: Infinite Loop**

```go
for i := 1; i <= 5; { // Missing i++
    fmt.Println(i)
}
```

✔ **Fix:**

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

✅ **Correct Way to Loop Backwards:**

```go
for i := 5; i >= 1; i-- {
    fmt.Println(i)
}
```

✅ **Skipping Every Other Number:**

```go
for i := 1; i <= 10; i += 2 {
    fmt.Println(i)
}
```

---

# **🚀 Challenges for You**

1️⃣ **Basic Loop:** Print numbers from **1 to 10** using a `for` loop.  
2️⃣ **Even Numbers Only:** Print even numbers from **1 to 20**.  
3️⃣ **Sum of First `n` Natural Numbers:** Compute the sum of the first `n` natural numbers (`n = 10`).  
4️⃣ **Multiplication Table:** Print the multiplication table for `5`.  
5️⃣ **Reverse Loop:** Print numbers from **10 to 1** using a `for` loop.  
6️⃣ **Skip Multiples of 3:** Print numbers from **1 to 20**, but **skip multiples of 3** using `continue`.  
7️⃣ **Loop Over a Slice:** Given `fruits := []string{"apple", "banana", "cherry"}`, loop over it and print each fruit.  
8️⃣ **Find Prime Numbers:** Print prime numbers between `1` and `50`.

---

Give these a shot and send me your solutions! I'll review and suggest improvements. 🚀
