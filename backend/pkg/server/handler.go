package server

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"

	"github.com/mrasoolmirzaei/words/backend/pkg/api"
)

const (
	addSynonymPathParam = "word"
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

		resp, err := s.api.AddWord(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
		req.WordTitle = word

		err = s.api.AddSynonym(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
