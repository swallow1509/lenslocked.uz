package main

import (
	"html/template"
	"os"
)

//User is exported to hello.gohtml
type User struct {
	Name string
	Dog  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := User{
		Name: "Umidbek",
		Dog:  "Morty",
	}
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	data.Name = "Sardorbek"
	data.Dog = "Goldie"
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
