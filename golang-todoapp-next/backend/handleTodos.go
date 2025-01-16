package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/internal/database"
)

func (apiConfig *ApiConfig) handlePostTodos(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	params := parameter{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&params)

	if err != nil {
		log.Printf("Failed to decode it %v", err)
		http.Error(w, "Error to decode it", http.StatusBadRequest)
		return
	}

	todo, err := apiConfig.DB.CreateTodos(r.Context(), database.CreateTodosParams{
		Title:     params.Title,
		Content:   params.Content,
		Createdat: time.Now(),
	})

	if err != nil {
		log.Printf("Failed to create todo %v", err)
		http.Error(w, "Error to create todo", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, todo)
}
