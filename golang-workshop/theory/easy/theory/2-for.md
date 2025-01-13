### Workshop for Computer Science Students on Go Language: Loops

## Part 1: Lecture 

### Introduction to Loops in Go

#### What are Loops?
- Loops are control flow statements that repeatedly execute a block of code as long as a specified condition is met.
- Go has only one looping construct: the `for` loop, which can be used in various forms.

### Basic `for` Loop

#### 1. Traditional `for` Loop
The most common form of the `for` loop includes three components: the initialization, the condition, and the post (increment/decrement).

**Syntax**:
```go
for initialization; condition; post {
    // code to be executed
}
```

**Example**:
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
}
```
- **Initialization**: `i := 0` sets the starting value.
- **Condition**: `i < 10` keeps the loop running while `i` is less than 10.
- **Post**: `i++` increments `i` after each iteration.

### 2. Conditional `for` Loop
The `for` loop can also be used with a single condition, similar to a `while` loop in other languages.

**Example**:
```go
package main

import "fmt"

func main() {
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
}
```
- The loop continues as long as the condition `i < 10` is true.

### 3. Infinite `for` Loop
A `for` loop without any condition is an infinite loop, which runs forever unless interrupted by a `break` statement or some other terminating condition.

**Example**:
```go
package main

import "fmt"

func main() {
    i := 0
    for {
        if i >= 10 {
            break
        }
        fmt.Println(i)
        i++
    }
}
```
- Use `break` to exit the loop when the condition `i >= 10` is met.

### 4. `for`-`range` Loop
The `for`-`range` loop is used to iterate over elements of a collection (arrays, slices, maps, strings, channels).

**Examples**:

- **Iterating Over a Slice**:
  ```go
  package main

  import "fmt"

  func main() {
      nums := []int{1, 2, 3, 4, 5}
      for index, value := range nums {
          fmt.Println(index, value)
      }
  }
  ```

- **Iterating Over a String**:
  ```go
  package main

  import "fmt"

  func main() {
      str := "hello"
      for index, char := range str {
          fmt.Printf("Index: %d, Character: %c\n", index, char)
      }
  }
  ```

- **Iterating Over a Map**:
  ```go
  package main

  import "fmt"

  func main() {
      myMap := map[string]int{"one": 1, "two": 2, "three": 3}
      for key, value := range myMap {
          fmt.Println(key, value)
      }
  }
  ```

### Control Statements in Loops

#### 1. `break` Statement
The `break` statement terminates the loop immediately.

**Example**:
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break
        }
        fmt.Println(i)
    }
}
```

#### 2. `continue` Statement
The `continue` statement skips the rest of the current iteration and proceeds with the next iteration of the loop.

**Example**:
```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue
        }
        fmt.Println(i)
    }
}
```

### Nested Loops
You can place one loop inside another loop to create nested loops.

**Example**:
```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i = %d, j = %d\n", i, j)
        }
    }
}
```

### Summary
- The `for` loop is the only looping construct in Go, used in various forms: traditional, conditional, infinite, and `for`-`range`.
- Control statements like `break` and `continue` modify the loop's behavior.
- Loops can be nested to perform more complex iterations.

---

## Part 2: Practical Session

### Practical Task: FizzBuzz Program

**Task Description**: Write a program that prints the numbers from 1 to 100. For multiples of three, print "Fizz" instead of the number, and for multiples of five, print "Buzz". For numbers which are multiples of both three and five, print "FizzBuzz".

**Requirements**:
1. Use a `for` loop to iterate from 1 to 100.
2. Use conditional statements to check the divisibility of each number.
3. Print "Fizz" for multiples of three, "Buzz" for multiples of five, and "FizzBuzz" for multiples of both.

**Example Implementation**:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 100; i++ {
        if i%3 == 0 && i%5 == 0 {
            fmt.Println("FizzBuzz")
        } else if i%3 == 0 {
            fmt.Println("Fizz")
        } else if i%5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
    }
}
```

### Task Recommendations:
1. **Extend the Range**: Modify the program to take user input for the range instead of hardcoding 1 to 100.
2. **Optimize Conditions**: Refactor the conditions to make the code more efficient or readable.
3. **Add Custom Words**: Allow users to input their own words for multiples of three and five.

### Conclusion
- Discuss the practical task and any challenges faced by students.
- Provide feedback and suggest additional exercises to reinforce learning.

This comprehensive lesson on loops in Go will help students understand the different forms of the `for` loop, control statements, and practical applications through exercises like FizzBuzz.