package routes

import (
	"auth-go/internal/api/controllers"
	"auth-go/internal/api/middlewares"
	"github.com/gin-gonic/gin"
)

// format:  https://localhost/api/v1/

func DefineRoutes(g *gin.Engine, c *controllers.Controller) {
	api := g.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/health", c.GetInfo)
	v1.POST("/login", c.Login)
	v1.POST("/register", c.Register)

	auth := v1.Group("/auth")
	auth.Use(middlewares.JwtMiddleware())
	auth.GET("/posts", provider.GetPosts)
}
