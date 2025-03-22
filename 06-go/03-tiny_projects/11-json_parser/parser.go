package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string
	Age  int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Syntax: ./json-parser <json_files>")
		return
	}

	json_file := os.Args[1]
	content, err := os.ReadFile(json_file)
	if err != nil {
		fmt.Println("Error reading from file")
		return
	}

	var user user
	err = json.Unmarshal(content, &user)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Age: %d\n", user.Age)
	}
}
