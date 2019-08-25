package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var gae *gaeInfo
var myVARenv string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.Error(w, "Incorrect request", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Working Golang App\n")
	fmt.Fprintf(w, " 1. Application ID : %s\n", gae.Application)
	fmt.Fprintf(w, " 2. Environment : %s\n", gae.AppengineEnv)
	fmt.Fprintf(w, " 3. Instance : %s\n", gae.InstanceID)
	fmt.Fprintf(w, " 4. Memory : %s MegaBytes\n", gae.MemoryMB)
	fmt.Fprintf(w, " 5. Runtime : %s\n", gae.Runtime)
	fmt.Fprintf(w, " 6. Service Name : %s\n", gae.ServiceName)
	fmt.Fprintf(w, " 7. Service Version : %s\n", gae.ServiceVersion)
	fmt.Fprintf(w, " 8. GCP Project ID : %s\n", gae.ProjectID)
	fmt.Fprintf(w, " 9. Run Environment(only in NodeJS) : %s\n", gae.RunEnvironment)
	fmt.Fprintf(w, "10. Server Port : %s\n", gae.PORT)
	fmt.Fprintf(w, "11. Our additional Environment Var: %s", myVARenv)
}

func main() {

	http.HandleFunc("/", indexHandler)

	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	// Read our Custom environment Variable
	myVARenv = os.Getenv("MY_VAR")
	if myVARenv == "" {
		log.Println("Environment Variable 'MY_VAR' not set")
	}

	log.Printf("Starting server on :%s", gae.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), nil))
}
