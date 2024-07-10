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

//go:embed sqls/searchWord.sql
var searchWordSQL string

//go:embed sqls/getSynonyms.sql
var getSynonymsSQL string

type Storer interface {
	Close() error

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

// AddWord adds a new word to the database.
func (db *DB) AddWord(title string) (*Word, error) {
	row := db.conn.QueryRow(addWordSQL, title)

	word := &Word{}
	err := row.Scan(&word.ID, &word.Title)
	if err != nil {
		return nil, err
	}

	return word, nil
}

// AddSynonym adds a new synonym relationship to the database.
func (db *DB) AddSynonym(word_id_1, word_id_2 int) error {
	_, err := db.conn.Exec(addSynonymSQL, word_id_1, word_id_2)
	if err != nil {
		return err
	}

	return nil
}

// SearchWord searches for a word by its title.
func (db *DB) SearchWord(title string) (*Word, error) {
	row := db.conn.QueryRow(searchWordSQL, title)

	var w Word
	err := row.Scan(&w.ID, &w.Title)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

// GetSynonyms find synonyms for a word by its ID.
func (db *DB) GetSynonyms(id int) ([]*Word, error) {
	rows, err := db.conn.Query(getSynonymsSQL, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var words []*Word
	for rows.Next() {
		var w Word
		err := rows.Scan(&w.ID, &w.Title)
		if err != nil {
			return nil, err
		}
		words = append(words, &w)
	}

	return words, nil
}
