package handler

import (
	"todo-service/pkg/usecase/task/service"
)

type Handler struct {
	Service service.TodoService
}

func NewHandler(service service.TodoService) *Handler {
	return &Handler{Service: service}
}
