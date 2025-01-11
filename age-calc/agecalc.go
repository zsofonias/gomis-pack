package main

import (
	"fmt"
	"time"
)

func main() {
	var birthyear int
	fmt.Print("Enter your year of birth: ")
	fmt.Scanln(&birthyear)

	currentYear := time.Now().Year()

	age := currentYear - birthyear 

	fmt.Printf("Your age is %d years old", age)
}