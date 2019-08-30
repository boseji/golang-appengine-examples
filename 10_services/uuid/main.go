package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := uuid.NewV4()
	w.Header().Set("Content-type", "text/html")
	fmt.Fprintf(w, "<h2>%v</h2><br>", u.String())
	if strings.Contains(r.Host, ":8081") {
		// We are in the Development Mode
		fmt.Fprintf(w, "<a href=\"http://%v\">Link to Service</a>",
			strings.Replace(r.Host, "8081", "8080", -1))
		return
	}
	// if strings.Contains(r.URL.RawPath, "https") {
	fmt.Fprintf(w, "<a href=\"https://%v\">Link to Main</a>",
		strings.Replace(r.Host, "uuid-dot-", "", 1))
	// } else {
	// 	fmt.Fprintf(w, "<a href=\"http://%v\">Link to Main</a>",
	// 		strings.Replace(r.Host, "uuid.", "", 1))
	// }
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// Get the Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("Defaulting to Port %s", port)
	}
	log.Printf("Listening on Port %s : Server", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
