package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lib/pq"
	"github.com/mrasoolmirzaei/words/backend/internal/db"
	"github.com/sirupsen/logrus"
)

const (
	UniquenessViolation = "23505"
	NotFound            = "20000"
	CheckViolation      = "23514"

	WordPrefix = "Word"
	SynonymPrefix = "Synonym"
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
	errMsg, isInvalid := req.Validate()
	if isInvalid {
		a.log.Debugf("invalid add word request: %v", errMsg)
		return nil, &Error{
			Message:  errMsg,
			HttpCode: http.StatusBadRequest,
		}
	}

	dbWord, err := a.db.AddWord(req.Title.String())
	if err != nil {
		a.log.Errorf("failed to add word: %v", err)
		return nil, customizeError(err, WordPrefix)
	}

	return &AddWordResponse{
		Word: &Word{
			ID:    dbWord.ID,
			Title: dbWord.Title,
		}}, nil
}

func (a *API) AddSynonym(req AddSynonymRequest) *Error {
	errMsg, isInvalid := req.Validate()
	if isInvalid {
		a.log.Debugf("invalid add synonym request: %v", errMsg)
		return &Error{
			Message:  errMsg,
			HttpCode: http.StatusBadRequest,
		}
	}
	word, err := a.db.SearchWord(req.WordTitle.String())
	if err != nil {
		a.log.Errorf("failed to find first word %v: %v", req.WordTitle, err)
		return customizeError(err, WordPrefix)
	}

	synonym, err := a.db.SearchWord(req.SynonymTitle.String())
	if err != nil {
		a.log.Errorf("failed to find second word: %v", err)
		return customizeError(err, WordPrefix)
	}

	if word.ID > synonym.ID {
		word, synonym = synonym, word
	}
	err = a.db.AddSynonym(word.ID, synonym.ID)
	if err != nil {
		a.log.Errorf("failed to add synonym: %v", err)
		return customizeError(err, SynonymPrefix)
	}

	return nil
}

func (a *API) GetSynonyms(req GetSynonymsRequest) (*GetSynonymsResponse, *Error) {
	errMsg, isInvalid := req.Validate()
	if isInvalid {
		a.log.Debugf("invalid get synonyms request: %v", errMsg)
		return nil, &Error{
			Message:  errMsg,
			HttpCode: http.StatusBadRequest,
		}
	}
	
	word, err := a.db.SearchWord(req.WordTitle.String())
	if err != nil {
		a.log.Errorf("failed to find word: %v", err)
		return nil, customizeError(err, WordPrefix)
	}

	synonyms, err := a.db.GetSynonyms(word.ID)
	if err != nil {
		a.log.Errorf("failed to get synonyms: %v", err)
		return nil, customizeError(err, SynonymPrefix)
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

func customizeError(err error, prefix string) *Error {
	customizedError := &Error{
		Message:  "internal server error",
		HttpCode: 500,
	}

	if err == sql.ErrNoRows {
		customizedError.Message = fmt.Sprintf("%v not found", prefix)
		customizedError.HttpCode = http.StatusNotFound
		return customizedError
	}

	pgErr, ok := err.(*pq.Error)
	if !ok {
		return customizedError
	}
	customizedError.DBErrorCode = string(pgErr.Code)

	switch pgErr.Code {
	case NotFound:
		customizedError.Message = fmt.Sprintf("%v not found", prefix)
		customizedError.HttpCode = http.StatusNotFound
	case UniquenessViolation:
		customizedError.Message = fmt.Sprintf("%v already exists", prefix)
		customizedError.HttpCode = http.StatusConflict
	}

	return customizedError
}
