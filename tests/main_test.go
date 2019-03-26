package tests

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/letsgo/routes"
	"gitlab.com/letsgo/types"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	. "gopkg.in/check.v1"
)

type MySuite struct{
	srv *gin.Engine
}

var _ = Suite(&MySuite{})

func (s *MySuite) SetUpTest(c *C) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	s.srv = routes.PaveRoutes()
	go s.srv.Run(":8084")
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func (s *MySuite) TestGetEnv(c *C) {

	dbPort := os.Getenv("DATABASE_PORT")
	fmt.Printf("db port %s",dbPort)
	if dbPort == "" {
		c.Error()
		c.Fail()
	}
	c.Assert(dbPort, Equals, "27017")
}


func (s *MySuite) TestNoApiToken(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/"
	resp, err := http.Get(requestURL)
	if err != nil {
		c.Error(err)
		c.Fail()
	}
	defer resp.Body.Close()
	c.Assert(resp.StatusCode, Equals, 401)
}

func (s *MySuite) TestApiTokenMismatch(c *C) {
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

func (s *MySuite) TestHelloWorld(c *C) {
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


func (s *MySuite) TestCredentials(c *C) {
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

func (s *MySuite) TestTokenSuccess(c *C) {
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

func (s *MySuite) TestTokenFail(c *C) {
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

func (s *MySuite) TestAccessTokenSuccess(c *C) {
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


func (s *MySuite) TestAccessTokenFail(c *C) {
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


func Test(t *testing.T) {
	TestingT(t)
}