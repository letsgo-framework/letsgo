package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/letsgo-framework/letsgo/database"
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

type TestSuite struct{
	srv *gin.Engine
}

var _ = Suite(&TestSuite{})

func (s *TestSuite) SetUpTest(c *C) {
	err := godotenv.Load("../.env.testing")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s.srv = routes.PaveRoutes()
	go s.srv.Run(os.Getenv("PORT"))
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func (s *TestSuite) TestGetEnv(c *C) {

	dbPort := os.Getenv("DATABASE_PORT")
	fmt.Printf("db port %s",dbPort)
	if dbPort == "" {
		c.Error()
		c.Fail()
	}
	c.Assert(dbPort, Equals, "27017")
}


func (s *TestSuite) TestNoApiToken(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/"
	resp, err := http.Get(requestURL)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestApiTokenMismatch(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,`lablabla`)
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestHelloWorld(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}


func (s *TestSuite) TestCredentials(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/credentials/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestTokenSuccess(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/credentials/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1:8084/api/v1/token?grant_type=client_credentials&client_id="+credResponse.CLIENT_ID+"&client_secret="+credResponse.CLIENT_SECRET+"&scope=read"

	req, err = http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 200)
}

func (s *TestSuite) TestTokenFail(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/token"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestAccessTokenSuccess(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/credentials/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1:8084/api/v1/token?grant_type=client_credentials&client_id="+credResponse.CLIENT_ID+"&client_secret="+credResponse.CLIENT_SECRET+"&scope=read"

	req, err = http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1:8084/api/v1/auth?access_token="+tokenResponse.Access_token

	req, err = http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()


	c.Assert(resp.StatusCode, Equals, 200)
}


func (s *TestSuite) TestAccessTokenFail(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/credentials/"
	client := &http.Client{}
	req, err := http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err := client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	responseData, err := ioutil.ReadAll(resp.Body)
	var credResponse types.CredentialResponse
	json.Unmarshal(responseData, &credResponse)

	requestURL = "http://127.0.0.1:8084/api/v1/token?grant_type=client_credentials&client_id="+credResponse.CLIENT_ID+"&client_secret="+credResponse.CLIENT_SECRET+"&scope=read"

	req, err = http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()

	respData, err := ioutil.ReadAll(resp.Body)
	var tokenResponse types.TokenResponse
	json.Unmarshal(respData, &tokenResponse)

	requestURL = "http://127.0.0.1:8084/api/v1/auth?access_token=mywrongaccesstoken"

	req, err = http.NewRequest("GET", requestURL, nil)
	req.Header.Set(`X_API_KEY`,os.Getenv("X_API_KEY"))
	resp, err = client.Do(req)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()


	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *TestSuite) TestDatabaseConnection(c *C) {
	database.TestConnect()
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

func Test(t *testing.T) {
	TestingT(t)
}