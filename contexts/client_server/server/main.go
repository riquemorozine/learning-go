package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("Request Iniciada")

	defer log.Println("Request Finalizada")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Request Processada com sucesso")
		w.Write([]byte("Request Processada com sucesso"))
	case <-ctx.Done():
		log.Println("Request cancelada pelo client")
		http.Error(w, "Request cancelada pelo client", http.StatusRequestTimeout)
	}
}
