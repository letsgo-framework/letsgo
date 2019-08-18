package routes

import (
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/letsgo-framework/letsgo/controllers"
	"github.com/letsgo-framework/letsgo/helpers"
)

func AuthRoutes(r *gin.RouterGroup) *gin.RouterGroup {

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

	r.GET("/credentials", controllers.GetCredentials)
	r.GET("/login", controllers.GetToken)
	r.POST("/register", controllers.Register)
	auth := r.Group("auth")
	{
		auth.Use(ginserver.HandleTokenVerify(config))
	}

	return auth
}
