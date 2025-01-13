package main

import (
	"context"
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
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(mongoConnectionURI)

	client, err := mongo.Connect(ctx, clientOptions)
	// cp 0.006 (0.01) -> 1CPU x6 -> 6CPU
	// RM 15Mb -> 128Mb
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
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
		Addr:    ":8080",
		Handler: nil,
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
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	// Disconnect from MongoDB
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("MongoDB Disconnect Failed:%+v", err)
	}

	log.Println("Server exited properly")
}
