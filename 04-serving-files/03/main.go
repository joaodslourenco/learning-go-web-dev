/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use only this:

``` Go

	fs := http.FileServer(http.Dir("public"))

```

Hint: look to see what type FileServer returns, then think it through.
*/
package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("starting-files/templates/index.gohtml"))
}

func dogs(res http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln("Error while executing the template:", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("starting-files/public"))

	http.HandleFunc("/", dogs)
	http.Handle("/pics/", fs)

	http.ListenAndServe(":8080", nil)
}
