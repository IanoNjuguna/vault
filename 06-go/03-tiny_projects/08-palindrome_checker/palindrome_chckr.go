package main

/**
 * A palindrome is a word, number, phrase, or other sequence of symbols
 * that reads the same backwards as forwards, such as:
 * madam or racecar,
 * the date "22/02/2022",
 * and the sentence: "A man, a plan, a canal – Panama".
 */

import (
	"bufio"
	"fmt"
	"strings"
)

/**
 * reverse_wrd - reverses word to
 * 		check whether the result
 *		is the same as original input.
 *
 * @word: original input
 */

func reverse_wrd(word string) string {
	reversed := ""

	for i := (len(word) - 1); i >= 0; i-- {
		reversed += string(word[i])
	}

	return reversed
}

func main() {
	var word string
	fmt.Print("Enter the word to check: ")

	/// FIX
	reader := bufio.NewScanner(strings.NewReader(word))

	reversed := reverse_wrd(word)

	if reversed == word {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}
}
