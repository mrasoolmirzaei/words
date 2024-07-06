package db

type Word struct {
	// ID of the word.
	ID   int
	// Title of the word.    
	Title string 
}

type Synonym struct {
	// ID of the first word.
	Word_ID_1   int    
	// ID of the second word.
	Word_ID_2   int
}