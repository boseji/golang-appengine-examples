package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Working Golang App")
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Get the Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to Port %s", port)
	}

	log.Printf("Listening on Port %s : Server", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
