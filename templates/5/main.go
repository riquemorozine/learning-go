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
	templates := []string{
		"./templates/5/header.html",
		"./templates/5/content.html",
		"./templates/5/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("content.html").ParseFiles(templates...))
		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 45},
			{"Java", 60},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8083", nil)

}
