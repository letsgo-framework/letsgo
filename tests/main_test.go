package tests

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/letsgo/routes"
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
	s.srv = routes.PaveRoutes()
	go s.srv.Run(":8084")
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func (s *MySuite) TestHelloWorld(c *C) {
	requestURL := "http://127.0.0.1:8084/api/v1/"
	resp, err := http.Get(requestURL)
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