/*
Take the previous program and change it so that:
func main uses http.Handle instead of http.HandleFunc
Contstraint: Do not change anything outside of func main
*/
package main

import (
	"log"
	"net/http"
	"text/template"
)

type person struct {
	FName string
	LName string
	Age   int
	Job   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func meFunc(res http.ResponseWriter, req *http.Request) {
	joao := person{
		"João", "Lourenço", 24, "Developer",
	}

	err := tpl.Execute(res, joao)
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	http.Handle("/me", http.HandlerFunc(meFunc))

	http.ListenAndServe(":8080", nil)

}
