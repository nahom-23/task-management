package handler

import (
	"encoding/json"
	// "log"
	"net/http"
	"task-management/internal/model"
	"task-management/internal/service"
	// "strconv"

	// "github.com/gorilla/mux"
)

type UserHandler struct {
	Service *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{Service: svc}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u model.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateUser(r.Context(), &u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(u)
}