package main

import "fmt"

func main() {
	for index := 0; index <= 100; index++ {
		if index%3 == 0 && index%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if index%5 == 0 {
			fmt.Println("Buzz")
		} else if index%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(index)
		}
	}
}
