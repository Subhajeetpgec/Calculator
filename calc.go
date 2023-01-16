package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Please provide an operation and two numbers For Example 1 + 5")
		return
	}

	op := os.Args[1]  // slice of strings
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid first number")
		return
	}
	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Invalid second number")
		return
	}

	switch op {
	case "+":
		fmt.Printf("%v + %v = %v\n", a, b, a+b)
	case "-":
		fmt.Printf("%v - %v = %v\n", a, b, a-b)
	case "*":
		fmt.Printf("%v * %v = %v\n", a, b, a*b)
	case "/":
		if b == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		fmt.Printf("%v / %v = %v\n", a, b, a/b)
	default:
		fmt.Println("Invalid operation")
	}
}
