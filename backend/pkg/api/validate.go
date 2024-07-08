package api

import (
	"strings"
)

func (w InputWord) String() string {
	return string(w)
}

func (w *InputWord) PreProcess() {
	// Remove leading and trailing spaces
	*w = InputWord(strings.Join(strings.Fields(w.String()), " "))
	// Convert to lowercase
	*w = InputWord(strings.ToLower(w.String()))
}

func (w InputWord) Validate() (map[string]string, bool) {
	errors := make(map[string]string)
	isInvalid := false
	if len(w) == 0 {
		errors["empty word"] = "input word is empty"
		isInvalid = true
	}

	return errors, isInvalid
}