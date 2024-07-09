package api

type AddWordRequest struct {
	Title InputWord `json:"title"`
}

type AddWordResponse struct {
	Word *Word `json:"word"`
}

type AddSynonymRequest struct {
	WordTitle    InputWord
	SynonymTitle InputWord `json:"synonym"`
}

type GetSynonymsRequest struct {
	WordTitle InputWord `json:"word"`
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

type InputWord string
