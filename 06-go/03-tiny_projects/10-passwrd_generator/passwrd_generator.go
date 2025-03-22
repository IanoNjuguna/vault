package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func main() {
	passwordLength := 12
	var password string
	const pool = "ABCDEFGHIJKLMNOPQRSTUVXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"*&^%$#@!()-`~_+=[]{};:/'\"\\?/"

	rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano()))

	for i := 0; i < passwordLength; i++ {
		// Cast to type int64 (reduce unintended behaviour)
		var lenPool int64 = int64(len(pool))
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(lenPool))
		character := pool[int64(randomNum.Int64())]
		password += string(character)
	}

	fmt.Println("Generated password is: ", password)
}
