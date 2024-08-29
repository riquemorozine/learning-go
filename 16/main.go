package main

import "fmt"

func main() {

	var x interface{} = 10 // Create a variable 'x' with the type 'interface{}' and the value 10

	var y interface{} = "Hello world" // Create a variable 'y' with the type 'interface{}' and the value "Hello world"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("Type: %T\n e o valor é %v \n", t, t)
}
