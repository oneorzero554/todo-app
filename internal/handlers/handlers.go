package handlers

import "todo-app/internal/services"

type Handler struct {
    services *services.Services
}

func InitHandlers(services *services.Services) *Handler {
    return &Handler{services: services}
}