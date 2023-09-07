package main

import "fmt"

func main() {
	const weeklyPay = 200
	itemCounter := 0
	total := 0.0
	value := 0.0
	salesPersonsPay := 0.0
	fmt.Println("Enter the value of the item:")
	fmt.Scanf("%f", &value)
	for value != -1 {
		total = total + value
		itemCounter++
		fmt.Println("Enter the value of the item:")
		fmt.Scanf("%f", &value)
	}
	if itemCounter > 0 {
		salesPersonsPay = weeklyPay + 0.09*total
		fmt.Println(salesPersonsPay)
	}
}
