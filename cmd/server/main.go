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

	stdhttp.HandleFunc("/api/users", h.HandleRegisterUser)
	stdhttp.HandleFunc("/api/login", h.HandleLogin)
	stdhttp.HandleFunc("/api/dashboard", apphttp.VerifyJWT(h.HandleDashboard))

	// Serve static files from the Vue build directory
	fs := stdhttp.FileServer(stdhttp.Dir("frontend/dist"))
	stdhttp.Handle("/", fs)

	log.Println("Server running on :8080")
	log.Fatal(stdhttp.ListenAndServe(":8080", nil))
}
