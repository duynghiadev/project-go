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
