package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"playground/allinone/util"
	"testing"

	"github.com/stretchr/testify/suite"
)

type APIHealthTestSuite struct {
	suite.Suite
}

func (s *APIHealthTestSuite) SetupSuite() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIHealthTestSuite) SetupTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIHealthTestSuite) TestHealthSuccess() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(HandleHealth))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	s.T().Log("GET", server.URL)
	res, err := http.Get(server.URL)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Logf("Response: %v, Error:%v", res, err)

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Payload:", string(payload))
}

func (s *APIHealthTestSuite) TestHealthFail() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(HandleHealth))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	s.T().Log("POST", server.URL)
	res, err := http.Post(server.URL, "text/plain", nil)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Logf("Response: %v, Error:%v", res, err)

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Payload:", string(payload))
}

func (s *APIHealthTestSuite) TearDownTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIHealthTestSuite) TearDownSuite() {
	s.T().Log(util.CurFuncDesc())
}

func TestAPIHealth(t *testing.T) {
	suite.Run(t, new(APIHealthTestSuite))
}
