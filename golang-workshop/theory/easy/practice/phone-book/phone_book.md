#### Task: Developing a "Phone Book" Program

**Task Description**: Write a simple program that will store and manage contacts in a phone book. The program should be able to add, delete, and search for contacts.

**Requirements**:
1. The program should allow the user to input commands in the console:
    - `add <name> <phone>`
    - `delete <name>`
    - `find <name>`
    - `list`
2. Contacts should be stored in a map with the name as the key and the phone number as the value.
3. Implement functions for each command.

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
    phoneBook := make(map[string]string)
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("> ")
        scanner.Scan()
        input := scanner.Text()
        command := strings.Fields(input)

        if len(command) == 0 {
            continue
        }

        switch command[0] {
        case "add":
            if len(command) != 3 {
                fmt.Println("Usage: add <name> <phone>")
                continue
            }
            name, phone := command[1], command[2]
            phoneBook[name] = phone
            fmt.Println("Added:", name)
        case "delete":
            if len(command) != 2 {
                fmt.Println("Usage: delete <name>")
                continue
            }
            name := command[1]
            delete(phoneBook, name)
            fmt.Println("Deleted:", name)
        case "find":
            if len(command) != 2 {
                fmt.Println("Usage: find <name>")
                continue
            }
            name := command[1]
            phone, exists := phoneBook[name]
            if exists {
                fmt.Println("Found:", name, "=>", phone)
            } else {
                fmt.Println("Not found:", name)
            }
        case "list":
            fmt.Println("Phone Book:")
            for name, phone := range phoneBook {
                fmt.Println(name, "=>", phone)
            }
        default:
            fmt.Println("Unknown command:", command[0])
        }
    }
}
```

### Task Recommendations:
1. **Divide the Code into Functions**: Separate the main loop into individual functions for adding, deleting, searching, and listing contacts.
2. **Error Handling**: Add error checks and user-friendly messages for incorrect input.
3. **Enhance the Program**: Add additional commands such as `update` for updating existing contacts.
