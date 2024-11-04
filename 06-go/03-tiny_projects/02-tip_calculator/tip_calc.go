package main

import (
	"fmt"
)

func main() {
	var bill_amount, tip_percentage float64

	fmt.Print("Enter the bill amount in USD ($): ")
	fmt.Scanln(&bill_amount)
	fmt.Print("Enter the tip percentage: ")
	fmt.Scanln(&tip_percentage)

	tip_amount := (tip_percentage / 100 * bill_amount)
	total_bill_amount := bill_amount - tip_amount

	fmt.Printf("Tip Amount: $%.2f\n", tip_amount)
	fmt.Printf("Total bill amount (tip included): $%.2f\n", total_bill_amount)
}
