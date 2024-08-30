package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	if err != nil {
		panic(err)
	}

	defer req.Body.Close() // Close the connection after the function ends

	res, err := io.ReadAll(req.Body) // Read the response body
	if err != nil {
		panic(err)
	}

	fmt.Println(string(res))
}
