package app

import (
	"fmt"
	"log"

	"bcc-workshop-2/config"
	pkggin "bcc-workshop-2/pkg/gin"
	"bcc-workshop-2/pkg/postgres"

	ginfwk "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	config *config.Config
	db     *gorm.DB
	router *ginfwk.Engine
}

func New(cfg *config.Config) *App {
	db := postgres.New(cfg)
	router := pkggin.New(cfg)

	return &App{
		config: cfg,
		db:     db,
		router: router,
	}
}

func (a *App) Run() {
	addr := fmt.Sprintf(":%s", a.config.App.Port)
	log.Printf("Server %s running on port %s", a.config.App.Name, a.config.App.Port)

	if err := a.router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
