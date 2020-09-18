package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template
var contactsTemplate *template.Template
var faqTemplate *template.Template

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactsTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := faqTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We are sorry! No such page found :(</h1>")
}

func main() {

	var router = mux.NewRouter()
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if homeTemplate, err = template.ParseFiles("views/home.gohtml"); err != nil {
		panic(err)
	}
	if contactsTemplate, err = template.ParseFiles("views/contacts.gohtml"); err != nil {
		panic(err)
	}
	if faqTemplate, err = template.ParseFiles("views/faq.gohtml"); err != nil {
		panic(err)
	}

	//handler for NotFound request
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//URL paths and corresponding handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/contacts", contactsHandler)
	router.HandleFunc("/faq", faqHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
