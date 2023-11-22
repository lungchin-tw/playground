package api

import (
	"io/ioutil"
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

type APIQuerySinglePersonTestSuite struct {
	suite.Suite
}

func (s *APIQuerySinglePersonTestSuite) SetupSuite() {
	s.T().Log(util.CurFuncDesc())
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s *APIQuerySinglePersonTestSuite) SetupTest() {
	s.T().Log(util.CurFuncDesc())
	model.EmptyDefaultMatchService()

	server := httptest.NewServer(http.HandlerFunc(AddSinglePersonAndMatch))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	samples := sample.GetSampleMaleUsers()
	samples = append(samples, sample.GetSampleFemaleUsers()[0])
	for _, v := range samples {
		model.DefaultMatchService().AddUser(v)
		s.T().Logf("Add: %v", v.Name())
	}

}

func (s *APIQuerySinglePersonTestSuite) TestQuery() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(QuerySinglePerson))
	defer server.Close()

	source := sample.GetSampleFemaleUsers()[0]
	url := model.BuildURL_QuerySinglePerson(
		server.URL,
		source.Name(),
	)
	s.T().Log("GET", url)
	res, err := http.Get(url)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Logf("Response: %v, Error:%v", res, err)

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Response Payload:", string(payload))
	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *APIQuerySinglePersonTestSuite) TearDownTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIQuerySinglePersonTestSuite) TearDownSuite() {
	s.T().Log(util.CurFuncDesc())
}

func TestAPIQuerySinglePerson(t *testing.T) {
	suite.Run(t, new(APIQuerySinglePersonTestSuite))
}
