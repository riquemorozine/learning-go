package main

import "fmt"

type NameType string

var name NameType = "Rique"

func main() {
	fmt.Printf("The type of name is %T", name) // Print the type of the variable 'name'
}
