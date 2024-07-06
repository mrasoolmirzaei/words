package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storer interface {
	Close() error

	// GetWord returns a word by its ID.
	GetWord(id string) (*Word, error)

	// AddWord adds a new word to the database.
	AddWord(title string) (*Word, error)
	// AddSynonym adds a new synonym relationship to the database.
	AddSynonym(word_id_1, word_id_2 string) error

	// SearchWord searches for a word by its title.
	SearchWord(title string) (*Word, error)
	// SearchSynonyms searches for synonyms of a word by its ID.
	SearchSynonyms(id string) ([]*Word, error)
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
	getWordSQL := `SELECT id, title FROM words WHERE id = $1`
	row := db.conn.QueryRow(getWordSQL, id)

	var w Word
	err := row.Scan(&w.ID, &w.Title)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (db *DB) AddWord(title string) (*Word, error) {
	addWordSQL := `INSERT INTO words (title) VALUES ($1) RETURNING id, title`
	row := db.conn.QueryRow(addWordSQL, title)

	word := &Word{}
	err := row.Scan(word.ID, word.Title)
	if err != nil {
		return nil, err
	}

	return word, nil
}

func (db *DB) AddSynonym(word_id_1, word_id_2 string) error {
	addSynonymSQL := `INSERT INTO synonyms (word_id_1, word_id_2) VALUES ($1, $2)`
	_, err := db.conn.Exec(addSynonymSQL, word_id_1, word_id_2)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) SearchWord(title string) (*Word, error) {
	searchWordSQL := `SELECT id, title FROM words WHERE title = $1`
	row := db.conn.QueryRow(searchWordSQL, title)

	var w Word
	err := row.Scan(&w.ID, &w.Title)
	if err != nil {
		return nil, err
	}

	return &w, nil
}

func (db *DB) SearchSynonyms(id string) ([]*Word, error) {
	searchSynonymsSQL := `SELECT id, title FROM words WHERE id IN (SELECT word_id_2 FROM synonyms WHERE word_id_1 = $1)`
	rows, err := db.conn.Query(searchSynonymsSQL, id)
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

