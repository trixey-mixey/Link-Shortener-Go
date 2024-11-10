package main

import (
	"fmt"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/link"
	"go/projcet-Adv/pkg/db"
	"go/projcet-Adv/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositoires
	linkRepository := link.NewLinkRepository(db)
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
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
