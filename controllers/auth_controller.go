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
	"github.com/letsgo-framework/letsgo/database"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"time"
)

var clientStore = store.NewClientStore()
var manager = manage.NewDefaultManager()

// AuthInit initializes authentication
func AuthInit() {
	cfg := &manage.Config{
		// access token expiration time
		AccessTokenExp: time.Hour * 2,
		// refresh token expiration time
		RefreshTokenExp: time.Hour * 24 * 7,
		// whether to generate the refreshing token
		IsGenerateRefresh: true,
	}

	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.SetPasswordTokenCfg(cfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	manager.MapClientStorage(clientStore)

	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	ginserver.SetPasswordAuthorizationHandler(login)
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

// register
func Register (c *gin.Context) {
	a := types.User{}
	ctx := context.Background()
	collection := database.DB.Collection("users")
	err := c.BindJSON(&a)
	a.Password,_ = generateHash(a.Password)
	a.Id = primitive.NewObjectID()


	if err != nil {
		letslog.Error(err.Error())
		c.Abort()
	}
	_, err = collection.InsertOne(ctx, a)
	if err != nil {
		letslog.Error(err.Error())
		c.Abort()
	}
	c.JSON(200, a)
	c.Done()
}

// Generate a salted hash for the input string
func generateHash(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

// Compare string to generated hash
func compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)

	return bcrypt.CompareHashAndPassword(existing, incoming)
}


func login (username, password string) (userID string, err error) {

	collection := database.DB.Collection("users")

	user := types.User{}
	err = collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		letslog.Error(err.Error())
		return userID, err
	}
	loginError := compare(user.Password, password)

	if loginError != nil {
		letslog.Error(loginError.Error())
		return userID, err
	} else {
		userID = user.Id.Hex()
		return userID, err
	}


}