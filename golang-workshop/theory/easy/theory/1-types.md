### Workshop for Computer Science Students on Go Language: Data Types and Variables

## Part 1: Lecture

### Data Types and Variables in Go

#### Introduction to Data Types and Variables
- **Data Types**: Define the kind of data a variable can hold.
- **Variables**: Named storage locations that hold data.

### Basic Data Types in Go

#### 1. Integer Types
Integers are whole numbers without a fractional component. Go provides several integer types of varying sizes and signedness:

- **Signed Integers**: Can hold both positive and negative numbers.
    - `int`: The default integer type, with a size dependent on the platform (typically 32 or 64 bits).
    - `int8`: 8-bit integer, ranging from -128 to 127.
    - `int16`: 16-bit integer, ranging from -32768 to 32767.
    - `int32`: 32-bit integer, ranging from -2147483648 to 2147483647.
    - `int64`: 64-bit integer, ranging from -9223372036854775808 to 9223372036854775807.

- **Unsigned Integers**: Can hold only non-negative numbers.
    - `uint`: The default unsigned integer type, with a size dependent on the platform.
    - `uint8`: 8-bit unsigned integer, ranging from 0 to 255 (also known as `byte`).
    - `uint16`: 16-bit unsigned integer, ranging from 0 to 65535.
    - `uint32`: 32-bit unsigned integer, ranging from 0 to 4294967295.
    - `uint64`: 64-bit unsigned integer, ranging from 0 to 18446744073709551615.

**Example**:
```go
package main

import "fmt"

func main() {
    var a int = 10
    var b int8 = -128
    var c uint16 = 65535
    var d int64 = 9223372036854775807

    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
    fmt.Println("d:", d)
}
```

#### 2. Floating-point Numbers
Floating-point numbers are numbers with a fractional component. Go provides two floating-point types:

- `float32`: 32-bit floating-point number.
- `float64`: 64-bit floating-point number (default type).

**Example**:
```go
package main

import "fmt"

func main() {
    var a float32 = 3.14
    var b float64 = 2.718281828459045

    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

#### 3. Strings
Strings are sequences of characters. They are immutable in Go, meaning once a string is created, it cannot be changed.

**Example**:
```go
package main

import "fmt"

func main() {
    var s string = "Hello, World!"
    fmt.Println(s)

    // Strings can also be concatenated
    var firstName string = "John"
    var lastName string = "Doe"
    fullName := firstName + " " + lastName
    fmt.Println(fullName)
}
```

#### 4. Booleans
Booleans represent true or false values.

**Example**:
```go
package main

import "fmt"

func main() {
    var t bool = true
    var f bool = false

    fmt.Println("t:", t)
    fmt.Println("f:", f)
}
```

### Variable Declaration and Initialization

#### 1. Using the `var` Keyword
Variables can be declared using the `var` keyword with an explicit type or with type inference.

**Example**:
```go
package main

import "fmt"

func main() {
    var a int = 10 // explicit type
    var b = 20     // type inference
    fmt.Println("a:", a)
    fmt.Println("b:", b)
}
```

#### 2. Short Variable Declaration
Short variable declaration (`:=`) is a concise way to declare and initialize variables. This form is only available inside functions.

**Example**:
```go
package main

import "fmt"

func main() {
    c := 30 // short variable declaration
    d := "Hello"
    e := 3.14
    fmt.Println("c:", c)
    fmt.Println("d:", d)
    fmt.Println("e:", e)
}
```

### Constants
Constants are immutable values that are known at compile time and do not change during the execution of the program. They are declared using the `const` keyword.

**Example**:
```go
package main

import "fmt"

const Pi = 3.14159
const Greeting = "Hello, World!"

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println(Greeting)
}
```

### Type Conversion
Go requires explicit type conversion when assigning a value of one type to a variable of another type.

**Example**:
```go
package main

import "fmt"

func main() {
    var a int = 10
    var b float64 = float64(a)
    var c uint = uint(b)

    fmt.Println("a:", a)
    fmt.Println("b:", b)
    fmt.Println("c:", c)
}
```

### Zero Values
Variables declared without an initial value are automatically assigned a "zero value" based on their type.

- Integers: `0`
- Floating-point numbers: `0.0`
- Strings: `""` (empty string)
- Booleans: `false`

**Example**:
```go
package main

import "fmt"

func main() {
    var a int
    var b float64
    var c string
    var d bool

    fmt.Println("a:", a) // 0
    fmt.Println("b:", b) // 0.0
    fmt.Println("c:", c) // ""
    fmt.Println("d:", d) // false
}
```

### Summary
- Go provides several basic data types, including integers, floating-point numbers, strings, and booleans.
- Variables can be declared using the `var` keyword or short variable declaration (`:=`).
- Constants are immutable values declared using the `const` keyword.
- Type conversion in Go requires explicit casting.
- Uninitialized variables are assigned zero values.

---

## Part 2: Practical Session

### Practical Task: Implementing a Simple Calculator

**Task Description**: Write a program that acts as a simple calculator. The program should read two numbers and an operator from the user and perform the corresponding arithmetic operation.

**Requirements**:
1. The program should prompt the user to enter two numbers and an operator (`+`, `-`, `*`, `/`).
2. Use appropriate data types for the numbers.
3. Implement functions to handle each arithmetic operation.
4. Print the result of the operation.

**Example Implementation**:

```go
package main

import (
    "fmt"
)

func add(a, b float64) float64 {
    return a + b
}

func subtract(a, b float64) float64 {
    return a - b
}

func multiply(a, b float64) float64 {
    return a * b
}

func divide(a, b float64) float64 {
    if b == 0 {
        fmt.Println("Error: Division by zero")
        return 0
    }
    return a / b
}

func main() {
    var num1, num2 float64
    var operator string

    fmt.Print("Enter first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter second number: ")
    fmt.Scanln(&num2)
    fmt.Print("Enter operator (+, -, *, /): ")
    fmt.Scanln(&operator)

    var result float64
    switch operator {
    case "+":
        result = add(num1, num2)
    case "-":
        result = subtract(num1, num2)
    case "*":
        result = multiply(num1, num2)
    case "/":
        result = divide(num1, num2)
    default:
        fmt.Println("Invalid operator")
        return
    }

    fmt.Printf("Result: %f\n", result)
}
```

### Task Recommendations:
1. **Modularize the Code**: Create separate functions for each arithmetic operation.
2. **Error Handling**: Add error handling for invalid input and division by zero.
3. **Extend Functionality**: Add support for more operations like modulus (`%`) and exponentiation.

### Conclusion
- Discuss the practical task and any challenges faced by students.
- Provide feedback and suggest additional exercises to reinforce learning.
