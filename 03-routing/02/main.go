/*
Take the previous program in the previous folder and change it so that:
- a template is parsed and served
- you pass data into the template
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
	http.HandleFunc("/me", meFunc)

	http.ListenAndServe(":8080", nil)

}
