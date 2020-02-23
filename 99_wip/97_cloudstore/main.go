package main

import (
	"fmt"
	"log"
	"net/http"
	// "cloud.google.com/go/storage"
)

var gae *gaeInfo

// var client *storage.Client

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI != "/" {
		http.Error(w, "In correct request", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "Working Golang App\n")

	// Create Client
	// client, err := storage.NewClient(r.Context())
	// if err != nil {
	// 	log.Printf("Error Failed to Create Client: %v", err)
	// 	fmt.Fprintf(w, "Error in Creating Client\n")
	// 	return
	// }
	// defer client.Close()
	// fmt.Fprintf(w, " 1. Client Creation: Success\n")
	fmt.Fprintf(w, " 2. Application ID : %s\n", gae.Application)
	fmt.Fprintf(w, " 3. Environment : %s\n", gae.AppengineEnv)
	fmt.Fprintf(w, " 4. Instance : %s\n", gae.InstanceID)
	fmt.Fprintf(w, " 5. Memory : %s MegaBytes\n", gae.MemoryMB)
	fmt.Fprintf(w, " 6. Runtime : %s\n", gae.Runtime)
	fmt.Fprintf(w, " 7. Service Name : %s\n", gae.ServiceName)
	fmt.Fprintf(w, " 8. Service Version : %s\n", gae.ServiceVersion)
	fmt.Fprintf(w, " 9. GCP Project ID : %s\n", gae.ProjectID)
	fmt.Fprintf(w, "10. Run Environment(only in NodeJS) : %s\n", gae.RunEnvironment)
	fmt.Fprintf(w, "11. Server Port : %s\n", gae.PORT)

	// Get Default Bucket Name
	// it := client.Buckets(r.Context(), projectID)
	// for {
	// 	bucketAttrs, err := it.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		// TODO: Handle error.
	// 		fmt.Fprint(w, "Found Error in Iteration\n")
	// 	} else {
	// 		fmt.Fprintln(w, bucketAttrs.Name)
	// 	}
	// }
}

func main() {

	// Check Client Creation
	// var err error
	// if client, err = storage.NewClient(context.Background()); err != nil {
	// 	log.Fatalf("Error Creating Client %v", err)
	// }

	http.HandleFunc("/", indexHandler)

	// Load the Environment Info from Appengine
	gae = getAppengineEnv()

	log.Printf("Starting server on :%s", gae.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", gae.PORT), nil))
}
