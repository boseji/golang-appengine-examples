package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	fmt.Fprint(w, "<h2>Main Route in Golang App</h2><br>")
	if strings.Contains(r.Host, ":8080") {
		// We are in the Development Mode
		fmt.Fprintf(w, "<a href=\"http://%v\">Link to Service</a>",
			strings.Replace(r.Host, "8080", "8081", -1))
		return
	}
	// Check in case We have HTTPS - To Allways HTTPS
	// if r.TLS != nil {
	fmt.Fprintf(w, "<a href=\"https://uuid-dot-%v\">Link to Service</a>", r.Host)
	// } else {
	// fmt.Fprintf(w, "<a href=\"http://uuid.%v\">Link to Service</a>", r.Host)
	// }
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
