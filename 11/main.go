package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
} // Create a struct 'Endereco' with the fields 'Logadouro', 'Numero', 'Cidade' and 'Estado'

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
} // Create a struct 'Cliente' with the fields 'Nome', 'Idade', 'Ativo' and 'Endereco'

func main() {
	rique := Cliente{
		"Rique",
		22,
		true,
		Endereco{
			"Rua 1",
			10,
			"São Paulo",
			"SP",
		},
	} // Create a new variable 'rique' with the type 'Cliente' and the values "Rique", 22, true, Endereco{"Rua 1", 10, "São Paulo", "SP"}

	rique.Ativo = false

	fmt.Println(rique)
}
