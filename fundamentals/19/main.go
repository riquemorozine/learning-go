package main

import (
	"LearningGo/fundamentals/19/matematica"
	"fmt"
)

func main() {
	conta := matematica.Sum(10, 20)
	carro := matematica.Carro{"Fiat"}

	fmt.Println(carro.Marca)
	fmt.Printf("a soma de %d com %d é %d\n", 10, 20, conta)
}
