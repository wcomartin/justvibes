package http

import (
	"encoding/json"
	"my-go-app/internal/application"
	"my-go-app/internal/mediator"
	"net/http"
)

type AuthHandler struct {
	mediator mediator.Mediator
}

func NewAuthHandler(m mediator.Mediator) *AuthHandler {
	return &AuthHandler{mediator: m}
}

func RegisterAuthRoutes(mux *http.ServeMux, authHandler *AuthHandler) {
	mux.HandleFunc("/users", authHandler.HandleRegisterUser)
	mux.HandleFunc("/login", authHandler.HandleLogin)
}

func (h *AuthHandler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
	var cmd application.RegisterUserCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.mediator.Send(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *AuthHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var cmd application.LoginCommand
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := h.mediator.Send(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
