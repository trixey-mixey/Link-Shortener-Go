package main

import (
	"fmt"
	"go/projcet-Adv/configs"
	"go/projcet-Adv/internal/auth"
	"go/projcet-Adv/internal/link"
	"go/projcet-Adv/internal/stat"
	"go/projcet-Adv/internal/users"
	"go/projcet-Adv/pkg/db"
	"go/projcet-Adv/pkg/event"
	"go/projcet-Adv/pkg/middleware"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositoires
	linkRepository := link.NewLinkRepository(db)
	userRepository := users.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})
	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
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

	go statService.AddClick()

	fmt.Println("Server is listening on port 8081")
	server.ListenAndServe()
}
