package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	passwordLength := 12
	var password string
	const pool = "ABCDEFGHIJKLMNOPQRSTUVXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"*&^%$#@!()-`~_+=[]{};:/'\"\\?/"

	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < passwordLength; i++ {
		character := pool[rand.Intn(len(pool))]
		password += string(character)
	}

	fmt.Println("Generated password is: ", password)
}
