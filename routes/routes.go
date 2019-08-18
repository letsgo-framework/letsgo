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
	"github.com/letsgo-framework/letsgo/controllers"
	"github.com/letsgo-framework/letsgo/gql"
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

	// Graphql Init
	gql.InitGraphql(r)

	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Greet)
		auth := AuthRoutes(v1)
		auth.GET("/", controllers.Verify)
		// websocket route
		r.GET("/ws", func(c *gin.Context) {
			controllers.ServeWebsocket(hub, c.Writer, c.Request)
		})

	}

	return r
}
