package main

import "fmt"

type Conta struct {
	saldo float64
}

func newConta() *Conta {
	return &Conta{0}
} // Create a new instance of the struct 'Conta'

func (c *Conta) simular(valor float64) float64 { // When we add a "*" before the type, we are saying that we are working with a pointer to the struct
	// This way, we can change the value of the struct
	c.saldo += valor

	// When we do not add a "*", we are working with the struct itself like a copy

	fmt.Printf("O saldo atual é de: %v ", c.saldo)

	return c.saldo
} // Create a method 'simular' for the struct 'Conta' that adds the value to the balance and prints the current balance

func main() {
	conta := Conta{100.0}

	conta.simular(100.0) // Call the method 'simular' for the variable 'conta' with the value 100.0

	fmt.Printf("O valor da struct é %v", conta.saldo)

}
