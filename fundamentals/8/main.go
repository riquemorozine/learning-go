package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 100, 300, 1000))
}

func sum(numeros ...int) int { // The '...' operator allows the function to receive a variable number of arguments
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
