package main

import "fmt"

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}

	return soma
}

func Comparar[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Rique": 1000, "Michael": 2000}
	mf := map[string]float64{"Rique": 1000.50, "Michael": 2000.50}

	m3 := map[string]MyNumber{"Rique": 1000, "Michael": 2000}

	fmt.Println(Soma(m))
	fmt.Println(Soma(mf))
	fmt.Println(Soma(m3))
	fmt.Println(Comparar(10, 10.0))
}
