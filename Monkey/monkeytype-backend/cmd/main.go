package main

import (
	"log"
	"net/http"

	"monkeytype-backend/internal/handlers"
)

func main() {
	typingHandler := handlers.NewTypingHandlerSource()

	http.HandleFunc("/typing", typingHandler.HandleTypingRequest)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
