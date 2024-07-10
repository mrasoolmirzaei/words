package db

import (
	"database/sql"

	"github.com/lib/pq"
)

type PQMock struct{}

func (db *PQMock) Close() error {
	return nil
}

func (db *PQMock) AddWord(title string) (*Word, error) {
	// 23505 is a unique violation.
	if title == "erroruniqueviolation" {
		return nil, &pq.Error{Code: "23505"}
	}

	return &Word{ID: 1, Title: title}, nil
}

func (db *PQMock) AddSynonym(word_id_1, word_id_2 int) error {
	// 23514 is a check violation, here we are using it to simulate a foreign key violation,
	// word_id_1 should be greater than word_id_2.
	if word_id_1 == 2 && word_id_2 == 1 {
		return &pq.Error{Code: "23514"}
	}
	// 23505 is a unique violation.
	if word_id_1 == 1 && word_id_2 == 2 {
		return &pq.Error{Code: "23505"}
	}

	return nil
}

func (db *PQMock) SearchWord(title string) (*Word, error) {
	if title == "errornotfound" {
		return nil, sql.ErrNoRows
	}
	if title == "testone" {
		return &Word{ID: 1, Title: "testone"}, nil
	}
	if title == "testtwo" {
		return &Word{ID: 2, Title: "testtwo"}, nil
	}

	return &Word{ID: 3, Title: "testthree"}, nil

}

func (db *PQMock) GetSynonyms(id int) ([]*Word, error) {
	return []*Word{{ID: 1, Title: "test"}}, nil
}
