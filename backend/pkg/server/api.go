package server

import (
	"encoding/json"
	"net/http"

	m "github.com/mrasoolmirzaei/words/backend/pkg/api_models"
)

func (s *Server) addWord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := m.AddWordRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			s.log.Errorf("invalid request format: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		word, err := s.db.AddWord(req.Title)
		if err != nil {
			s.log.Errorf("failed to add word: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp := m.AddWordResponse{
			Title: word.Title,
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
