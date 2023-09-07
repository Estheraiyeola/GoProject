package main

import "fmt"

func main() {
	biggest := 0
	secondBiggest := 0
	number := 0
	for index := 0; index < 10; index++ {
		fmt.Println("Enter a number: ")
		fmt.Scanf("%d", &number)
		if number > biggest {
			secondBiggest = biggest
			biggest = number
		}
		if number > secondBiggest && number < biggest {
			secondBiggest = number
		}
	}
	fmt.Println("Second Biggest = ", secondBiggest)
	fmt.Println("Biggest = ", biggest)
}
