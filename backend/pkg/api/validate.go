package api

import (
	"fmt"
	"strings"
)

func (r GetSynonymsRequest) Validate() (string, bool) {
	var errorMsg string
	isInvalid := false

	if r.WordTitle == "" {
		errorMsg = "input word is empty"
		isInvalid = true
		return errorMsg, isInvalid
	}

	errorMsg, isInvalid = r.WordTitle.Validate()

	return errorMsg, isInvalid
}

func (r AddSynonymRequest) Validate() (string, bool) {
	var errorMsg string
	isInvalid := false

	if r.WordTitle == "" {
		errorMsg = "input word is empty"
		isInvalid = true
		return errorMsg, isInvalid
	}

	errorMsg, isInvalid = r.WordTitle.Validate()
	if isInvalid {
		return errorMsg, isInvalid
	}

	if r.SynonymTitle == "" {
		errorMsg = "input synonym is empty"
		isInvalid = true
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
	var errorMsg string
	isInvalid := false

	if r.Title == "" {
		errorMsg = "input word is empty"
		isInvalid = true
		return errorMsg, isInvalid
	}

	errorMsg, isInvalid = r.Title.Validate()

	return errorMsg, isInvalid
}

func (w InputWord) Validate() (string, bool) {
	var errorMsg string
	isInvalid := false

	w.PreProcess()

	// Longest word in English dictionary is 45 characters
	if len(w) > 45 || len(w) < 2  {
		errorMsg = fmt.Sprintf("input word is too long or too short, maximum 45 and minimum 2 characters. current length: %v", len(w))
		isInvalid = true
	}
	// Check if all characters are alphabets
	for _, c := range w {
		if c < 'a' || c > 'z' {
			errorMsg = fmt.Sprintf("input word should contain only alphabets. invalid character: %v", c)
			isInvalid = true
			break
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