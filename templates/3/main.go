package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	t := template.Must(template.New("content.html").ParseFiles("templates/3/content.html"))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Python", 45},
		{"Java", 60},
	})

	if err != nil {
		panic(err)
	}
}
