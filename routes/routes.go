package routes

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/letsgo/controllers"
	"gitlab.com/letsgo/middlewares"
)

func PaveRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.TokenAuthMiddleware())
	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.HelloWorld)

	}

	return r
}
