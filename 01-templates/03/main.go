/*
Create a data structure to pass to a template which
contains information about restaurant's menu including Breakfast, Lunch, and Dinner items
*/
package main

import (
	"os"
	"text/template"
)

type menu struct {
	Breakfast, Lunch, Dinner []string
}

type restaurant struct {
	Name string
	Menu menu
}

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurants := []restaurant{
		{
			Name: "Pizza Bob",
			Menu: menu{Breakfast: []string{"none"}, Lunch: []string{"none"}, Dinner: []string{"Pizza"}},
		}, {
			Name: "Burger Joe",
			Menu: menu{Breakfast: []string{"none"}, Lunch: []string{"none"}, Dinner: []string{"Burger"}},
		},
	}

	tpl.Execute(os.Stdout, restaurants)
}
