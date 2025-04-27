package handler

import (
	service "github.com/Quszlet/subscription-service/internal/services"
	"github.com/Quszlet/subscription-service/pkg/logger"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
	log logger.Logger
}

func NewHandler(service *service.Service, log logger.Logger) *Handler {
	return &Handler{services: service, log: log}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	users := r.PathPrefix("/users").Subrouter()
	users.HandleFunc("/create", h.Create).Methods("POST")
	users.HandleFunc("/{id:[0-9]+}", h.Get).Methods("GET")
	users.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("UPDATE")
	users.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")

	subscriptions := r.PathPrefix("/subscriptions").Subrouter()
	subscriptions.HandleFunc("/create", h.Create).Methods("POST")
	subscriptions.HandleFunc("/{id:[0-9]+}", h.Get).Methods("GET")
	subscriptions.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("UPDATE")
	subscriptions.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")

	subscriptionsMembers := r.PathPrefix("/subscriptions/members").Subrouter()
	subscriptionsMembers.HandleFunc("/{id:[0-9]+}", h.Get).Methods("GET")
	subscriptionsMembers.HandleFunc("/update/{id:[0-9]+}", h.Update).Methods("UPDATE")
	subscriptionsMembers.HandleFunc("/delete/{id:[0-9]+}", h.Delete).Methods("DELETE")

	return r
}
