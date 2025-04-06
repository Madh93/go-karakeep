package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Madh93/go-karakeep"
)

func main() {
	// Basic configuration
	apiUrl := "https://<YOUR_KARAKEEP_HOSTNAME>/api/v1" // Replace this with your API URL
	apiKey := "<YOUR_KARAKEEP_API_KEY>"                 // Replace this with your actual token

	// Set up Bearer authentication
	auth := func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		return nil
	}

	// Create the Karakeep client
	client, err := karakeep.NewClient(apiUrl, karakeep.WithRequestEditorFn(auth))
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	log.Printf("Hello world from %s", client.Server)
}
