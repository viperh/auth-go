package app

import (
	"auth-go/internal/api/controllers"
	"auth-go/internal/api/routes"
	"auth-go/internal/config"
	"auth-go/internal/provider"
	"auth-go/internal/statistics"
	"github.com/gin-gonic/gin"
	"time"
)

type App struct {
	Controller *controllers.Controller
	Config     *config.Config
	Provider   *provider.Postgres
	Engine     *gin.Engine
}

func New() *App {
	cfg := config.New()
	prv := provider.New(cfg)
	stats := statistics.New()
	stats.SetStartTime(uint64(time.Now().Unix()))
	cnt := controllers.New(prv, stats, cfg)

	engine := gin.Default()
	routes.DefineRoutes(engine, cnt)
	return &App{
		Controller: cnt,
		Config:     cfg,
		Provider:   prv,
		Engine:     engine,
	}
}

func (a *App) Run() error {
	return a.Engine.Run(":" + a.Config.Port)
}
