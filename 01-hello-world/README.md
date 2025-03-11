### **Topic 1: Hello World in Go**

#### **1. Introduction**

"Hello World" is the simplest Go program and serves as a starting point for understanding Go syntax, the structure of a Go program, and how Go handles execution.

---

#### **2. Basic "Hello World" Program**

Here’s the simplest form of a Go program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

#### **3. Explanation**

- **`package main`**: Every Go program starts with a package declaration. The `main` package is special—it tells the Go compiler that this is an executable program.
- **`import "fmt"`**: The `fmt` package provides functions for formatted I/O (input/output).
- **`func main()`**: This is the entry point of a Go program. Execution starts from `main()`.
- **`fmt.Println("Hello, World!")`**: `Println` prints a string followed by a newline.

---

#### **4. Edge Cases & Variations**

##### **a) Using `Print` Instead of `Println`**

```go
fmt.Print("Hello, ")
fmt.Print("World!\n")
```

- `Print` does not add a newline automatically, so we need `\n` manually.

##### **b) Using `Printf` for Formatting**

```go
fmt.Printf("Hello, %s!\n", "Go")
```

- `Printf` allows formatted strings, useful when working with variables.

##### **c) Missing `import "fmt"`**

```go
package main

func main() {
    Println("Hello, World!") // ERROR: undefined Println
}
```

- `fmt` must be imported before using its functions.

##### **d) Declaring `main()` in a Different Package**

```go
package mypackage  // ERROR: main function must be in 'main' package

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

- The `main()` function **must** be inside the `main` package for executable programs.

##### **e) Using Multiple Print Statements**

```go
fmt.Println("Hello")
fmt.Println("World")
```

- Each `Println` prints on a new line.

##### **f) Unicode Characters in Print**

```go
fmt.Println("こんにちは, 世界!") // Hello, World! in Japanese
fmt.Println("😊 Hello, World!") // Using emoji
```

- Go supports Unicode natively, so you can print non-English characters and emojis.

---

### **5. Challenges** 🚀

Try these exercises:

1. **Basic Printing**: Write a program that prints your name and a greeting on separate lines.
2. **Formatted Output**: Use `Printf` to print your name and age like: `"My name is John and I am 25 years old."`
3. **Multi-line String**: Print a multi-line string using a single `Println`.
4. **Experiment with Unicode**: Print a message containing emojis and non-English characters.

Try solving these challenges 🚀
