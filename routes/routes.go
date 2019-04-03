/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| TokenAuthMiddleware middleware is used for X_API_KEY authentication.
| Enjoy building your API!
|
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"github.com/letsGo/controllers"
	"github.com/letsGo/helpers"
	"github.com/letsGo/middlewares"
)

func PaveRoutes() *gin.Engine {
	r := gin.Default()

	// websocket setup
	hub := controllers.NewHub()
	go hub.Run()

	r.Use(middlewares.TokenAuthMiddleware())

	config := ginserver.Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			helpers.RespondWithError(ctx, 401, "invalid access_token")
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}
	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Home)
		v1.GET("/credentials", controllers.GetCredentials)
		v1.GET("/token", controllers.GetToken)
		auth := v1.Group("auth")
		{
			auth.Use(ginserver.HandleTokenVerify(config))
			auth.GET("/", controllers.Verify)
		}

		// websocket route
		r.GET("/ws", func(c *gin.Context) {
			controllers.ServeWebsocket(hub, c.Writer, c.Request)
		})

	}

	return r
}
