package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5}

	for _, b := range numeros {
		println(b)
	}

	x := 0
	for x < 100 {
		println(x)
		x++
	}
}
