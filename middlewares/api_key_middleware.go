package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("X_API_KEY")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set X_API_KEY environment variable")
	}

	return func(c *gin.Context) {
		token := c.Request.Header.Get(`X_API_KEY`)

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}
