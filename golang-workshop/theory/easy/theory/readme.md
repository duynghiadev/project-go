### Workshop for Computer Science Students on Go Language

## Part 1: Lecture (1 hour)

### Introduction to Go
- **History and Features of the Language**
    - Created at Google in 2009
    - Compiles to native code
    - Strong typing, automatic memory management
    - High performance, simplicity, and safety

### Go Syntax Basics

#### 1. Installation and Running a Program
- **Installing Go**:
    - Website: [https://golang.org/dl/](https://golang.org/dl/)
    - Installing packages: `go get <package>`

- **Program Structure**:
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, World!")
  }
  ```

- **Running a Program**:
    - `go run <filename.go>`
    - Compiling: `go build <filename.go>`

#### 2. Data Types and Variables
- **Basic Data Types**:
    - Integers: `int`, `int8`, `int16`, `int32`, `int64`
    - Floating-point numbers: `float32`, `float64`
    - Strings: `string`
    - Booleans: `bool`

- **Variable Declaration**:
  ```go
  var a int = 10
  var b = 20 // implicit type
  c := 30 // short declaration
  ```

#### 3. Control Flow
- **Conditional Statements**:
  ```go
  if condition {
      // code
  } else {
      // code
  }
  ```

- **Loops**:
  ```go
  for i := 0; i < 10; i++ {
      fmt.Println(i)
  }
  ```

#### 4. Collections and Working with Them
- **Arrays**:
  ```go
  var arr [5]int
  arr[0] = 1
  ```

- **Slices**:
  ```go
  s := []int{1, 2, 3}
  s = append(s, 4)
  ```

- **Maps**:
  ```go
  m := make(map[string]int)
  m["one"] = 1
  ```
