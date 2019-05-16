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

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/letsgo-framework/letsgo/docs"
)

// PaveRoutes sets up all api routes
func PaveRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	// websocket setup
	hub := controllers.NewHub()
	go hub.Run()

	// CORS
	r.Use(cors.Default())

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
