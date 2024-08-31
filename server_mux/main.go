package main

import "net/http"

func main() {
	mux := http.NewServeMux() // Create a new ServeMux

	mux.HandleFunc("/", hello{message: "Hello, World!"}.ServeHTTP)

	http.ListenAndServe(":8080", mux) // You can pass the ServeMux as the second parameter
	// The difference when you use a NewServerMux is that you can create multiple routes and different servers with different ServeMux
	// http.ListenAndServe(":8080", nil)  If you pass nil as the second parameter, the default ServeMux will be used
}

type hello struct {
	message string
}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.message))
}
