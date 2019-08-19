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
	"github.com/letsgo-framework/letsgo/jobs"
)

// PaveRoutes sets up all api routes
func PaveRoutes() *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(cors.Default())

	// CRON
	jobs.Run()

	// Grouped api
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", controllers.Greet)
		auth := AuthRoutes(v1)
		auth.GET("/", controllers.Verify)
	}

	return r
}
