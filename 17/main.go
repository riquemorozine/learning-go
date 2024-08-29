package main

import "fmt"

func main() {
	var minhaVar interface{} = "Rique Morozine"

	res, ok := minhaVar.(int) // Type assertion
	// If the assertion is successful, the variable 'ok' will be true and the variable 'res' will be the value of the variable 'minhaVar' converted to the type 'int'

	if !ok {
		fmt.Println("Não é um inteiro")
	} // Não é um inteiro

	res2 := minhaVar.(int)

	fmt.Printf("O valor de res é %v e o resultado de ok é %v \n", res, ok) // Não é um inteiro

	fmt.Printf("O valor de res2 é %v\n", res2) // panic: interface conversion: interface {} is string, not int
}
