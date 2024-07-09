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
	addWordApiURL     = "http://localhost:8080/word"
	addSynonymApiURL  = "http://localhost:8080/synonym"
	getSynonymsApiURL = "http://localhost:8080/synonyms"
	contentType       = "application/json"
)

func (suite *testSuite) TestAddWord() {
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

func TestIntegration(t *testing.T) {
	suite.Run(t, new(testSuite))
}
