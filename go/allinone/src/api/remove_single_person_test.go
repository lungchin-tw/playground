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

	samples := sample.GetSampleMaleUsers()
	samples = append(samples, sample.GetSampleFemaleUsers()[0])
	for _, v := range samples {
		model.DefaultMatchService().AddUser(v)
		s.T().Logf("Add: %v", v.Name())
	}
}

func (s *APIRemoveSinglePersonTestSuite) TestRemovePerson() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(RemoveSinglePerson))

	s.T().Log("Server Started:", server.URL)
	defer server.Close()

	source := sample.GetSampleFemaleUsers()[0]
	before := model.DefaultMatchService().FindPossiblePeopleByName(source.Name())
	s.Greater(len(before), 0)
	s.T().Log("Before:", before)

	url := model.BuildURL_RemoveSinglePerson(
		server.URL,
		source.Name(),
	)
	s.T().Log(http.MethodPost, url)
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		s.T().Fatal(err)
	}

	s.Equal(http.StatusOK, res.StatusCode)

	after := model.DefaultMatchService().FindPossiblePeopleByName(source.Name())
	s.Equal(0, len(after))
	s.T().Log("After:", after)
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
