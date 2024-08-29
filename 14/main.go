package main

func soma(a, b *int) int {
	*a = 100 // Change the value of the variable 'a' through the pointer 'a'
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(soma(&minhaVar1, &minhaVar2)) // Pass the memory address of the variables 'minhaVar1' and 'minhaVar2' to the function 'soma'
}
