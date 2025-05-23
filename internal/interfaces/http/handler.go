package http

import (
    "encoding/json"
    "fmt"
    "my-go-app/internal/application"
    "my-go-app/internal/mediator"
    "net/http"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

type Handler struct {
    mediator mediator.Mediator
}

func NewHandlerWithMediator(m mediator.Mediator) *Handler {
    return &Handler{mediator: m}
}

func (h *Handler) HandleRegisterUser(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) HandleHome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to the home page!",
    })
}

func (h *Handler) HandleDashboard(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to your dashboard!",
    })
}

var jwtKey = []byte("supersecretkey")

func VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}