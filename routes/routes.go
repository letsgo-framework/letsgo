/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application.
| Enjoy building your API!
|
*/

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"github.com/letsgo-framework/letsgo/controllers"
	"github.com/letsgo-framework/letsgo/gql"
	"github.com/letsgo-framework/letsgo/helpers"
	"github.com/letsgo-framework/letsgo/jobs"
)

// PaveRoutes sets up all api routes
func PaveRoutes() *gin.Engine {
	r := gin.Default()

	// websocket setup
	hub := controllers.NewHub()
	go hub.Run()

	// CORS
	r.Use(cors.Default())

	// CRON
	jobs.Run()

	// Auth Init
	controllers.AuthInit()
	config := ginserver.Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			helpers.RespondWithError(ctx, 401, "invalid access_token")
		},
		TokenKey: "github.com/go-oauth2/gin-server/access-token",
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}

	// Graphql Init
	gql.InitGraphql(r)

	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Home)
		v1.GET("/credentials", controllers.GetCredentials)
		v1.GET("/login", controllers.GetToken)
		v1.POST("/register", controllers.Register)
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
