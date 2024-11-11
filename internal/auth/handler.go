package auth

import (
	"fmt"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/pkg/request"
	"go/projcet-Adv/pkg/response"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Прочитать body
		body, err := request.HandleBody[LoginRequest](&w, req)
		if err != nil {
			return
		}

		fmt.Println(body)
		res := LoginResponse{
			Token: "123",
		}
		response.Json(w, res, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](&w, req)
		if err != nil {
			return
		}
		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}
