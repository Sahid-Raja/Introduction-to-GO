# Introduction to Go

Welcome to the **Introduction to Go** project! This repository is designed to help you learn the Go programming language through practical examples divided into **Basic**, **Fundamentals**, and **Advanced** sections.

## Table of Contents
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [How to Run](#how-to-run)
- [Project Structure Explanation](#project-structure-explanation)

---

## Project Structure

This project is divided into three main sections: **Basic**, **Fundamentals**, and **Advanced**. Each section contains Go examples that demonstrate core concepts and features of the language. The entry point (`main.go`) is set up so that you can easily switch between different parts of the code by commenting and uncommenting relevant sections.

```
Introduction-to-GO/
├── cmd/
│   └── myApp/
│       └── main.go                # Entry point of the application
├── pkg/
│   ├── basic/                     # Basic Go concepts and examples
│   ├── fundamentals/              # Fundamental Go concepts
│   └── advanced/                  # Advanced Go topics
├── go.mod                         # Go module definition
└── README.md                      # Project documentation
```

---

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/Sahid-Raja/Introduction-to-GO.git
cd Introduction-to-GO
```

2. Run `go mod tidy` to download dependencies:

```bash
go mod tidy
```

3. Install Go if you don't have it already: [Go Installation Guide](https://golang.org/doc/install).

---

## How to Run

In the `cmd/myApp/main.go` file, you'll find different Go examples from **Basic**, **Fundamentals**, and **Advanced** packages. You can switch between these examples by commenting and uncommenting the relevant code sections.

For example:

1. **To run the Basic example**, comment out all other code sections and uncomment the line that imports and calls the function from `basic`:

```go
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/basic" // Uncomment for basic example
// basic.function()                                         // Uncomment for basic example
```

2. **To test Fundamentals**, uncomment the lines for the `fundamentals` example:

```go
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/fundamentals" // Uncomment for fundamentals example
// fundamentals.function()                                           // Uncomment for fundamentals example
```

3. **To try Advanced features**, uncomment the code in the `advanced` example:

```go
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/advanced" // Uncomment for advanced example
// advanced.function()                                            // Uncomment for advanced example
```

Once you've made the changes, you can run the project with:

```bash
go run cmd/myApp/main.go
```

---

## Project Structure Explanation

### `cmd/myApp/main.go`
This file is the main entry point of the application. It imports different packages from the `pkg/` directory and allows you to run different examples. Each section of code (Basic, Fundamentals, Advanced) can be tested by commenting or uncommenting the corresponding import and function calls.

### `pkg/basic/`
This directory contains basic Go concepts and examples, helping you understand the foundation of Go programming.

### `pkg/fundamentals/`
This directory contains more detailed Go concepts that build upon the basic knowledge and introduce more advanced functionality.

### `pkg/advanced/`
This directory focuses on advanced Go topics, such as concurrency, goroutines, channels, and more.

---

### **Example Usage**
In your `main.go` file, you can comment/uncomment blocks as follows:

```go
package main

// Uncomment for Basic Example
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/basic"
// basic.HelloWorld()

// Uncomment for Fundamentals Example
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/fundamentals"
// fundamentals.HelloWorld()

// Uncomment for Advanced Example (Concurrency)
// import "github.com/Sahid-Raja/Introduction-to-GO/pkg/advanced"
// advanced.RunConcurrencyExample()
```

---

This setup allows users to explore different Go concepts by simply commenting or uncommenting sections in the `main.go` file. Let me know if you need any more adjustments or additional details!