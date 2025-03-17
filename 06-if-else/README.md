# **If/Else Statements in Go**

## **🔹 Overview**

- `if` statements allow conditional execution of code.
- **Parentheses (`()`) are not required** around conditions.
- **Curly braces `{}` are mandatory**, even for single-line blocks.
- The `else if` and `else` keywords can be used for multiple conditions.

---

## **🔹 Syntax**

### **✅ Basic If/Else**

```go
if condition {
    // Code to execute if condition is true
} else if anotherCondition {
    // Code to execute if another condition is true
} else {
    // Code to execute if all conditions are false
}
```

### **✅ Example**

```go
package main

import "fmt"

func main() {
	num := 10

	if num > 0 {
		fmt.Println("Positive number")
	} else if num < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Zero")
	}
}
```

✅ **No `()` around conditions.**  
✅ **Curly braces `{}` are required.**

---

## **🔹 Short Variable Declaration in If**

- Go allows **declaring a variable inside `if`**, which is only accessible within that block.

```go
if value := 5; value > 3 {
	fmt.Println("Value is greater than 3")
}
// fmt.Println(value)  // ❌ ERROR: value is not accessible outside `if`
```

✅ **The variable `value` exists only inside the `if` block.**

---

## **🔹 Logical Operators**

| Operator | Meaning | Example                     |
| -------- | ------- | --------------------------- | --- | --------------------- | --- | ---------------- |
| `&&`     | AND     | `if age >= 18 && age <= 60` |
| `        |         | `                           | OR  | `if day == "Saturday" |     | day == "Sunday"` |
| `!`      | NOT     | `if !isRainy`               |

### **Example**

```go
age := 20

if age >= 18 && age <= 60 {
	fmt.Println("You are an adult.")
} else {
	fmt.Println("You are either too young or a senior citizen.")
}
```

---

## **🔹 Edge Cases**

### **1️⃣ No Ternary Operator (`? :`)**

- Go does not support the ternary operator (`? :`).
- ✅ **Use an `if` statement instead**:
  ```go
  var min int
  if a < b {
      min = a
  } else {
      min = b
  }
  ```

---

### **2️⃣ Boolean Must Be Explicit**

- **Only boolean expressions are allowed in `if` conditions**:

  ```go
  someVar := true
  if someVar {  // ✅ Works
      fmt.Println("Valid condition")
  }

  num := 10
  if num {  // ❌ ERROR: Non-boolean condition
      fmt.Println("Invalid in Go")
  }
  ```

  ✅ **Go does not allow non-boolean values in `if` conditions.**

---

## **🚀 Challenges**

### **1️⃣ Even or Odd**

- Write a program that checks if a number is even or odd.
- **Input:** `10`
- **Output:** `"Even number"`

### **2️⃣ Grade Calculator**

- Take a score and print the grade:
  - 90+ → `"A"`
  - 80-89 → `"B"`
  - 70-79 → `"C"`
  - Below 70 → `"Fail"`
- **Input:** `85`
- **Output:** `"Grade: B"`

### **3️⃣ Leap Year Check**

- Write a program to check if a year is a leap year.
- **Leap Year Conditions**:
  - Divisible by 400 **✅ Leap Year**
  - Divisible by 4 but **not** by 100 **✅ Leap Year**
  - Otherwise **❌ Not a Leap Year**
- **Input:** `2024`
- **Output:** `"Leap Year"`

### **4️⃣ Login System**

- Take a **username** and **password** as input.
- Check if they match `"admin"` and `"1234"`, respectively.
- If correct, print `"Login Successful"`, otherwise `"Invalid Credentials"`.

---

Try these out and let me know your solutions! 🚀🔥
