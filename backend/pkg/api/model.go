package api

type AddWordRequest struct {
	Title string `json:"title"`
}

type AddWordResponse struct {
	Title string `json:"title"`
}

type AddSynonymRequest struct {
	WordTitle  string
	SynonymTitle string `json:"synonym"`
}
