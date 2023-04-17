/*
# ListenAndServe on port 8080 of localhost

For the default route "/"
Have a func called "foo"
which writes to the response "foo ran"

For the route "/dog/"
Have a func called "dog"
which parses a template called "dog.gohtml"
and writes to the response "<h1>This is from dog</h1>"
and also shows a picture of a dog when the template is executed.

Use "http.ServeFile"
to serve the file "dog.jpeg"
*/
package main

import (
	"io"
	"net/http"
	"text/template"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo ran")
}

func dog(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl := *template.Must(template.ParseFiles("dog.gohtml"))
	tpl.Execute(res, "This is from dog.")
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "dog.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}
