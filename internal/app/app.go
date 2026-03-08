package app

import (
	"fmt"
	"log"

	"bcc-workshop-2/config"
	"bcc-workshop-2/internal/controller/rest"
	"bcc-workshop-2/internal/entity"
	"bcc-workshop-2/internal/repository"
	"bcc-workshop-2/internal/usecase"
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

	app := &App{
		config: cfg,
		db:     db,
		router: router,
	}

	app.migrate()
	app.setupRoutes()

	return app
}

func (a *App) migrate() {
	if err := a.db.AutoMigrate(
		&entity.Restaurant{},
		&entity.Item{},
	); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("Database migration completed")
}

func (a *App) setupRoutes() {
	repo := repository.NewRepository(a.db)
	uc := usecase.NewUsecase(repo)

	rest.NewRouter(a.router, uc)
}

func (a *App) Run() {
	addr := fmt.Sprintf(":%s", a.config.App.Port)
	log.Printf("Server %s running on port %s", a.config.App.Name, a.config.App.Port)

	if err := a.router.Run(addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
