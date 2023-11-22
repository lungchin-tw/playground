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

type APIMatchTestSuite struct {
	suite.Suite
}

func (s *APIMatchTestSuite) SetupSuite() {
	s.T().Log(util.CurFuncDesc())
	rand.Seed(time.Now().UTC().UnixNano())
}

func (s *APIMatchTestSuite) SetupTest() {
	s.T().Log(util.CurFuncDesc())
	model.EmptyDefaultMatchService()
}

func (s *APIMatchTestSuite) TestMatchSuccess() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(AddSinglePersonAndMatch))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	for _, v := range sample.GetSampleMaleUsers() {
		url := model.BuildURL_AddSinglePersonAndMatch(
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

	source := sample.GetSampleFemaleUsers()[0]
	url := model.BuildURL_AddSinglePersonAndMatch(
		server.URL,
		source.Name(),
		source.Height(),
		source.Gender(),
		source.NumDates(),
	)
	s.T().Log("POST", url)
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Logf("Response: %v, Error:%v", res, err)

	payload, err := io.ReadAll(res.Body)
	if err != nil {
		s.T().Fatal(err)
	}

	s.T().Log("Response Payload:", string(payload))
	s.Equal(http.StatusOK, res.StatusCode)
}

func (s *APIMatchTestSuite) TestMatchUntilFail() {
	s.T().Log(util.CurFuncDesc())

	server := httptest.NewServer(http.HandlerFunc(AddSinglePersonAndMatch))

	s.T().Log("Server Started.", server.URL)
	defer server.Close()

	// Add Male
	{
		male := sample.GetSampleMaleUsers()[0]

		url := model.BuildURL_AddSinglePersonAndMatch(
			server.URL,
			male.Name(),
			male.Height(),
			male.Gender(),
			male.NumDates(),
		)
		s.T().Log("POST", url)
		res, err := http.Post(url, "text/plain", nil)
		if err != nil {
			s.T().Fatal(err)
		}

		s.T().Logf("Add: %v", male.Name())
		s.Equal(http.StatusNotFound, res.StatusCode)
	}

	counter := 0
	source := sample.GetSampleFemaleUsers()[0]
	for {
		url := model.BuildURL_AddSinglePersonAndMatch(
			server.URL,
			source.Name(),
			source.Height(),
			source.Gender(),
			source.NumDates(),
		)
		s.T().Log("POST", url)
		res, err := http.Post(url, "text/plain", nil)
		if err != nil {
			s.T().Fatal(err)
		}

		s.T().Logf("Response: %v, Error:%v", res, err)

		payload, err := io.ReadAll(res.Body)
		if err != nil {
			s.T().Fatal(err)
		}

		s.T().Log("Response Payload:", string(payload))
		if res.StatusCode == http.StatusOK {
			counter++
		} else {
			break
		}
	}

	s.Equal(1, counter)
}

func (s *APIMatchTestSuite) TearDownTest() {
	s.T().Log(util.CurFuncDesc())
}

func (s *APIMatchTestSuite) TearDownSuite() {
	s.T().Log(util.CurFuncDesc())
}

func TestAPIMatch(t *testing.T) {
	suite.Run(t, new(APIMatchTestSuite))
}
