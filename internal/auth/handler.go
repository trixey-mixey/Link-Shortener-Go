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
}

type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
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
		fmt.Println(body)
		res := RegisterResoponse{
			Token: "123",
		}
		response.Json(w, res, 200)
	}
}
