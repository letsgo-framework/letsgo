package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/letsGo/helpers"
	"log"
	"os"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("X_API_KEY")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set X_API_KEY environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get(`X_API_KEY`)

		if token == "" {
			helpers.RespondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			helpers.RespondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}
