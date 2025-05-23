package main

import (
	"log"
	stdhttp "net/http"

	"my-go-app/internal/application"
	"my-go-app/internal/infrastructure/persistence"
	apphttp "my-go-app/internal/interfaces/http"
	"my-go-app/internal/mediator"
)

func main() {
	sqliteRepo, err := persistence.NewSqliteUserRepo("users.db")
	if err != nil {
		log.Fatalf("failed to open sqlite db: %v", err)
	}
	userService := application.NewUserService(sqliteRepo)
	registerHandler := application.NewRegisterUserHandler(userService)
	loginHandler := application.NewLoginHandler(sqliteRepo)

	med := mediator.NewSimpleMediator()
	med.Register("application.RegisterUserCommand", registerHandler)
	med.Register("application.LoginCommand", loginHandler)

	h := apphttp.NewHandlerWithMediator(med)
	authHandler := apphttp.NewAuthHandler(med)

	mux := stdhttp.NewServeMux()

	authMux := stdhttp.NewServeMux()
	apphttp.RegisterAuthRoutes(authMux, authHandler)
	mux.Handle("/api/auth/", stdhttp.StripPrefix("/api/auth", authMux))

	mux.HandleFunc("/api/dashboard", apphttp.VerifyJWT(h.HandleDashboard))

	// Serve static files from the Vue build directory
	fs := stdhttp.FileServer(stdhttp.Dir("frontend/dist"))
	mux.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(stdhttp.ListenAndServe(":8080", mux))
}
