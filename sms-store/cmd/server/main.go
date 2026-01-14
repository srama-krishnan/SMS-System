package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"sms-store/internal/httpapi"
	"sms-store/internal/store"
)

func main() {
	// MongoDB connection configuration
	connectionString := getEnv("MONGODB_URI", "mongodb://localhost:27017")
	databaseName := getEnv("MONGODB_DATABASE", "sms_store")
	collectionName := getEnv("MONGODB_COLLECTION", "messages")

	// Initialize MongoDB store
	log.Println("Connecting to MongoDB...")
	mongoStore, err := store.NewMongoStore(connectionString, databaseName, collectionName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	log.Println("Successfully connected to MongoDB")

	// Ensure MongoDB connection is closed on shutdown
	defer func() {
		log.Println("Closing MongoDB connection...")
		if err := mongoStore.Close(); err != nil {
			log.Printf("Error closing MongoDB connection: %v", err)
		}
	}()

	// Create handler with MongoDB store
	h := httpapi.NewHandler(mongoStore)

	mux := http.NewServeMux()

	// GET /ping - Health check endpoint
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.Ping(w, r)
	})

	// GET /v1/user/{user_id}/messages - Required endpoint for SMS Store
	mux.HandleFunc("/v1/user/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		h.GetUserMessages(w, r)
	})

	// POST /messages - Optional endpoint for testing (can be removed later)
	mux.HandleFunc("/messages", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateMessage(w, r)
		case http.MethodGet:
			h.ListMessages(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := ":8082"
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	// Setup graceful shutdown
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		log.Println("Shutting down server...")
		if err := server.Close(); err != nil {
			log.Printf("Error closing server: %v", err)
		}
	}()

	log.Println("sms-store server started at", addr)
	log.Println("Available endpoints:")
	log.Println("  GET  /ping")
	log.Println("  GET  /v1/user/{user_id}/messages")
	log.Println("  POST /messages (testing only)")

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
