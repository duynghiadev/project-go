package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"todo-service/pkg/handler"
	"todo-service/pkg/usecase/task/repository"
	"todo-service/pkg/usecase/task/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get credentials from environment variables
	mongoUsername := os.Getenv("MONGO_USERNAME")
	mongoPassword := os.Getenv("MONGO_PASSWORD")
	mongoURI := os.Getenv("MONGO_URI")
	mongoDB := os.Getenv("MONGO_DB")

	if mongoUsername == "" || mongoPassword == "" || mongoURI == "" || mongoDB == "" {
		log.Fatal("Missing required MongoDB credentials in environment variables")
	}

	// Construct the MongoDB URI
	mongoConnectionURI := "mongodb+srv://" + mongoUsername + ":" + mongoPassword + "@" + mongoURI + "/" + mongoDB + "?retryWrites=true&w=majority"

	// Create a background context for the MongoDB client
	ctx := context.Background()

	// Set MongoDB client options with longer timeouts
	clientOptions := options.Client().
		ApplyURI(mongoConnectionURI).
		SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)).
		SetTimeout(30 * time.Second).
		SetSocketTimeout(30 * time.Second).
		SetTLSConfig(&tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: false,
		})

	// Connect to MongoDB with retry logic
	var client *mongo.Client
	var err error

	for i := 0; i < 3; i++ {
		connectCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		client, err = mongo.Connect(connectCtx, clientOptions)
		cancel()

		if err != nil {
			log.Printf("Failed to connect to MongoDB, attempt %d: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		// Try to ping the database
		pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		err = client.Ping(pingCtx, nil)
		cancel()

		if err != nil {
			log.Printf("Failed to ping MongoDB, attempt %d: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB after retries: %v", err)
	}

	log.Println("Connected to MongoDB!")

	// Initialize repository, service, and handler
	todoRepo := repository.NewMongoRepository(client)
	todoService := service.NewService(todoRepo)
	todoHandler := handler.NewHandler(todoService)

	// Set up routes
	http.HandleFunc("/api/v1/add", todoHandler.AddTask)
	http.HandleFunc("/api/v1/get-all", todoHandler.GetAllTasks)
	http.HandleFunc("/api/v1/update", todoHandler.UpdateTask)
	http.HandleFunc("/api/v1/complete", todoHandler.CompleteTask)

	// Create a server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      nil,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	// Channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start the server in a goroutine
	go func() {
		log.Println("Starting server on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Wait for an OS signal to shutdown
	<-stop
	log.Println("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server Shutdown Failed: %+v", err)
	}

	// Disconnect from MongoDB
	if err := client.Disconnect(shutdownCtx); err != nil {
		log.Printf("MongoDB Disconnect Failed: %+v", err)
	}

	log.Println("Server exited properly")
}
