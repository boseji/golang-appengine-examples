package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
)

// Firebase App Instance
var app *firebase.App

func init() {
	var err error
	// Initialize the Global instance
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// Note: This would be done only Once
	log.Println("Init Firebase Completed...")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Working Golang App")
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	// Get the Port
	port := os.Getenv("PORT")
	isDev := false // Detect if we are running in Development Machine
	if port == "" {
		port = "8080"
		isDev = true
		log.Printf("Defaulting to Port %s", port)
	}

	// This it check if we are able to get the Auth Object
	if _, err := app.Auth(context.Background()); err != nil {
		log.Fatalf("Could not get Firebase Auth object")
	}

	log.Printf("Listening on Port %s : Server", port)
	log.Printf("Running in Development Environment : %v", isDev)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
