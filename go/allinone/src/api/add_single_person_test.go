package api

import (
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"playground/allinone/model"
	"playground/allinone/sample"
	"playground/allinone/util"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type APIAddSinglePersonTestSuite struct {
	suite.Suite
}

func (s *APIAddSinglePersonTestSuite) SetupSuite() {
	s.T().Log(util.CurFuncDesc())
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s *APIAddSinglePersonTestSuite) SetupTest() {
	s.T().Log(util.CurFuncDesc())
	model.EmptyDefaultMatchService()
}

func (s *APIAddSinglePersonTestSuite) TestAddSinglePerson() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(AddSinglePersonAndMatch))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	user := sample.GetSampleMaleUsers()[0]
	url := model.BuildURL_AddSinglePersonAndMatch(
		server.URL,
		user.Name(),
		user.Height(),
		user.Gender(),
		user.NumDates(),
	)
	s.T().Log("POST", url)
	res, err := http.Post(url, "text/plain", nil)
	s.Equal(http.StatusNotFound, res.StatusCode)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Logf("Response: %v, Error:%v", res, err)

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Payload:", string(payload))
}

func (s *APIAddSinglePersonTestSuite) TearDownTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIAddSinglePersonTestSuite) TearDownSuite() {
	s.T().Log(util.CurFuncDesc())
}

func TestAPIAddSinglePerson(t *testing.T) {
	suite.Run(t, new(APIAddSinglePersonTestSuite))
}
