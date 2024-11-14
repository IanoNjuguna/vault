package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func find_longest_word(text string) string {
	longest_word := ""
	max_length := 0

	words := strings.Fields((text))

	for _, word := range words {
		if len(word) > max_length {
			longest_word = word
			max_length = len(word)
		}
	}

	return longest_word
}

func main() {
	var text string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a sentence or paragraph: ")
	text, _ = reader.ReadString('\n')

	longest_word := find_longest_word(text)

	fmt.Printf("The longest word is: %s\n", longest_word)
}
