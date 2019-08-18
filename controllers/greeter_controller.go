package controllers

import (
	"github.com/gin-gonic/gin"
)

// Welcome !! The content below is only a placeholder and can be replaced.
type Welcome struct {
	Greet    string `json:"greet"`
	Doc      string `json:"link_to_doc"`
	Github   string `json:"github"`
	Examples string `json:"examples"`
}

// Greet is the response for api/v1
func Greet(c *gin.Context) {

	welcome := Welcome{
		Greet:    "Welcome to letsGo",
		Doc:      "https://letsgo-framework.github.io/",
		Github:   "https://github.com/letsgo-framework/letsgo",
		Examples: "Coming Soon",
	}
	c.JSON(200, welcome)
	c.Done()
}
