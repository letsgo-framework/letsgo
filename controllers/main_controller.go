package controllers

import (
	"github.com/gin-gonic/gin"
)


// The content below is only a placeholder and can be replaced.
type Welcome struct {
	Greet string `json:"greet"`
	Doc string `json:"link_to_doc"`
	Github string `json:"github"`
	Examples string `json:"examples"`
}
func Home(c *gin.Context) {
	var welcome Welcome
	welcome.Greet = `Welcome to letsGo`
	welcome.Doc = `Link to readme`
	welcome.Github = `Link to github`
	welcome.Examples = `Link To examples`
	c.JSON(200, welcome)
	c.Done()
}