package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Working Golang App")
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
