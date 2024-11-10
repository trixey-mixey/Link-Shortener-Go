package main

import (
	"fmt"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/link"
	"go/projcet-Adv/pkg/db"
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

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
