package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Post("/todos", apiConfig.handlePostTodos)
	r.Delete("/todos/{id}", apiConfig.handleDeleteTodos)
	r.Patch("/todos/{id}", apiConfig.handleEditTodos)
	r.Get("/todos", apiConfig.handleAllTodos)
	r.Get("/todos/{id}", apiConfig.handleOneTodos)

	http.ListenAndServe(":8000", r)
}
