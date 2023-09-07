package main

import "fmt"

func main() {
	for index := 0; index <= 100; index++ {
		if index%3 == 0 {
			fmt.Println(index)
		}
	}
}
