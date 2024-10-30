package main

import (
	"go/projcet-Adv/configs"
	"go/projcet-Adv/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	server.ListenAndServe()
}
