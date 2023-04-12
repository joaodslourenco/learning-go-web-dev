/*
1. Create a data structure to pass to a template which
- contains information about California hotels including Name, Address, City, Zip, Region
- region can be: Southern, Central, Northern
- can hold an unlimited number of hotels
*/

package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotels := hotels{
		{
			Name:    "Hotel California",
			Address: "42 Sunset Boulevard",
			City:    "Los Angeles",
			Zip:     "95612",
			Region:  "Southern"},
		{
			Name:    "Another Hotel",
			Address: "Different Address",
			City:    "Los Angeles",
			Zip:     "12345",
			Region:  "Southern"},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
