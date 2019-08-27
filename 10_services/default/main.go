package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h2>Main Route in Golang App</h2><br>")
	if r.TLS == nil {
		fmt.Fprintf(w, "<a href=\"http://uuid.%v\">Link to Service</a>", r.Host)
	} else {
		fmt.Fprintf(w, "<a href=\"https://uuid-dot-%v\">Link to Service</a>", r.Host)
	}
}
