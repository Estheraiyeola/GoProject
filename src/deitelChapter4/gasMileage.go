package main

import "fmt"

func main() {
	var stopValue = -1
	var miles, gallon, gallonCounter, milesCounter int
	fmt.Print("Enter the miles or press -1 to quit: ")
	fmt.Scanf("%d", &miles)
	for miles != stopValue {
		milesCounter += miles
		fmt.Print("Enter the gallon used or press -1 to quit: ")
		fmt.Scanf("%d", &gallon)
		gallonCounter += gallon
		milesPerGallon := float32(miles / gallon)
		fmt.Printf("The miles per gallon for this trip is %.2f\n", milesPerGallon)
		fmt.Print("Enter the miles or press -1 to quit: ")
		fmt.Scanf("%d", &miles)
	}
	fmt.Printf("The combined miles per gallon is %d", milesCounter/gallonCounter)
}
