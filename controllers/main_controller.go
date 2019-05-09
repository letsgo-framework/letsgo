package controllers

import (
	"github.com/gin-gonic/gin"
)

// Welcome!! The content below is only a placeholder and can be replaced.
type Welcome struct {
	Greet    string `json:"greet"`
	Doc      string `json:"link_to_doc"`
	Github   string `json:"github"`
	Examples string `json:"examples"`
}

// Response for api/v1
func Home(c *gin.Context) {
	var welcome Welcome
	welcome.Greet = `Welcome to letsGo`
	welcome.Doc = `https://letsgo-framework.github.io/`
	welcome.Github = `https://github.com/letsgo-framework/letsgo`
	welcome.Examples = `Link To examples`
	c.JSON(200, welcome)
	c.Done()
}
