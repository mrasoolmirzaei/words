package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrasoolmirzaei/words/backend/pkg/api"
)

const (
	addSynonymPathParam  = "word"
	getSynonymsPathParam = "word"
)

func (s *Server) addWord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := api.AddWordRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			s.log.Errorf("invalid request format: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, cErr := s.api.AddWord(req)
		if cErr != nil {
			http.Error(w, cErr.Message, cErr.HttpCode)
			return
		}

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			s.log.Errorf("failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) addSynonym() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := api.AddSynonymRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			s.log.Errorf("invalid request format: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		word, found := vars[addSynonymPathParam]
		if !found {
			errMsg := fmt.Sprintf("missing %v in parameters", addSynonymPathParam)
			s.log.Errorf(errMsg)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
		req.WordTitle = api.InputWord(word)

		cErr := s.api.AddSynonym(req)
		if cErr != nil {
			http.Error(w, cErr.Message, cErr.HttpCode)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) getSynonyms() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word, found := vars[getSynonymsPathParam]
		if !found {
			errMsg := fmt.Sprintf("missing %v in parameters", getSynonymsPathParam)
			s.log.Errorf(errMsg)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}

		req := api.GetSynonymsRequest{WordTitle: api.InputWord(word)}
		resp, cErr := s.api.GetSynonyms(req)
		if cErr != nil {
			http.Error(w, cErr.Message, cErr.HttpCode)
			return
		}

		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			s.log.Errorf("failed to encode response: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
