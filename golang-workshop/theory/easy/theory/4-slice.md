### Workshop for Computer Science Students on Go Language: Slices

---

## Part 1: Lecture

### Introduction to Slices in Go

#### What are Slices?
- Slices are a flexible, powerful built-in feature in Go that provide a more convenient and dynamic way to work with sequences of data than arrays.
- Unlike arrays, slices have a dynamic size. They are essentially references to arrays.

### Declaring and Initializing Slices

#### 1. Declaration
Slices can be declared using the `var` keyword or shorthand declaration.

**Example**:
```go
package main

import "fmt"

func main() {
    var s []int // declares a slice of integers
    fmt.Println(s) // prints []
}
```

#### 2. Initialization
Slices can be initialized using a composite literal.

**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3} // declares and initializes a slice
    fmt.Println(s) // prints [1 2 3]
}
```

### Using `append` to Add Elements
The `append` function is used to add elements to a slice. It returns a new slice with the added elements.

**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    s = append(s, 4) // appends 4 to the slice
    fmt.Println(s) // prints [1 2 3 4]
}
```

### Accessing and Modifying Slice Elements
You can access and modify slice elements using indices.

**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    fmt.Println(s[0]) // prints 1
    s[0] = 10
    fmt.Println(s) // prints [10 2 3]
}
```

### Slicing Slices
You can create a new slice from an existing slice by specifying a range of indices.

**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    s1 := s[1:4] // creates a slice from index 1 to 3 (excluding 4)
    fmt.Println(s1) // prints [2 3 4]
}
```

### Length and Capacity of Slices
Slices have a length and a capacity. The length is the number of elements in the slice. The capacity is the number of elements in the underlying array, counting from the first element in the slice.

**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    fmt.Println("Length:", len(s)) // prints Length: 5
    fmt.Println("Capacity:", cap(s)) // prints Capacity: 5
}
```

### Creating Slices with `make`
The `make` function is used to create slices with a specified length and capacity.

**Example**:
```go
package main

import "fmt"

func main() {
    s := make([]int, 5, 10) // creates a slice with length 5 and capacity 10
    fmt.Println("Slice:", s) // prints Slice: [0 0 0 0 0]
    fmt.Println("Length:", len(s)) // prints Length: 5
    fmt.Println("Capacity:", cap(s)) // prints Capacity: 10
}
```

### Iterating Over Slices
You can use loops to iterate over slice elements.

#### Using `for` Loop
**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}
```

#### Using `for`-`range` Loop
**Example**:
```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3, 4, 5}
    for index, value := range s {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}
```

### Copying Slices
You can copy one slice to another using the `copy` function.

**Example**:
```go
package main

import "fmt"

func main() {
    s1 := []int{1, 2, 3, 4, 5}
    s2 := make([]int, len(s1)) // create a slice with the same length as s1
    copy(s2, s1) // copy elements of s1 to s2
    fmt.Println("s1:", s1) // prints s1: [1 2 3 4 5]
    fmt.Println("s2:", s2) // prints s2: [1 2 3 4 5]
}
```

### Summary
- Slices in Go are dynamic and more flexible than arrays.
- You can append elements to slices, access and modify elements, and create sub-slices.
- Slices have length and capacity, which can be managed using the `len` and `cap` functions.
- The `make` function creates slices with specified length and capacity.
- The `copy` function copies elements from one slice to another.

---

## Part 2: Practical Session

### Practical Task: Working with Slices

**Task Description**: Write a program that performs the following tasks:
1. Declares and initializes a slice of integers.
2. Appends new elements to the slice.
3. Computes and prints the sum of all elements in the slice.
4. Removes an element at a specified index from the slice.
5. Finds and prints the maximum and minimum elements in the slice.

**Requirements**:
1. Use a dynamic slice.
2. Implement functions for sum, removal, maximum, and minimum operations.
3. Use loops to iterate over slice elements.

**Example Implementation**:

```go
package main

import "fmt"

// Function to compute the sum of slice elements
func sum(slice []int) int {
    total := 0
    for _, value := range slice {
        total += value
    }
    return total
}

// Function to find the maximum element in the slice
func max(slice []int) int {
    maxValue := slice[0]
    for _, value := range slice {
        if value > maxValue {
            maxValue = value
        }
    }
    return maxValue
}

// Function to find the minimum element in the slice
func min(slice []int) int {
    minValue := slice[0]
    for _, value := range slice {
        if value < minValue {
            minValue = value
        }
    }
    return minValue
}

// Function to remove an element at a specified index
func remove(slice []int, index int) []int {
    return append(slice[:index], slice[index+1:]...)
}

func main() {
    s := []int{5, 3, 8, 1, 2}
    fmt.Println("Initial Slice:", s)

    s = append(s, 10) // append an element to the slice
    fmt.Println("After Append:", s)

    sumValue := sum(s)
    fmt.Println("Sum:", sumValue)

    maxValue := max(s)
    fmt.Println("Max:", maxValue)

    minValue := min(s)
    fmt.Println("Min:", minValue)

    s = remove(s, 2) // remove the element at index 2
    fmt.Println("After Removal:", s)
}
```

### Task Recommendations:
1. **Experiment with Different Slices**: Try different initial values and slice operations.
2. **Optimize Functions**: Refactor functions to handle larger slices efficiently.
3. **Extend Functionality**: Add functions to find the average of the elements and to sort the slice.

### Conclusion
- Discuss the practical task and any challenges faced by students.
- Provide feedback and suggest additional exercises to reinforce learning.

This comprehensive lesson on slices in Go will help students understand slice declaration, initialization, modification, and manipulation, along with practical applications through hands-on exercises.