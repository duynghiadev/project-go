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

func (apiConfig *ApiConfig) handleOneTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")
	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)

	if err != nil {
		log.Printf("Failed to get all todos %v", err)
		http.Error(w, "Error to get all todos", http.StatusBadRequest)
		return
	}

	todo, err := apiConfig.DB.GetTodos(r.Context(), todoIntId)

	if err != nil {
		log.Printf("Failed to Get one todo %v", err)
		http.Error(w, "Error to Getting one todos", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, todo)
}

func (apiConfig *ApiConfig) handleAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := apiConfig.DB.ListAllTodos(r.Context())

	if err != nil {
		log.Printf("Failed to create todo %v", err)
		http.Error(w, "Error to create todo", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, todos)
}

func (apiConfig *ApiConfig) handleEditTodos(w http.ResponseWriter, r *http.Request) {
	todoIdString := chi.URLParam(r, "id")
	todoIntId, err := strconv.ParseInt(todoIdString, 10, 64)

	if err != nil {
		log.Printf("Failed to parse Id %v", err)
		http.Error(w, "Error to parse Id", http.StatusBadRequest)
		return
	}

	type parameter struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	params := parameter{}

	decode := json.NewDecoder(r.Body)
	err = decode.Decode(&params)

	if err != nil {
		log.Printf("Failed to decode it %v", err)
		http.Error(w, "Error to decode it", http.StatusBadRequest)
		return
	}

	err = apiConfig.DB.UpdateTodos(r.Context(), database.UpdateTodosParams{
		ID:      todoIntId,
		Title:   params.Title,
		Content: params.Content,
	})

	if err != nil {
		log.Printf("Failed to edit todo %v", err)
		http.Error(w, "Error to edit todo", http.StatusBadRequest)
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"message": "successfully Edit Todo"})
}

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
