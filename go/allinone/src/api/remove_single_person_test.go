package api

import (
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

type APIRemoveSinglePersonTestSuite struct {
	suite.Suite
}

func (s *APIRemoveSinglePersonTestSuite) SetupSuite() {
	s.T().Log(util.CurFuncDesc())
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s *APIRemoveSinglePersonTestSuite) SetupTest() {
	s.T().Log(util.CurFuncDesc())
	model.EmptyDefaultMatchService()
}

func (s *APIRemoveSinglePersonTestSuite) TestQuery() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(QuerySinglePerson))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	for _, v := range sample.GetSampleMaleUsers() {
		url := model.BuildURL(
			server.URL,
			v.Name(),
			v.Height(),
			v.Gender(),
			v.NumDates(),
		)
		s.T().Log("POST", url)
		res, err := http.Post(url, "text/plain", nil)
		if err != nil {
			s.T().Fatal(err)
		}

		s.T().Logf("Add: %v", v.Name())
		s.Equal(http.StatusNotFound, res.StatusCode)
	}
}

func (s *APIRemoveSinglePersonTestSuite) TearDownTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIRemoveSinglePersonTestSuite) TearDownSuite() {
	s.T().Log(util.CurFuncDesc())
}

func TestAPIRemoveSinglePerson(t *testing.T) {
	suite.Run(t, new(APIRemoveSinglePersonTestSuite))
}
