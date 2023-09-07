package main

import "fmt"

func main() {
	feetToMetre := 0.3048
	fmt.Println("Convert From Feet to Metre. Enter a measurement in Feet")
	var userInput float64
	_, err := fmt.Scanf("%f", &userInput)
	if err != nil {
		return
	}
	fmt.Printf("%.2f feet is equals to %.2f metre", userInput, feetToMetre*userInput)
}
