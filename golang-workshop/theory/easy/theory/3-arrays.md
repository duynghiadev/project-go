### Workshop for Computer Science Students on Go Language: Arrays

## Part 1: Lecture 

### Introduction to Arrays in Go

#### What are Arrays?
- An array is a collection of elements of the same type placed in contiguous memory locations.
- Arrays have a fixed size that is defined at the time of declaration.

### Declaring and Initializing Arrays

#### 1. Declaration
Arrays can be declared by specifying the type of elements and the number of elements required by an array.

**Syntax**:
```go
var arrayName [size]Type
```

**Example**:
```go
package main

import "fmt"

func main() {
    var arr [5]int // declares an array of 5 integers
    fmt.Println(arr) // prints [0 0 0 0 0]
}
```

- `arr` is an array of 5 integers. By default, all elements are initialized to the zero value of the element type (`0` for `int`).

#### 2. Initialization
Arrays can be initialized at the time of declaration.

**Example**:
```go
package main

import "fmt"

func main() {
    var arr [5]int = [5]int{1, 2, 3, 4, 5} // declares and initializes an array
    fmt.Println(arr) // prints [1 2 3 4 5]
}
```

#### 3. Shorthand Declaration
You can use the shorthand declaration to initialize an array without specifying the size explicitly.

**Example**:
```go
package main

import "fmt"

func main() {
    arr := [...]int{1, 2, 3, 4, 5} // the size is inferred from the number of elements
    fmt.Println(arr) // prints [1 2 3 4 5]
}
```

### Accessing Array Elements
Array elements can be accessed using the index. The index of the first element is `0`.

**Example**:
```go
package main

import "fmt"

func main() {
    var arr [5]int
    arr[0] = 1 // assigns the value 1 to the first element
    fmt.Println(arr[0]) // prints 1
}
```

### Looping Through Arrays
You can use loops to iterate over array elements.

#### Using `for` Loop
**Example**:
```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
```

#### Using `for`-`range` Loop
**Example**:
```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    for index, value := range arr {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### Multi-dimensional Arrays
Go supports multi-dimensional arrays. A common use is the 2-dimensional array, also known as a matrix.

**Example**:
```go
package main

import "fmt"

func main() {
    var matrix [2][3]int
    matrix[0][0] = 1
    matrix[0][1] = 2
    matrix[0][2] = 3
    matrix[1][0] = 4
    matrix[1][1] = 5
    matrix[1][2] = 6
    fmt.Println(matrix)
}
```

### Array Length
You can get the length of an array using the `len` function.

**Example**:
```go
package main

import "fmt"

func main() {
    arr := [5]int{1, 2, 3, 4, 5}
    fmt.Println("Length of array:", len(arr)) // prints 5
}
```

### Summary
- Arrays in Go have a fixed size and hold elements of the same type.
- Elements can be accessed and modified using indices.
- You can iterate over arrays using `for` loops and `for`-`range` loops.
- Go supports multi-dimensional arrays.
- The `len` function returns the length of an array.

---

## Part 2: Practical Session

### Practical Task: Working with Arrays

**Task Description**: Write a program that performs the following tasks:
1. Declares and initializes an array of integers.
2. Computes and prints the sum of all elements in the array.
3. Finds and prints the maximum and minimum elements in the array.
4. Reverses the array and prints the reversed array.

**Requirements**:
1. Use a fixed-size array.
2. Implement functions for sum, maximum, minimum, and reverse operations.
3. Use loops to iterate over array elements.

**Example Implementation**:

```go
package main

import "fmt"

// Function to compute the sum of array elements
func sum(arr [5]int) int {
    total := 0
    for _, value := range arr {
        total += value
    }
    return total
}

// Function to find the maximum element in the array
func max(arr [5]int) int {
    maxValue := arr[0]
    for _, value := range arr {
        if value > maxValue {
            maxValue = value
        }
    }
    return maxValue
}

// Function to find the minimum element in the array
func min(arr [5]int) int {
    minValue := arr[0]
    for _, value := range arr {
        if value < minValue {
            minValue = value
        }
    }
    return minValue
}

// Function to reverse the array
func reverse(arr [5]int) [5]int {
    var reversed [5]int
    for i, value := range arr {
        reversed[len(arr)-1-i] = value
    }
    return reversed
}

func main() {
    arr := [5]int{5, 3, 8, 1, 2}
    fmt.Println("Array:", arr)

    sumValue := sum(arr)
    fmt.Println("Sum:", sumValue)

    maxValue := max(arr)
    fmt.Println("Max:", maxValue)

    minValue := min(arr)
    fmt.Println("Min:", minValue)

    reversedArray := reverse(arr)
    fmt.Println("Reversed Array:", reversedArray)
}
```

### Task Recommendations:
1. **Modify Array Size**: Experiment with different array sizes and values.
2. **Optimize Functions**: Try optimizing the functions to handle larger arrays efficiently.
3. **Extend Functionality**: Add functions to compute the average of the elements and to sort the array.

### Conclusion
- Discuss the practical task and any challenges faced by students.
- Provide feedback and suggest additional exercises to reinforce learning.

This comprehensive lesson on arrays in Go will help students understand array declaration, initialization, access, iteration, and manipulation, along with practical applications through hands-on exercises.