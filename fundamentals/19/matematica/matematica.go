package matematica

func Sum[T int | float64](a, b T) T { // if Sum is typed with "S" in uppercase, it will be exported, otherwise it will be private
	return a + b
}

type Carro struct {
	Marca string // if Marca is typed with "M" in uppercase, it will be exported, otherwise it will be private
}
