package handler

import (
	service "github.com/Quszlet/subscription-service/internal/services"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	// services *service.Service
}

func NewHandler(service *service.Service, log logger.Logger) *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// users := r.PathPrefix("/users").Subrouter()
	// users.HandleFunc("/create", h.Create).Methods("POST")
	// users.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("UPDATE")
	// users.HandleFunc("/{id:[0-9]+}", h.Get).Methods("GET")
	// users.HandleFunc("", h.GetAll).Methods("GET")
	// users.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")
	return r
}
