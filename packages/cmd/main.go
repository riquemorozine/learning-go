package main

import (
	"fmt"
	"github.com/riquemorozine/learning-go/packages/math"
)

func main() {
	sum := math.Math{A: 10, B: 20}

	fmt.Println(sum.Sum())
	fmt.Println(sum.Multiply())
}
