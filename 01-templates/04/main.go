/*

Using the data structure created in the previous folder, modify it to hold menu information for an unlimited number of restaurants.

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

type restaurants []restaurant

var tpl template.Template

func init() {
	tpl = *template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	restaurantsList := restaurants{
		{
			Name: "Pizza Bob",
			Menu: menu{Breakfast: []string{"none"}, Lunch: []string{"none"}, Dinner: []string{"Pizza"}},
		}, {
			Name: "Burger Joe",
			Menu: menu{Breakfast: []string{"none"}, Lunch: []string{"none"}, Dinner: []string{"Burger"}},
		},
	}

	tpl.Execute(os.Stdout, restaurantsList)
}
