### Option 3: Word Count Program

**Task Description**: Write a program that reads text from the input stream and counts the number of occurrences of each word.

**Requirements**:
1. The program should read text from standard input until end-of-input.
2. Word counting should be performed using a map with words as keys and counts as values.
3. Implement a function to display the results.

**Example Implementation**:

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    wordCount := make(map[string]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    fmt.Println("Enter text (Ctrl+D to end):")
    for scanner.Scan() {
        word := strings.ToLower(scanner.Text())
        wordCount[word]++
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    fmt.Println("Word count:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

### Option 4: Random Number Generator Program

**Task Description**: Write a program that generates a specified number of random numbers within a given range and prints them.

**Requirements**:
1. The program should ask the user for the number of numbers and the range (minimum and maximum).
2. Use the `math/rand` package to generate random numbers.
3. Print the generated numbers to the screen.

**Example Implementation**:

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    var count, min, max int
    fmt.Print("Enter the number of random numbers to generate: ")
    fmt.Scan(&count)
    fmt.Print("Enter the minimum value: ")
    fmt.Scan(&min)
    fmt.Print("Enter the maximum value: ")
    fmt.Scan(&max)

    if min > max {
        fmt.Println("Minimum value cannot be greater than maximum value")
        return
    }

    fmt.Println("Generated random numbers:")
    for i := 0; i < count; i++ {
        num := rand.Intn(max-min+1) + min
        fmt.Println(num)
    }
}
```

