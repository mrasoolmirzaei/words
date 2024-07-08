package db

import (
	"database/sql"
	"log"

	_ "embed"
)

//go:embed sqls/addWord.sql
var addWordSQL string

//go:embed sqls/addSynonym.sql
var addSynonymSQL string

//go:embed sqls/getWord.sql
var getWordSQL string

//go:embed sqls/searchWord.sql
var searchWordSQL string

//go:embed sqls/getParentSynonyms.sql
var getParentSynonymsSQL string

//go:embed sqls/getChildSynonyms.sql
var getChildSynonymsSQL string

type Storer interface {
	Close() error

	// GetWord returns a word by its ID.
	GetWord(id string) (*Word, error)
	// GetDirectSynonyms returns a list of direct synonyms for a word.
	GetSynonyms(id int) ([]*Word, error)

	// AddWord adds a new word to the database.
	AddWord(title string) (*Word, error)
	// AddSynonym adds a new synonym relationship to the database.
	AddSynonym(word_1_id, word_2_id int) error

	// SearchWord searches for a word by its title.
	SearchWord(title string) (*Word, error)
}

type DB struct {
	conn *sql.DB
}

func NewDB(connStr string) (Storer, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open db connection: %v", err)
		return nil, err
	}

	return &DB{conn: db}, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) GetWord(id string) (*Word, error) {
	row := db.conn.QueryRow(getWordSQL, id)

	var w Word
	err := row.Scan(&w.ID, &w.Title)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (db *DB) AddWord(title string) (*Word, error) {
	row := db.conn.QueryRow(addWordSQL, title)

	word := &Word{}
	err := row.Scan(&word.ID, &word.Title)
	if err != nil {
		return nil, err
	}

	return word, nil
}

func (db *DB) AddSynonym(word_id_1, word_id_2 int) error {
	_, err := db.conn.Exec(addSynonymSQL, word_id_1, word_id_2)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) SearchWord(title string) (*Word, error) {
	row := db.conn.QueryRow(searchWordSQL, title)

	var w Word
	err := row.Scan(&w.ID, &w.Title)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (db *DB) GetSynonyms(id int) ([]*Word, error) {
	parentWords, err := db.listWordsByID(id, getParentSynonymsSQL)
	if err != nil {
		return nil, err
	}

	childWords, err := db.listWordsByID(id, getChildSynonymsSQL)
	if err != nil {
		return nil, err
	}

	words := append(parentWords, childWords...)

	return words, nil
}

func (db *DB) listWordsByID(id int, sqlQuery string) ([]*Word, error) {
	childrenRows, err := db.conn.Query(sqlQuery, id)
	if err != nil {
		return nil, err
	}
	defer childrenRows.Close()

	var words []*Word
	for childrenRows.Next() {
		var w Word
		err := childrenRows.Scan(&w.ID, &w.Title)
		if err != nil {
			return nil, err
		}
		words = append(words, &w)
	}

	return words, nil
}
