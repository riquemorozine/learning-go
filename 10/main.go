package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
} // Create a struct 'Cliente' with the fields 'Nome', 'Idade' and 'Ativo'

func main() {
	rique := Cliente{"Rique", 22, true}

	rique.Ativo = false

	fmt.Println(rique)
}
