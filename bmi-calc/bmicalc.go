package main

import (
	"fmt"
)

func main() {
	var inputWeight, inputHeight float64
	var bmi float64
	var bmiCategory string

	fmt.Print("Enter the weight(Kgs): ")
	// fmt.Scanln(&inputWeight)
	_, err := fmt.Scanf("%f", &inputWeight)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.", err)
		return
	}

	
	fmt.Print("Enter the height(Cm): ")
	_, err2 := fmt.Scanf("%f", &inputHeight)
	if err2 != nil {
		fmt.Println("Invalid input. Please enter a number.", err)
		return
	}

	heightInMeters := inputHeight/100
	
	bmi = inputWeight / (heightInMeters * heightInMeters)

	if bmi < 18.5 {
		bmiCategory = "Underweight"
	} else if bmi < 25 {
        bmiCategory = "Normal weight"
    } else if bmi < 30 {
        bmiCategory = "Overweight"
    } else {
        bmiCategory = "Obese"
    }

	// fmt.Printf("The BMI is %f", bmi)
	fmt.Printf("BMI: %.2f\n", bmi)
    fmt.Println("BMI Category:", bmiCategory)
	
}