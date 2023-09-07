package main

import (
	"fmt"
	"math"
)

func main() {
	for index := 1; index < 5; index++ {
		if index == 1 {
			fmt.Print("N\t\t")
		} else {
			fmt.Printf("N%d\t\t", index)
		}
	}
	fmt.Println()
	for index := 1; index < 6; index++ {
		for idx := 1; idx < 5; idx++ {
			fmt.Print(math.Pow(float64(index), float64(idx)), "\t\t")
		}
		fmt.Println()
	}
}
