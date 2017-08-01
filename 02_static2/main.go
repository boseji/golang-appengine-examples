package main

import (
	"html/template"
	"net/http"

	"log"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
