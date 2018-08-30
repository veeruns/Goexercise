package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	sum := 0
	for _, val := range args {
		num, err := strconv.Atoi(val)
		if err == nil {
			fmt.Printf("Value is %d\n", num)
			sum = sum + num
		}
	}
	fmt.Printf("Sum of all the integers is %d\n", sum)
}
