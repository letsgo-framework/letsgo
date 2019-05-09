package helpers

import "github.com/gin-gonic/gin"

// RespondWithError creates response for error
func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
