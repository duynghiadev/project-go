package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
	"todo/internal/database"

	"github.com/go-chi/chi/v5"
)

func (apiConfig *ApiConfig) handleDeleteTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")
	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)

	if err != nil {
		log.Printf("Failed to parse Id %v", err)
		http.Error(w, "Error to parse Id", http.StatusBadRequest)
		return
	}

	err = apiConfig.DB.DeleteTodos(r.Context(), todoIntId)

	if err != nil {
		log.Printf("Failed to create todo %v", err)
		http.Error(w, "Error to create todo", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "deleted well!"})
}

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
