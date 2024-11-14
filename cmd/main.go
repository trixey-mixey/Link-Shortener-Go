package main

import (
	"context"
	"fmt"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/link"
	"go/projcet-Adv/internal/users"
	"go/projcet-Adv/pkg/db"
	"go/projcet-Adv/pkg/middleware"
	"net/http"
	"time"
)

func tickOperation(ctx context.Context) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-ctx.Done():
			fmt.Println("Cancel")
			return
		}
	}
}

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositoires
	linkRepository := link.NewLinkRepository(db)
	userRepository := users.NewUserRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
	})

	// Middlewares
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
