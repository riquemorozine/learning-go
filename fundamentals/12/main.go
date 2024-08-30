package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
} // Create a struct 'Endereco' with the fields 'Logadouro', 'Numero', 'Cidade' and 'Estado'

type Pessoa interface {
	Desativar()
} // Create an interface 'Pessoa' with the method 'Desativar' Interface can only have method signatures

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
} // Create a struct 'Cliente' with the fields 'Nome', 'Idade', 'Ativo' and 'Endereco'

type Empresa struct {
	Nome string
}

func (c Cliente) Desativar() {
	c.Ativo = false

	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func (e Empresa) Desativar() {} // Implement the method 'Desativar' for the struct 'Empresa'

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

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

	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa) // Call the function 'Desativacao' with the variable 'minhaEmpresa'
	Desativacao(rique)        // Call the function 'Desativacao' with the variable 'rique'
}
