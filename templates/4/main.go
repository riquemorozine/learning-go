package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles("./templates/4/content.html"))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 45},
			{"Java", 60},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8082", nil)

}
