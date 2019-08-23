package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Get the Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to Port %s", port)
		// May Be we are Running in Development
		//  So Lets serve the Assets Also
		http.Handle("/assets/",
			http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}

	log.Printf("Listening on Port %s : Server", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
