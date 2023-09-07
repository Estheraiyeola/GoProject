package main

import "fmt"

func main() {
	var fahrenheit float64
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)
	var celcius float64 = (fahrenheit - 32) * 5 / 9
	fmt.Println("Temperature in celcius is ", celcius)
}
