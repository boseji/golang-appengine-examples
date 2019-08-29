package main

import (
	"fmt"
	"log"
	"net/http"
)

var gae *gaeInfo

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		return404(w, r)
		return
	}
	fmt.Fprint(w, "Working Golang App\n")
}

func main() {

	http.HandleFunc("/", indexHandler)
	// Create the File System
	fs := dotFileHidingFileSystem{http.Dir("public")}
	// Create the File Server
	// hfileServer := http.FileServer(http.Dir("public"))
	hfileServer := http.FileServer(fs)
	// Create the 404 Page Handler
	h404 := &handlerFor404{}
	// Finally the Wrapping of File Server and 404 Handler
	hfileServerWith404 := static404Handler(hfileServer, h404)
	http.Handle("/home/", http.StripPrefix("/home/", hfileServerWith404))
	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	log.Printf("Starting server on :%s", gae.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), nil))
}
