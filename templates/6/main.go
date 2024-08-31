package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"./templates/6/header.html",
		"./templates/6/content.html",
		"./templates/6/footer.html",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.New("content.html")
		t.Funcs(template.FuncMap{"ToUpper": ToUpper})
		t = template.Must(t.ParseFiles(templates...))

		err := t.Execute(w, Cursos{
			{"Go", 40},
			{"Python", 45},
			{"Java", 60},
		})

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8084", nil)

}
