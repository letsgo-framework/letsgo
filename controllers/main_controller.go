package controllers

import (
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	c.String(200, `Hello World`)
	c.Done()
}