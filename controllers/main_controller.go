package controllers

import (
	"github.com/gin-gonic/gin"
)


// The content below is only a placeholder and can be replaced.
func Home(c *gin.Context) {
	c.String(200, `Welcome to letsGo`)
	c.Done()
}