package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to read godotenv")
	}

	connStr := os.Getenv("POSTGRES_URL")

	log.Println(connStr)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	db := database.New(conn)

	apiConfig := ApiConfig{
		DB: db,
	}

	r := chi.NewRouter()

	r.Post("/todos", apiConfig.handlePostTodos)
	r.Delete("/todos/{id}", apiConfig.handleDeleteTodos)
	r.Patch("/todos/{id}", apiConfig.handleEditTodos)

	http.ListenAndServe(":8000", r)
}
