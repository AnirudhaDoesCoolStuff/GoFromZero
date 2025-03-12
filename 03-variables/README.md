### **Topic 3: Variables in Go**

#### **1. Introduction**

Variables in Go are used to store values that can change during program execution. Go is a statically typed language, meaning each variable has a defined type.

---

#### **2. Declaring Variables in Go**

There are **three ways** to declare variables in Go:

##### **a) Using `var` (Explicit Type Declaration)**

```go
var name string = "Anirudha"
var age int = 24
var height float64 = 5.9
```

- You specify the type explicitly.
- Useful when you need a default value.

##### **b) Using `var` Without Initialization (Default Values)**

```go
var city string  // Defaults to ""
var score int    // Defaults to 0
var isActive bool // Defaults to false
```

- When no value is provided, Go assigns a **zero value** based on the type:
  - `0` for integers
  - `""` (empty string) for strings
  - `false` for booleans

##### **c) Using `:=` (Short Variable Declaration - Implicit Type)**

```go
name := "Anirudha"
age := 24
height := 5.9
```

- Go **infers** the type automatically.
- Can only be used inside functions.

---

#### **3. Multiple Variable Declaration**

##### **a) Declaring Multiple Variables in a Single Line**

```go
var a, b, c int = 1, 2, 3
x, y, z := "Go", true, 3.14
```

##### **b) Declaring Multiple Variables Using a Block**

```go
var (
    firstName string = "Anirudha"
    lastName  string = "XYZ"
    age       int    = 24
)
```

- Helps keep code organized when declaring multiple variables.

---

#### **4. Constants (`const`)**

Constants are variables whose values **cannot change** after being assigned.

```go
const pi = 3.14159
const appName = "GoMaster"
```

- **Cannot use `:=` for constants**.

---

#### **5. Edge Cases & Special Considerations**

##### **a) Redeclaring Variables (Scope Issues)**

```go
name := "John"
name := "Doe" // ERROR: cannot redeclare name in the same scope
```

- Variables **cannot be redeclared** in the same scope with `:=`.

##### **b) Unused Variables**

```go
var unusedVar int = 10 // ERROR: declared but not used
```

- Go **does not allow unused variables**.
- Use `_` (blank identifier) to avoid this error:
  ```go
  _ = unusedVar
  ```

##### **c) Constants Cannot Be Changed**

```go
const pi = 3.14
pi = 3.14159 // ERROR: cannot assign to pi
```

##### **d) Type Mismatch**

```go
var num int = 10
num = "hello" // ERROR: cannot assign string to int
```

- Go is **statically typed**, meaning you **cannot** change a variable’s type after declaration.

---

### **6. Challenges 🚀**

1. **Basic Variable Declaration:** Declare and print a `name`, `age`, and `height` using both `var` and `:=`.
2. **Multiple Variables:** Declare multiple variables in a single line and print them.
3. **Constants:** Define a constant `pi` and use it in a calculation (e.g., area of a circle).
4. **Scope Issue:** Try to redeclare a variable inside a function and observe the error.
5. **Unused Variable Fix:** Declare a variable but prevent Go from throwing an unused variable error.

Try these and share your solutions! I’ll review and suggest improvements. 🚀
