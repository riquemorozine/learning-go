package main

import "fmt"

func main() {
	salarios := map[string]int{"Rique": 1000, "Michael": 2000}

	fmt.Println(salarios["Rique"])
	delete(salarios, "Rique")
	fmt.Println(salarios["Michael"])
	salarios["Alexia"] = 3000

	fmt.Println(salarios)

	sal := make(map[string]int) // Create a map with the type 'map[string]int'
	sal1 := map[string]int{}    // Same as above

	sal["Rique"] = 1000
	sal1["Rique"] = 1000

	for nome, salario := range sal {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	} // Print the name and the salary of each employee
}
