package main

import "fmt"

func main() {
	for index := 0; index <= 10; index++ {
		if index%2 == 0 {
			fmt.Println(index, "even")
		} else {
			fmt.Println(index, "odd")
		}
	}
}
