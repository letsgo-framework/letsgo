package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"gitlab.com/letsgo/controllers"
	"gitlab.com/letsgo/middlewares"
)

func PaveRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.TokenAuthMiddleware())

	config := ginserver.Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			respondWithError(ctx, 401, "invalid access_token")
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}
	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.HelloWorld)
		v1.GET("/credentials", controllers.GetCredentials)
		v1.GET("/token", controllers.GetToken)
		auth := v1.Group("auth")
		{
			auth.Use(ginserver.HandleTokenVerify(config))
			auth.GET("/", controllers.Verify)
		}

	}

	return r
}

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
