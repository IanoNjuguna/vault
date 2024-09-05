package main

/*
 * a function can take zero or more arguments.
 * add() takes two parameters of type int.
 *
 * in golang, type comes after the variable name
 */

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
