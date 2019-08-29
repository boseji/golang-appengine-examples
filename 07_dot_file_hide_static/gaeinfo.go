package main

import (
	"os"
)

// gaeInfo Stores the Appengine environment variables
type gaeInfo struct {
	Application    string // Application ID "GAE_APPLICATION"
	AppengineEnv   string // Appengine Env (standard / flex) "GAE_ENV"
	InstanceID     string // Instance "GAE_INSTANCE"
	MemoryMB       string // Memory in MegaBytes "GAE_MEMORY_MB"
	Runtime        string // Runtime "GAE_RUNTIME"
	ServiceName    string // Service Name "GAE_SERVICE"
	ServiceVersion string // Service Version "GAE_VERSION"
	ProjectID      string // GCP Project ID "GOOGLE_CLOUD_PROJECT"
	RunEnvironment string // Run Environment(only in NodeJS) "NODE_ENV"
	PORT           string // Server Port "PORT"
}

// getAppengineEnv Load and Reports the Appengine Environment
//   variables in the 'gaeInfo' data structure
func getAppengineEnv() *gaeInfo {
	gae := &gaeInfo{
		Application:    os.Getenv("GAE_APPLICATION"),
		AppengineEnv:   os.Getenv("GAE_ENV"),
		InstanceID:     os.Getenv("GAE_INSTANCE"),
		MemoryMB:       os.Getenv("GAE_MEMORY_MB"),
		Runtime:        os.Getenv("GAE_RUNTIME"),
		ServiceName:    os.Getenv("GAE_SERVICE"),
		ServiceVersion: os.Getenv("GAE_VERSION"),
		ProjectID:      os.Getenv("GOOGLE_CLOUD_PROJECT"),
		RunEnvironment: os.Getenv("NODE_ENV"),
		PORT:           os.Getenv("PORT"),
	}

	// For Development
	if gae.PORT == "" {
		gae.PORT = "8080"
		gae.RunEnvironment = "development"
	}
	return gae
}
