package main

import (
	"fmt"
)

func isPrime(nmbr int) bool {
	var factors int = 0

	for i := 1; i <= nmbr; i++ {
		if nmbr%i == 0 {
			factors++
		}
	}

	// prime numbers have only 2 factors (1 and itself)
	if factors == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	// prompt user for number
	var nmbr int

	fmt.Print("Enter the number to check: ")
	_, err := fmt.Scanln(&nmbr)

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	/**
	* find the number of factors & display the result
	 */
	if isPrime(nmbr) {
		fmt.Printf("%d is a prime number\n", nmbr)
	} else {
		fmt.Printf("%d is not a prime number\n", nmbr)
	}
}
