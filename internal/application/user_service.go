package application

import (
	"errors"
	"fmt"
	"my-go-app/internal/domain/user"
	"my-go-app/internal/mediator"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(name, email, password string) (*user.User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("name, email and password are required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		ID:           user.UserID(generateID()),
		Name:         name,
		Email:        email,
		PasswordHash: string(hashed),
		CreatedAt:    time.Now(),
	}

	err = s.Repo.Save(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func generateID() string {
	return fmt.Sprintf("user-%d", time.Now().UnixNano())
}

type RegisterUserCommand struct {
	Name     string
	Email    string
	Password string
}

type RegisterUserHandler struct {
	UserService *UserService
}

func NewRegisterUserHandler(us *UserService) *RegisterUserHandler {
	return &RegisterUserHandler{UserService: us}
}

func (h *RegisterUserHandler) Handle(req mediator.Request) (mediator.Response, error) {
	cmd, ok := req.(RegisterUserCommand)
	if !ok {
		return nil, fmt.Errorf("invalid request type")
	}
	return h.UserService.RegisterUser(cmd.Name, cmd.Email, cmd.Password)
}

type LoginCommand struct {
	Email    string
	Password string
}

type LoginHandler struct {
	Repo user.Repository
}

func NewLoginHandler(repo user.Repository) *LoginHandler {
	return &LoginHandler{Repo: repo}
}

func (h *LoginHandler) Handle(req mediator.Request) (mediator.Response, error) {
	cmd, ok := req.(LoginCommand)
	if !ok {
		return nil, errors.New("invalid request")
	}
	user, err := h.Repo.FindByEmail(cmd.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(cmd.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	token, err := generateJWT(string(user.ID))
	if err != nil {
		return nil, err
	}

	return map[string]string{"token": token}, nil
}

var jwtKey = []byte("supersecretkey")

func generateJWT(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
