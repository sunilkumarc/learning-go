package main

import (
	"fmt"
)

func main() {
	slice := []int{}

	for i := 0; i <= 10; i++ {
		slice = append(slice, i)
	}

	for _, number := range slice {
		if number%2 == 0 {
			fmt.Println(number, " is even")
		} else {
			fmt.Println(number, " is odd")
		}
	}
}
