package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"os"
	"strconv"
)

func main() {
	var passwordLength int
	var password string
	const pool = "ABCDEFGHIJKLMNOPQRSTUVXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" +
		"*&^%$#@!()-`~_+=[]{};:/'\"\\?/"

	rand.Int(rand.Reader, big.NewInt(time.Now().UnixNano()))

	while (os.Args) {
		passwordLength, err := strconv.Atoi((os.Args[1]))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	passwordLength := 12

	for i := 0; i < passwordLength; i++ {
		// Cast to type int64 (reduce unintended behaviour)
		var lenPool int64 = int64(len(pool))
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(lenPool))
		character := pool[int64(randomNum.Int64())]
		password += string(character)
	}

	fmt.Println("Generated password is: ", password)
}
