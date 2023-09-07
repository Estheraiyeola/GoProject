package main

import "fmt"

func main() {
	number := 0
	largest := 0
	for counter := 0; counter < 10; counter++ {
		fmt.Println("Enter a number: ")
		fmt.Scanf("%d", &number)
		if number > largest {
			largest = number
		}
	}
	fmt.Println(largest)
}
