package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 2, 3, 4, 5, 100, 300, 1000) * 2
	}() // Call the function sum and multiply the result by 2

	fmt.Println(total)
}

func sum(numeros ...int) int { // The '...' operator allows the function to receive a variable number of arguments
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
