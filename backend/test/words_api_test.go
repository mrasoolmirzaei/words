package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mrasoolmirzaei/words/backend/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const (
	addWordApiURL     = "http://localhost:8090/word"
	addSynonymApiURL  = "http://localhost:8090/synonym/"
	getSynonymsApiURL = "http://localhost:8090/synonyms/test"
	contentType       = "application/json"
)

func (suite *testSuite) TestAddWordHappyCases() {
	type request struct {
		Title string `json:"title"`
	}

	cases := []struct {
		testName string
		request  request
		expected api.AddWordResponse
	}{
		{
			testName: "add word",
			request:  request{Title: "test"},
			expected: api.AddWordResponse{
				Word: &api.Word{
					ID:    1,
					Title: "test",
				},
			},
		},
	}

	for _, tc := range cases {
		body, err := json.Marshal(tc.request)
		require.NoError(suite.T(), err)
		res, err := http.Post(addWordApiURL, contentType, bytes.NewBuffer(body))
		require.NoError(suite.T(), err)
		defer res.Body.Close()

		response := api.AddWordResponse{}
		err = json.NewDecoder(res.Body).Decode(&response)
		require.NoError(suite.T(), err)

		assert.Equal(suite.T(), tc.expected, response)
	}
}

func (suite *testSuite) TestAddWordFailedCases() {
	type request struct {
		Title string `json:"title"`
	}

	cases := []struct {
		testName string
		request  request
		expected int
	}{
		{
			testName: "add word",
			request:  request{Title: "erroruniqueviolation"},
			expected: http.StatusConflict,
		},
	}

	for _, tc := range cases {
		body, err := json.Marshal(tc.request)
		require.NoError(suite.T(), err)
		res, err := http.Post(addWordApiURL, contentType, bytes.NewBuffer(body))
		require.NoError(suite.T(), err)
		defer res.Body.Close()

		assert.Equal(suite.T(), tc.expected, res.StatusCode)
	}
}

func (suite *testSuite) TestAddSynonym() {
	type request struct {
		WordTitle    string
		SynonymTitle string `json:"synonym"`
	}

	cases := []struct {
		testName string
		request  request
		expected int
	}{
		{
			testName: "add already created synonym",
			request:  request{WordTitle: "testone", SynonymTitle: "testtwo"},
			expected: http.StatusConflict,
		},
		{
			testName: "violate ID constraint",
			request:  request{WordTitle: "testtwo", SynonymTitle: "testone"},
			expected: http.StatusConflict,
		},
		{
			testName: "add synonym",
			request:  request{WordTitle: "testone", SynonymTitle: "testthree"},
			expected: http.StatusCreated,
		},
		{
			testName: "add not found synonym",
			request:  request{WordTitle: "testone", SynonymTitle: "errornotfound"},
			expected: http.StatusNotFound,
		},
	}

	for _, tc := range cases {
		body, err := json.Marshal(tc.request)
		require.NoError(suite.T(), err)
		res, err := http.Post(addSynonymApiURL+tc.request.WordTitle, contentType, bytes.NewBuffer(body))
		require.NoError(suite.T(), err)
		defer res.Body.Close()

		assert.Equal(suite.T(), tc.expected, res.StatusCode)
	}
}

func (suite *testSuite) TestGetSynonyms() {
	type request struct {
		WordTitle string `json:"word"`
	}

	cases := []struct {
		testName string
		request  request
		expected api.GetSynonymsResponse
	}{
		{
			testName: "get synonyms",
			request:  request{WordTitle: "test"},
			expected: api.GetSynonymsResponse{
				Synonyms: []api.Word{{ID: 1, Title: "test"}},
			},
		},
	}

	for _, tc := range cases {
		res, err := http.Get(getSynonymsApiURL)
		require.NoError(suite.T(), err)
		defer res.Body.Close()

		response := api.GetSynonymsResponse{}
		err = json.NewDecoder(res.Body).Decode(&response)
		require.NoError(suite.T(), err)

		assert.Equal(suite.T(), tc.expected, response)
	}
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(testSuite))
}
