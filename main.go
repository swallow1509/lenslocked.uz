package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to my Awesome Gallery!</h1>")
}

func contactsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to Contacts page</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We are sorry! No such page found :(</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to FAQ page. You can ask your questions here</h1>")
}

func main() {

	var router = mux.NewRouter()

	//handler for NotFound request
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	//URL paths and corresponding handlers
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/contacts", contactsHandler)
	router.HandleFunc("/faq", faqHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}
