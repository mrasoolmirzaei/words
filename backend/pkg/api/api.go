package api

import (
	"net/http"

	"github.com/lib/pq"
	"github.com/mrasoolmirzaei/words/backend/internal/db"
	"github.com/sirupsen/logrus"
)

const (
	UniquenessViolation = "23505"
	NotFound            = "20000"
	CheckViolation      = "23514"
)

type API struct {
	db  db.Storer
	log logrus.FieldLogger
}

func NewAPI(db db.Storer, log logrus.FieldLogger) *API {
	return &API{
		db:  db,
		log: log,
	}
}

func (a *API) AddWord(req AddWordRequest) (*AddWordResponse, *Error) {
	word, err := a.db.AddWord(req.Title)
	if err != nil {
		a.log.Errorf("failed to add word: %v", err)
		return nil, customizeError(err)
	}

	return &AddWordResponse{
		Word: &Word{
			ID:    word.ID,
			Title: word.Title,
		}}, nil
}

func (a *API) AddSynonym(req AddSynonymRequest) *Error {
	word, err := a.db.SearchWord(req.WordTitle)
	if err != nil {
		a.log.Errorf("failed to find first word %v: %v", req.WordTitle, err)
		return customizeError(err)
	}

	synonym, err := a.db.SearchWord(req.SynonymTitle)
	if err != nil {
		a.log.Errorf("failed to find second word: %v", err)
		return customizeError(err)
	}

	err = a.db.AddSynonym(word.ID, synonym.ID)
	if err == nil {
		return nil
	}

	cErr := customizeError(err)
	if cErr.DBErrorCode != CheckViolation {
		a.log.Errorf("failed to add synonym: %v", err)
		return cErr
	}

	// here the error is a check violation, which means the synonym id is smaller than the word id
	// so we try to add the synonym in the opposite order
	// if it fails again, we return the error
	err = a.db.AddSynonym(synonym.ID, word.ID)
	if err != nil {
		a.log.Errorf("failed to add synonym: %v", err)
		return customizeError(err)
	}

	return nil
}

func (a *API) GetSynonyms(req GetSynonymsRequest) (*GetSynonymsResponse, *Error) {
	word, err := a.db.SearchWord(req.WordTitle)
	if err != nil {
		a.log.Errorf("failed to find word: %v", err)
		return nil, customizeError(err)
	}

	synonyms, err := a.db.GetSynonyms(word.ID)
	if err != nil {
		a.log.Errorf("failed to get synonyms: %v", err)
		return nil, customizeError(err)
	}

	resp := &GetSynonymsResponse{}
	for _, s := range synonyms {
		w := Word{
			ID:    s.ID,
			Title: s.Title,
		}
		resp.Synonyms = append(resp.Synonyms, w)
	}

	return resp, nil
}

func customizeError(err error) *Error {
	customizedError := &Error{
		Message:  "internal server error",
		HttpCode: 500,
	}

	pgErr, ok := err.(*pq.Error)
	if !ok {
		return customizedError
	}
	customizedError.DBErrorCode = string(pgErr.Code)

	switch pgErr.Code {
	case NotFound:
		customizedError.Message = "word not found"
		customizedError.HttpCode = http.StatusNotFound
	case UniquenessViolation:
		customizedError.Message = "word already exists"
		customizedError.HttpCode = http.StatusConflict
	}

	return customizedError
}
