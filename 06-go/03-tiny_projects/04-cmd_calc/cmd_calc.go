package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("This is a CLI calculator")

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&num1)

	for err != nil {
		fmt.Println("Invalid Input. This program only accepts integer and float types")
		fmt.Print("Enter the first number: ")
		_, err = fmt.Scanln(&num1)
		continue
	}

	fmt.Print("Enter the Operator (+, -, / or *): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	for err != nil {
		fmt.Println("Invalid Input. This program only accepts integer and float types")
		fmt.Print("Enter the second number: ")
		_, err = fmt.Scanln(&num2)
		continue
	}

	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
		}
	default:
		fmt.Println("You provided an invalid operator. Provide either +, *, - or /")
	}
}
