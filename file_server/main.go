package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()

	println(http.Dir("./public"))

	mux.Handle("/", fileServer)

	log.Fatal(http.ListenAndServe(":8081", mux))
}
