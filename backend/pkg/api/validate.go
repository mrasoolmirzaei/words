package api

import (
	"fmt"
	"strings"
)

const (
	// Longest word in English dictionary is 45 characters
	maximumWordLength = 45
	minimumWordLength = 2
)

func (r GetSynonymsRequest) Validate() (string, bool) {
	return r.WordTitle.Validate()
}

func (r AddSynonymRequest) Validate() (string, bool) {
	var errorMsg string
	isInvalid := false

	errorMsg, isInvalid = r.WordTitle.Validate()
	if isInvalid {
		return errorMsg, isInvalid
	}

	errorMsg, isInvalid = r.SynonymTitle.Validate()
	if isInvalid {
		return errorMsg, isInvalid
	}

	if r.WordTitle == r.SynonymTitle {
		errorMsg = "input word and synonym are the same"
		isInvalid = true
		return errorMsg, isInvalid
	}

	return errorMsg, isInvalid
}

func (r AddWordRequest) Validate() (string, bool) {
	return r.Title.Validate()
}

func (w InputWord) Validate() (string, bool) {
	var errorMsg string
	isInvalid := false

	// Pre-process the word string, remove leading and trailing spaces and convert to lowercase
	w.PreProcess()
	// Check if word is empty
	if len(w) == 0 {
		errorMsg = "input word is empty"
		isInvalid = true
		return errorMsg, isInvalid
	}
	// Check if word is too long or too short
	if len(w) > maximumWordLength || len(w) < minimumWordLength {
		errorMsg = fmt.Sprintf("input word is too long or too short, maximum 45 and minimum 2 characters. current length: %v", len(w))
		isInvalid = true
		return errorMsg, isInvalid
	}
	// Check if all characters are alphabets
	for _, c := range w {
		if c < 'a' || c > 'z' {
			errorMsg = fmt.Sprintf("input word should contain only alphabets. invalid character: %v", string(c))
			isInvalid = true
			return errorMsg, isInvalid
		}
	}

	return errorMsg, isInvalid
}

func (w *InputWord) PreProcess() {
	// Remove leading and trailing spaces
	*w = InputWord(strings.Join(strings.Fields(w.String()), " "))
	// Convert to lowercase
	*w = InputWord(strings.ToLower(w.String()))
}

func (w InputWord) String() string {
	return string(w)
}
