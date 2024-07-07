package api

import (
	"github.com/mrasoolmirzaei/words/backend/internal/db"
	"github.com/sirupsen/logrus"
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

func (a *API) AddWord(req AddWordRequest) (*AddWordResponse, error) {
	word, err := a.db.AddWord(req.Title)
	if err != nil {
		a.log.Errorf("failed to add word: %v", err)
		return nil, err
	}

	return &AddWordResponse{
		Title: word.Title,
	}, nil
}

func (a *API) AddSynonym(req AddSynonymRequest)  error {
	word, err := a.db.SearchWord(req.WordTitle)
	if err != nil {
		a.log.Errorf("failed to find first word %v: %v",req.WordTitle, err)
		return err
	}

	synonym, err := a.db.SearchWord(req.SynonymTitle)
	if err != nil {
		a.log.Errorf("failed to find second word: %v", err)
		return err
	}

	err = a.db.AddSynonym(word.ID, synonym.ID)
	if err != nil {
		a.log.Errorf("failed to add synonym: %v", err)
		return err
	}

	return nil
}
