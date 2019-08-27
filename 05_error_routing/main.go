package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var gae *gaeInfo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		// Display the 404 file
		content, _ := ioutil.ReadFile("404.html")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", content)
		return
	}
	fmt.Fprint(w, "Working Golang App\n")
}

func main() {

	http.HandleFunc("/", indexHandler)

	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	log.Printf("Starting server on :%s", gae.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), nil))
}
