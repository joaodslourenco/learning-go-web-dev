/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/" "/dog/" "/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"io"
	"net/http"
)

func dogFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func meFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "João Lourenço")
}

func main() {

	http.HandleFunc("/dog", dogFunc)
	http.HandleFunc("/me", meFunc)

	http.ListenAndServe(":8080", nil)
}
