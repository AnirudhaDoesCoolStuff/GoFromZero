Great! The next topic is **Switch Statements**.

I'll first explain everything in depth, covering all edge cases. Then, I'll give you challenges to solve. Let's get started! 🚀

---

## **🔹 Switch Statements in Go**

A `switch` statement is an alternative to `if-else` when you have multiple conditions to check. It makes your code **cleaner and more readable** than using many `if-else` blocks.

### **1️⃣ Basic Syntax**

```go
package main

import "fmt"

func main() {
	num := 2

	switch num {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number not recognized")
	}
}
```

✅ **How it works:**

- `num` is evaluated once.
- It **compares** the value against each `case`.
- If a match is found, that `case` block runs.
- If no match is found, the `default` case runs (optional but useful).

---

### **2️⃣ Multiple Case Values**

You can check multiple values in a single `case` statement:

```go
switch day := "Monday"; day {
case "Saturday", "Sunday":
	fmt.Println("It's the weekend!")
default:
	fmt.Println("It's a weekday.")
}
```

✅ **Why use this?**

- Reduces repetition.
- Makes conditions cleaner.

---

### **3️⃣ Switch Without Expression**

Go allows a **switch statement without an expression**, acting like an `if-else` ladder.

```go
score := 85

switch {
case score >= 90:
	fmt.Println("Grade: A")
case score >= 80:
	fmt.Println("Grade: B")
case score >= 70:
	fmt.Println("Grade: C")
default:
	fmt.Println("Grade: F")
}
```

✅ **Why use this?**

- The condition inside each `case` behaves like an `if-else` statement.
- It makes **range-based comparisons** easier.

---

### **4️⃣ Using `fallthrough`**

By default, Go **does NOT fall through** to the next case like in other languages (e.g., C).  
If you want to force fallthrough, use the `fallthrough` keyword:

```go
num := 1

switch num {
case 1:
	fmt.Println("Case 1")
	fallthrough
case 2:
	fmt.Println("Case 2 (Executed because of fallthrough)")
default:
	fmt.Println("Default case")
}
```

✅ **How it works:**

- **Without** `fallthrough`: Only **Case 1** would execute.
- **With** `fallthrough`: Case 2 **also executes** after Case 1.

---

### **5️⃣ Type Switch (for Interfaces)**

A `type switch` is used when you have an **interface** and want to check its actual type.

```go
var value interface{} = 42

switch v := value.(type) {
case int:
	fmt.Println("Integer:", v)
case string:
	fmt.Println("String:", v)
default:
	fmt.Println("Unknown Type")
}
```

✅ **Why use this?**

- Useful for working with **interfaces** and **dynamic types**.

---

## **🛑 Edge Cases & Important Notes**

1️⃣ **No Fallthrough by Default:**

- Unlike some languages, Go **does NOT fall through** unless you explicitly use `fallthrough`.

2️⃣ **Default Case Placement:**

- `default` can be **anywhere** in the switch, but it usually comes last.

3️⃣ **Switch Without a Variable:**

- Using a **switch without an expression** lets you handle **complex conditions** without writing `if-else`.

---

## **📝 Challenges for You**

Now, it's your turn! Solve these challenges using `switch`, and I'll review your solutions.

### **1️⃣ Basic Switch Statement**

Write a program that checks a number and prints:

- `"One"` if it's `1`
- `"Two"` if it's `2`
- `"Three"` if it's `3`
- `"Unknown Number"` otherwise

### **2️⃣ Multiple Case Values**

Write a switch statement that checks if a given day is a weekday or a weekend.

### **3️⃣ Grade Calculator (Switch Without Expression)**

Given a score, print the grade:

- `90+` → `"A"`
- `80-89` → `"B"`
- `70-79` → `"C"`
- Otherwise → `"F"`

### **4️⃣ Fallthrough Challenge**

Write a switch statement where:

- Case `1` prints `"One"` and falls through.
- Case `2` prints `"Two"` and stops.
- Default prints `"Not One or Two"`.

### **5️⃣ Type Switch**

Create a `switch` that takes an **interface{}** value and checks if it's an `int`, `string`, or `bool`.

---

🚀 Solve these challenges, and I'll review your solutions! Let me know if anything needs more explanation. Happy coding! 😊
