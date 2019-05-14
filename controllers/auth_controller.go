/*
|--------------------------------------------------------------------------
| Authentication Controller
|--------------------------------------------------------------------------
|
| GetCredentials works on oauth2 Client Credentials Grant and returns CLIENT_ID, CLIENT_SECRET
| GetToken takes CLIENT_ID, CLIENT_SECRET, grant_type, scope and returns access_token and some other information
*/

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/gin-server"
	"github.com/google/uuid"
	letslog "github.com/letsgo-framework/letsgo/log"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

var clientStore = store.NewClientStore()
var manager = manage.NewDefaultManager()

// AuthInit initializes authentication
func AuthInit() {
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapClientStorage(clientStore)

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)
}

// GetCredentials sends client credentials
func GetCredentials(c *gin.Context) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:8000",
	})
	if err != nil {
		letslog.Error(err.Error())
	}
	c.JSON(200, map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
	c.Done()
}

// GetToken sends accecc_token
func GetToken(c *gin.Context) {
	ginserver.HandleTokenRequest(c)
}

// Verify accessToken with client
func Verify(c *gin.Context) {
	ti, exists := c.Get(ginserver.DefaultConfig.TokenKey)
	if exists {
		c.JSON(200, ti)
		return
	}
	c.String(200, "not found")
}
