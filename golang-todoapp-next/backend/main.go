package main

import (
	"database/sql"
	"log"
	"net/http"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	connStr := "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

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

	http.ListenAndServe(":8000", r)
}
