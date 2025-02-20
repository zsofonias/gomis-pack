package main

import "fmt"

func main() {
	fmt.Println("Tip Calculator")

	var billAmount, tipPercentage float64

	fmt.Print("Enter the bill amount: ")

	_, err := fmt.Scanf("%f", &billAmount)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.", err)
		return
	}

	fmt.Print("Enter the tip percentage: ")
	_, err = fmt.Scanf("%f", &tipPercentage)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.", err)
		return
	}

	tipAmount := billAmount * (tipPercentage / 100)
	totalAmount := billAmount + tipAmount

	fmt.Printf("Tip amount: $%.2f\n", tipAmount)
	fmt.Printf("Total bill amount: $%.2f\n", totalAmount)
}