package routes

import (
	"auth-go/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

// format:  https://localhost/api/v1/

func DefineRoutes(g *gin.Engine, c *controllers.Controller) {
	v1 := DefineMainRoutes(g)
	DefineAuthRoutes(v1, c)
	DefinePublicRoutes(v1, c)
}

func DefineMainRoutes(g *gin.Engine) *gin.RouterGroup {
	api := g.Group("/api")
	v1 := api.Group("/v1")
	return v1
}

func DefineAuthRoutes(auth *gin.RouterGroup, c *controllers.Controller) {

	auth.POST("/login", c.Login)
	auth.POST("/register", c.Register)
}

func DefinePublicRoutes(public *gin.RouterGroup, c *controllers.Controller) {
	public.GET("/health", c.GetInfo)
}
