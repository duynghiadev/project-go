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
