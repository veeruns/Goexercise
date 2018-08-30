package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	sum := 0.0
	for _, val := range args {
		num, err := strconv.ParseFloat(val, 64)
		if err == nil {
			fmt.Printf("Value is %f\n", num)
			sum = sum + num
		}
	}
	fmt.Printf("Sum of all the integers is %f\n", sum)
}
