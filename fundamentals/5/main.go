package main

import "fmt"

func main() {
	var a [3]int

	a[0] = 10
	a[1] = 20
	a[2] = 30

	println(len(a))

	for i, v := range a {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	} // Print the index and the value of each element in the array
}
