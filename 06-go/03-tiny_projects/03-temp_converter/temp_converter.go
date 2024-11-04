package main

import (
	"fmt"
)

func celsius_to_fahrenheit(celsius float64) float64 {
	// (celsius * 9/5) + 32
	return (celsius * 9 / 5) + 32
}

func fahrenheit_to_celsius(fahrenheit float64) float64 {
	// (fahrenheit - 32) *5/9
	return (fahrenheit - 32) * 5 / 9
}

func main() {

	fmt.Println("This is a Temperature Converter")
	fmt.Println("1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("Enter the conversion type (1 or 2): ")

	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}
	if choice != 1 && choice != 2 {
		fmt.Println("Provide either 1 or 2")
		return
	}

	var temp float64
	fmt.Print("Enter the temperature value: ")
	_, err = fmt.Scanln(&temp)
	if err != nil {
		fmt.Println("Invalid Temperature Input")
		return
	}

	switch choice {
	case 1:
		result := celsius_to_fahrenheit(temp)
		fmt.Printf("%.2f C = %.2f f\n", temp, result)
	case 2:
		result := fahrenheit_to_celsius(temp)
		fmt.Printf("%.2f F = %.2f C\n", temp, result)
	}
}
