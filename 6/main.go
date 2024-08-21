package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len=%d cap=%d, %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d, %v\n", len(s[:0]), cap(s[:0]), s[:0]) // Slice the first element of the slice

	fmt.Printf("len=%d cap=%d, %v\n", len(s[:4]), cap(s[:4]), s[:4]) // Slice the first four elements of the slice

	fmt.Printf("len=%d cap=%d, %v\n", len(s[2:]), cap(s[2:]), s[2:]) // Slice the first two elements of the slice

	s = append(s, 2, 5, 6, 7, 8, 9) // append elements to the slice

	fmt.Printf("len=%d cap=%d, %v\n", len(s), cap(s), s) // print the length, capacity and the slice

	// When the capacity of the slice is reached, the append function creates a new array with double the capacity of the original array and copies the elements from the original array to the new array.
}
