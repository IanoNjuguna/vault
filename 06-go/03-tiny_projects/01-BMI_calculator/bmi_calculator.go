package main

import (
	"fmt"
	"math"
)

// Body Mass Index (BMI)
// Used to assess whether a person's weight is healthy ...
// in relation to their height

func main() {
	var weight, height float64

	fmt.Print("Enter the weight (kg): ")
	fmt.Scanln(&weight)

	fmt.Print("Enter the height (m): ")
	fmt.Scanln(&height)

	// math.Pow() calculates base-2 exponentiation in Golang
	bmi := weight/ (math.Pow(height, 2))
	fmt.Printf("Your BMI is: %.2f\n", bmi)

	/*
	 * underweight: bmi < 18.5
	 * normal weight: bmi 18.5 - 24.9
	 * overweight: bmi 25 - 29.9
	 * obese: bmi >= 30
	 */
	if bmi < 18.5 {
		fmt.Println("Category: Underweight")
	} else if bmi >= 18.5 && bmi <= 24.9 {
		fmt.Println("Category: Normal Weight")
	} else if bmi > 24.9 && bmi <= 29.9 {
		fmt.Println("Category: Over Weight")
	} else {
		fmt.Println("Category: Obese")
	}
}
