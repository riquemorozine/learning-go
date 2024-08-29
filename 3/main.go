package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New(" The sum is greater than 50")
	} // If the sum of the two numbers is greater than 50, return an error

	return a + b, nil
} // Create a function that receives two integers and returns the sum of them

func main() {
	res, err := sum(1, 2)

	if err != nil {
		fmt.Println(res)
	}

	fmt.Println(res)
} // Call the function sum and print the result of the sum
