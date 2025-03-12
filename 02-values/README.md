### **Topic 2: Values in Go**

#### **1. Introduction**

Values in Go refer to the different types of data you can store and manipulate, such as numbers, strings, and booleans. Understanding how Go handles values is key to working efficiently with variables and expressions.

---

#### **2. Basic Types of Values in Go**

##### **a) Numeric Values**

Go supports integers and floating-point numbers:

```go
fmt.Println(42)       // Integer
fmt.Println(3.14)     // Floating-point number
```

##### **b) String Values**

Strings in Go are sequences of characters enclosed in double quotes:

```go
fmt.Println("Hello, Go!")
fmt.Println("Go is awesome")
```

##### **c) Boolean Values**

Boolean values are `true` or `false`:

```go
fmt.Println(true)
fmt.Println(false)
```

---

#### **3. Edge Cases & Special Considerations**

##### **a) Integer Overflow**

Go’s integer types have fixed sizes. If an integer exceeds its limit, it causes an overflow.

Example (assuming `uint8` range is 0-255):

```go
var x uint8 = 255
x++  // ERROR: overflows uint8
```

##### **b) Floating-Point Precision Issues**

Floating-point numbers are not always exact due to precision limitations:

```go
fmt.Println(0.1 + 0.2) // Output: 0.30000000000000004 (Not exactly 0.3)
```

##### **c) Empty String and Zero Values**

Go assigns default values when variables are declared but not initialized:

```go
var s string
var n int
var b bool
fmt.Println(s) // Empty string: ""
fmt.Println(n) // Zero: 0
fmt.Println(b) // False: false
```

##### **d) Escape Characters in Strings**

```go
fmt.Println("Hello\nWorld")  // New line
fmt.Println("Tab\tSpace")     // Tab space
fmt.Println("\"Quoted Text\"") // Double quotes
```

##### **e) Unicode and Emojis in Strings**

```go
fmt.Println("こんにちは")  // Japanese
fmt.Println("🚀 Go!")      // Emojis
```

---

### **4. Challenges 🚀**

1. **Basic Operations**: Print different types of values (integer, float, string, boolean) in a single `Println` statement.
2. **Formatted Output**: Use `Printf` to display a sentence like: `"The price of an item is $49.99"` with a floating-point value.
3. **Escape Sequences**: Print a string containing new lines, tabs, and quotes.
4. **Unicode Challenge**: Print a sentence in your native language along with an emoji.

---

Try solving these challenges 🚀
