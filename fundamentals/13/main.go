package main

func main() {
	a := 10
	var ponteiro *int = &a // Create a pointer 'ponteiro' to the variable 'a'
	*ponteiro = 20         // Change the value of the variable 'a' through the pointer 'ponteiro'
	println(a)
}
