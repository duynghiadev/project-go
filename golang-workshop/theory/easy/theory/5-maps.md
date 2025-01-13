### Workshop for Computer Science Students on Go Language: Maps

## Part 1: Lecture

### Introduction to Maps in Go

#### What are Maps?
- Maps are Go's built-in associative data type (also known as hash tables or dictionaries in other languages).
- Maps store key-value pairs and provide efficient retrieval of values based on their corresponding keys.

### Declaring and Initializing Maps

#### 1. Declaration
Maps can be declared using the `make` function.

**Example**:
```go
package main

import "fmt"

func main() {
    var m map[string]int // declares a map without initializing it
    fmt.Println(m) // prints map[]
}
```

#### 2. Initialization
Maps are usually initialized using the `make` function, which allocates and initializes a map.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int) // creates an empty map
    fmt.Println(m) // prints map[]
}
```

### Adding and Retrieving Elements
Elements can be added to a map using the assignment operator. Elements are retrieved using the key.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1 // adds a key-value pair to the map
    fmt.Println(m["one"]) // prints 1
}
```

### Updating Elements
To update an element in a map, assign a new value to an existing key.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1
    m["one"] = 11 // updates the value for the key "one"
    fmt.Println(m["one"]) // prints 11
}
```

### Deleting Elements
Elements can be deleted from a map using the `delete` function.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1
    delete(m, "one") // removes the key "one" from the map
    fmt.Println(m) // prints map[]
}
```

### Checking for Key Existence
You can check if a key exists in a map by using the value, ok idiom.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1
    value, ok := m["one"]
    if ok {
        fmt.Println("Key exists with value", value) // prints Key exists with value 1
    } else {
        fmt.Println("Key does not exist")
    }
}
```

### Iterating Over Maps
You can iterate over map elements using the `for`-`range` loop.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1
    m["two"] = 2
    m["three"] = 3

    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

### Map Length
The `len` function returns the number of key-value pairs in a map.

**Example**:
```go
package main

import "fmt"

func main() {
    m := make(map[string]int)
    m["one"] = 1
    m["two"] = 2
    fmt.Println("Length:", len(m)) // prints Length: 2
}
```

### Nested Maps
Maps can contain other maps, creating nested structures.

**Example**:
```go
package main

import "fmt"

func main() {
    nestedMap := make(map[string]map[string]int)
    nestedMap["first"] = make(map[string]int)
    nestedMap["first"]["one"] = 1
    fmt.Println(nestedMap) // prints map[first:map[one:1]]
}
```

### Summary
- Maps in Go are dynamic collections that store key-value pairs.
- You can add, retrieve, update, delete, and check for elements in maps.
- Maps can be iterated using the `for`-`range` loop.
- The `len` function returns the number of elements in a map.
- Maps can be nested to create complex data structures.

---

## Part 2: Practical Session

### Practical Task: Working with Maps

**Task Description**: Write a program that performs the following tasks:
1. Declares and initializes a map of string keys and integer values.
2. Adds several key-value pairs to the map.
3. Updates the value for an existing key.
4. Deletes a key-value pair from the map.
5. Checks if a key exists in the map and prints an appropriate message.
6. Iterates over the map and prints all key-value pairs.
7. Finds and prints the key with the maximum value.

**Requirements**:
1. Use a dynamic map.
2. Implement functions for adding, updating, deleting, checking, iterating, and finding the maximum value.
3. Use loops and conditional statements as needed.

**Example Implementation**:

```go
package main

import "fmt"

// Function to add key-value pairs to the map
func add(m map[string]int, key string, value int) {
    m[key] = value
}

// Function to update the value for an existing key
func update(m map[string]int, key string, value int) {
    if _, ok := m[key]; ok {
        m[key] = value
    } else {
        fmt.Printf("Key %s does not exist\n", key)
    }
}

// Function to delete a key-value pair from the map
func remove(m map[string]int, key string) {
    delete(m, key)
}

// Function to check if a key exists in the map
func exists(m map[string]int, key string) bool {
    _, ok := m[key]
    return ok
}

// Function to find the key with the maximum value
func maxKey(m map[string]int) string {
    maxKey := ""
    maxValue := 0
    for key, value := range m {
        if value > maxValue {
            maxValue = value
            maxKey = key
        }
    }
    return maxKey
}

func main() {
    m := make(map[string]int)

    // Adding key-value pairs
    add(m, "one", 1)
    add(m, "two", 2)
    add(m, "three", 3)
    fmt.Println("Map after adding elements:", m)

    // Updating a value
    update(m, "two", 22)
    fmt.Println("Map after updating 'two':", m)

    // Deleting a key-value pair
    remove(m, "one")
    fmt.Println("Map after removing 'one':", m)

    // Checking if a key exists
    if exists(m, "two") {
        fmt.Println("Key 'two' exists")
    } else {
        fmt.Println("Key 'two' does not exist")
    }

    // Iterating over the map
    fmt.Println("Iterating over map:")
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }

    // Finding the key with the maximum value
    maxKey := maxKey(m)
    fmt.Printf("Key with the maximum value: %s\n", maxKey)
}
```

### Task Recommendations:
1. **Experiment with Different Data**: Try different keys and values.
2. **Optimize Functions**: Refactor functions for efficiency.
3. **Extend Functionality**: Add functions to find the minimum value and to sort the map by values.

### Conclusion
- Discuss the practical task and any challenges faced by students.
- Provide feedback and suggest additional exercises to reinforce learning.

This comprehensive lesson on maps in Go will help students understand map declaration, initialization, modification, and manipulation, along with practical applications through hands-on exercises.