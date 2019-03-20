package tests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/letsgo/routes"
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

func Test(t *testing.T) {
	TestingT(t)
}