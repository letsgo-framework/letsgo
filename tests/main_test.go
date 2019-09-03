package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
	letslog "github.com/letsgo-framework/letsgo/log"
	"github.com/letsgo-framework/letsgo/routes"
	"github.com/letsgo-framework/letsgo/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	. "gopkg.in/check.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

type TestInsert struct {
	Name string `form:"name" binding:"required" json:"name" bson:"name"`
}

type TestSuite struct {
	srv *gin.Engine
}

var _ = Suite(&TestSuite{})

func TestMain(m *testing.M) {
	// Setup log writing
	letslog.InitLogFuncs()
	err := godotenv.Load("../.env.testing")
	database.TestConnect()

	database.DB.Drop(context.Background())

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s := TestSuite{
		srv: routes.PaveRoutes(),
	}
	go s.srv.Run(os.Getenv("PORT"))

	os.Exit(m.Run())
}


func (s *TestSuite) TestGetEnv(c *C) {

	dbPort := os.Getenv("DATABASE_PORT")
	fmt.Printf("db port %s", dbPort)
	if dbPort == "" {
		c.Error()
		c.Fail()
	}
	c.Assert(dbPort, Equals, "27017")
}

func (s *TestSuite) TestHelloWorld(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestCredentials(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestTokenSuccess(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestTokenFail(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestAccessTokenSuccess(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/auth?access_token=" + tokenResponse.AccessToken

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestAccessTokenFail(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/credentials/"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, _ := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=client_credentials&client_id=" + credResponse.CLIENT_ID + "&client_secret=" + credResponse.CLIENT_SECRET + "&scope=read"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	respData, _ := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/auth?access_token=mywrongaccesstoken"

	req, _ = http.NewRequest("GET", requestURL, nil)

	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestDatabaseTestConnection(c *C) {
	database.TestConnect()
	err := database.Client.Ping(context.Background(), readpref.Primary())
	c.Assert(err, Equals, nil)
}

func (s *TestSuite) TestDatabaseConnection(c *C) {
	database.Connect()
	err := database.Client.Ping(context.Background(), readpref.Primary())
	c.Assert(err, Equals, nil)
}

func (s *TestSuite) TestDBInsert(c *C) {
	database.TestConnect()
	input := TestInsert{Name: "testname"}
	collection := database.DB.Collection("test_collection")
	_, err := collection.InsertOne(context.Background(), input)
	if err != nil {
		c.Error(err)
	}
	result := TestInsert{}
	err = collection.FindOne(context.Background(), bson.M{"name": "testname"}).Decode(&result)
	if err != nil {
		c.Error(err)
	}
	c.Assert(result, Equals, input)
}

//Test User-registartion
func (s *TestSuite) Test1UserRegistration(c *C) {
	data := types.User{
		Name:     "Letsgo User",
		Username:    "letsgoUs3r",
		Password: "qwerty",
	}

	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/register"
	client := &http.Client{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	req, _ := http.NewRequest("POST", requestURL, b)
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	respData, _ := ioutil.ReadAll(resp.Body)
	var user types.User
	json.Unmarshal(respData, &user)
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) Test2UserLoginPasswordGrant(c *C) {
	requestURL := "http://127.0.0.1" + os.Getenv("PORT") + "/api/v1/login?grant_type=password&client_id=client@letsgo&client_secret=Va4a8bFFhTJZdybnzyhjHjj6P9UVh7UL&scope=read&username=letsgoUs3r&password=qwerty"
	req, _ := http.NewRequest("GET", requestURL, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		letslog.Debug(err.Error())
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	c.Assert(resp.StatusCode, Equals, 200)
}

func Test(t *testing.T) {
	TestingT(t)
}
