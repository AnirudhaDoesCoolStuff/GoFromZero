### **Topic 4: Constants in Go**

#### **1. What are Constants?**

Constants (`const`) are **fixed values** that **cannot be changed** after being assigned. Unlike variables, they must be assigned at the time of declaration.

```go
const pi = 3.14159
const appName = "GoMaster"
```

- You **cannot** use `:=` to declare constants.
- Constants improve code **readability** and **prevent accidental modification**.

---

### **2. Declaring Constants in Go**

#### **a) Single Constant Declaration**

```go
const name = "Anirudha"
const age = 24
const pi = 3.14
```

- The value must be **known at compile-time**.

#### **b) Multiple Constants (Block Declaration)**

```go
const (
    country   = "India"
    city      = "Bangalore"
    isGoGreat = true
)
```

- This keeps constants organized.

---

### **3. Typed vs Untyped Constants**

#### **a) Typed Constants**

A **typed constant** has a fixed type.

```go
const x int = 10
const y float64 = 3.14
```

- You **cannot** assign `y` (float64) to an `int` variable without explicit conversion.

#### **b) Untyped Constants**

An **untyped constant** does **not** have a fixed type until it is used.

```go
const z = 10 // Untyped, can be used as int or float
```

Example:

```go
var a int = z      // Works (z behaves as int)
var b float64 = z  // Works (z behaves as float)
```

This makes **untyped constants more flexible**.

---

### **4. Constant Expressions & Computations**

Constants can be used in expressions.

```go
const x = 10
const y = x * 2  // Computed at compile-time
const z = x + y  // Works fine
```

- Since `x` and `y` are known at **compile-time**, Go allows computations.

---

### **5. `iota` - Constant Enumerator**

`iota` is a special constant used for **enumeration** in Go.

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

- `iota` starts at `0` and increments automatically.
- Useful for **assigning unique values** to constants.

Example: **Days of the Week**

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
    Wednesday      // 3
    Thursday       // 4
    Friday         // 5
    Saturday       // 6
)
```

Example: **Bit Flags**

```go
const (
    Read   = 1 << iota // 1
    Write              // 2
    Execute            // 4
)
```

- This is used in permission systems (Read, Write, Execute).

---

### **6. Edge Cases & Rules**

❌ **Constants Cannot Change**

```go
const x = 10
x = 20 // ERROR: cannot assign to x
```

❌ **Cannot Use Variables to Define Constants**

```go
var num = 5
const newNum = num // ERROR: num is not a constant
```

- Constants must be **known at compile-time**.

✔ **Allowed:**

```go
const x = 10
const y = x * 2 // Works! (Compile-time computation)
```

---

### **7. Challenges 🚀**

1️⃣ **Basic Constant Declaration:** Define and print constants for your `name`, `country`, and `age`.  
2️⃣ **Typed vs. Untyped Constants:** Declare an **untyped constant** and assign it to different types of variables (`int`, `float64`).  
3️⃣ **Constant Expression:** Define constants `length = 5`, `width = 10`, and compute the **area of a rectangle** using a constant expression.  
4️⃣ **iota Challenge:** Use `iota` to create an enumeration for `Gold`, `Silver`, and `Bronze`, with values **1, 2, and 3** respectively.  
5️⃣ **Bitwise iota Challenge:** Use `iota` to define constants for file permissions (`Read = 1`, `Write = 2`, `Execute = 4`).

Try these challenges 🚀
