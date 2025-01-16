package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	description string
	completed   bool
}

func main() {
	var tasks []Task
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
			if len(command) < 2 {
				fmt.Println("Usage: add <task>")
				continue
			}
			task := strings.Join(command[1:], " ")
			tasks = append(tasks, Task{description: task})
			fmt.Println("Added:", task)
		case "complete":
			if len(command) != 2 {
				fmt.Println("Usage: complete <task_number>")
				continue
			}
			taskNumber := command[1]
			index, err := strconv.Atoi(taskNumber)
			if err != nil || index < 1 || index > len(tasks) {
				fmt.Println("Invalid task number")
				continue
			}
			tasks[index-1].completed = true
			fmt.Println("Completed task", index)
		case "list":
			fmt.Println("To-Do List:")
			for i, task := range tasks {
				status := " "
				if task.completed {
					status = "x"
				}
				fmt.Printf("[%s] %d: %s\n", status, i+1, task.description)
			}
		default:
			fmt.Println("Unknown command:", command[0])
		}
	}
}
