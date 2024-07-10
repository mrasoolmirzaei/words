package db

type DBMock struct{}

func (db *DBMock) Close() error {
	return nil
}

func (db *DBMock) AddWord(title string) (*Word, error) {
	return &Word{ID: 1, Title: title}, nil
}

func (db *DBMock) AddSynonym(word_id_1, word_id_2 int) error {
	return nil
}

func (db *DBMock) SearchWord(title string) (*Word, error) {
	return &Word{ID: 1, Title: title}, nil
}

func (db *DBMock) GetSynonyms(id int) ([]*Word, error) {
	return []*Word{{ID: 1, Title: "test"}}, nil
}
