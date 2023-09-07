package main

import "fmt"

func main() {
	accountNumber := 1
	balance := 0
	charges := 0
	credit := 0
	creditLimit := 0
	fmt.Println("Enter your account number: ")
	fmt.Scanf("%d", &accountNumber)
	for accountNumber != 0 {
		fmt.Println("Enter your beginning balance: ")
		fmt.Scanf("%d", &balance)
		fmt.Println("Enter your charges: ")
		fmt.Scanf("%d", &charges)
		fmt.Println("Enter your credit: ")
		fmt.Scanf("%d", &credit)
		fmt.Println("Enter your creditLimit: ")
		fmt.Scanf("%d", &creditLimit)

		newBalance := balance + charges - credit
		fmt.Println("The new balance is ", newBalance)

		if newBalance > creditLimit {
			fmt.Println("Credit Limit Exceeded")
		}
	}
}
