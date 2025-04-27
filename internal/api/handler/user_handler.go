package handler

import (
	// "fmt"
	"net/http"
	// "strconv"
	// "github.com/Quszlet/CRUD-operations/internal/models"
	// "github.com/Quszlet/CRUD-operations/pkg/JSON"
	// "github.com/gorilla/mux"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.log.Debug("Запрос пришел")
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	//id, _ := strconv.Atoi(mux.Vars(r)["id"])
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	
}