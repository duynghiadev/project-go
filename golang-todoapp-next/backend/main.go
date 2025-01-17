package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	connStr := os.Getenv("POSTGRES_URL")

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
	r.Get("/todos", apiConfig.handleAllTodos)
	r.Get("/todos/{id}", apiConfig.handleOneTodos)

	http.ListenAndServe(":8000", r)
}
