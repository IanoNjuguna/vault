package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	target := rand.Intn(10) + 1
	attempts := 0
	var guess int

	fmt.Println("This is a number guesser.")
	fmt.Println("I am thinking of a number between 1 and 10. Can you guess it?")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low!. Try again")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congrats! You guessed the number %d after %d attempts\n", target, attempts)
			break
		}
	}
}
