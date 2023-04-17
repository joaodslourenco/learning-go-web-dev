/*
# Serve the files in the "starting-files" folder

## To get your images to serve, use:

``` Go

	func StripPrefix(prefix string, h Handler) Handler
	func FileServer(root FileSystem) Handler

```

Constraint: you are not allowed to change the route being used for images in the template file
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
		log.Fatalln("Error while executing template:", err)
	}
}

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("starting-files/public"))))
	http.ListenAndServe(":8080", nil)
}
