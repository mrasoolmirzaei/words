package api

type AddWordRequest struct {
	Title string `json:"title"`
}

type AddWordResponse struct {
	Word *Word `json:"word"`
}

type AddSynonymRequest struct {
	WordTitle    string
	SynonymTitle string `json:"synonym"`
}

type GetSynonymsRequest struct {
	WordTitle string `json:"word"`
}

type GetSynonymsResponse struct {
	Synonyms []Word `json:"synonyms"`
}

type Word struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Error struct {
	Message     string
	HttpCode    int
	DBErrorCode string
}
