package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument.")
		return
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Couldn't get an argument: ", err)
	}

	fmt.Println("Multiples of 3 or 5")

	sum := 0
	for i := 0; i < input; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	fmt.Println("Sum: ", sum)
}
