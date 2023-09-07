package main

import "fmt"

func main() {
	tax := 0.0
	fmt.Println("What's your name? ")
	name := ""
	fmt.Scanf("%s", &name)
	for name != "" {
		fmt.Println("How much do you earn per year?")
		earnings := 0.0
		fmt.Scanf("%f", &earnings)
		if earnings <= 30000 && earnings > 0 {
			tax = earnings * 0.15
		}
		if earnings > 30000 {
			tax = earnings * 0.20
		}
		fmt.Println("Tax is, ", tax)
		fmt.Println("What's your name? ")
		name = ""
		fmt.Scanf("%s", &name)
	}
	if name == "" {
		fmt.Println("Empty... Bye")
	}
}
