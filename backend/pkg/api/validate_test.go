package api

import (
	"testing"
)

func TestGetSynonymsRequestValidate(t *testing.T) {
	tests := []struct {
		request GetSynonymsRequest
		wantErr bool
		errMsg  string
	}{
		{GetSynonymsRequest{WordTitle: ""}, true, "input word is empty"},
		{GetSynonymsRequest{WordTitle: "validword"}, false, ""},
	}

	for _, tt := range tests {
		errMsg, isInvalid := tt.request.Validate()
		if isInvalid != tt.wantErr || errMsg != tt.errMsg {
			t.Errorf("Validate() got = %v, %v, want %v, %v", errMsg, isInvalid, tt.errMsg, tt.wantErr)
		}
	}
}

func TestAddSynonymRequestValidate(t *testing.T) {
	tests := []struct {
		request AddSynonymRequest
		wantErr bool
		errMsg  string
	}{
		{AddSynonymRequest{WordTitle: "", SynonymTitle: ""}, true, "input word is empty"},
		{AddSynonymRequest{WordTitle: "validword", SynonymTitle: ""}, true, "input synonym is empty"},
		{AddSynonymRequest{WordTitle: "validword", SynonymTitle: "validsynonym"}, false, ""},
		{AddSynonymRequest{WordTitle: "sameword", SynonymTitle: "sameword"}, true, "input word and synonym are the same"},
	}

	for _, tt := range tests {
		errMsg, isInvalid := tt.request.Validate()
		if isInvalid != tt.wantErr || errMsg != tt.errMsg {
			t.Errorf("Validate() got = %v, %v, want %v, %v", errMsg, isInvalid, tt.errMsg, tt.wantErr)
		}
	}
}

func TestAddWordRequestValidate(t *testing.T) {
	tests := []struct {
		request AddWordRequest
		wantErr bool
		errMsg  string
	}{
		{AddWordRequest{Title: ""}, true, "input word is empty"},
		{AddWordRequest{Title: "validword"}, false, ""},
	}

	for _, tt := range tests {
		errMsg, isInvalid := tt.request.Validate()
		if isInvalid != tt.wantErr || errMsg != tt.errMsg {
			t.Errorf("Validate() got = %v, %v, want %v, %v", errMsg, isInvalid, tt.errMsg, tt.wantErr)
		}
	}
}

func TestInputWordValidate(t *testing.T) {
	tests := []struct {
		word    InputWord
		wantErr bool
		errMsg  string
	}{
		{"a", true, "input word is too long or too short, maximum 45 and minimum 2 characters. current length: 1"},
		{"toolongwordthatwillcausevalidationtofailbecauseitisveryverylong", true, "input word is too long or too short, maximum 45 and minimum 2 characters. current length: 63"},
		{"validword", false, ""},
		{"invalid1word", true, "input word should contain only alphabets. invalid character: 49"},
	}

	for _, tt := range tests {
		errMsg, isInvalid := tt.word.Validate()
		if isInvalid != tt.wantErr || errMsg != tt.errMsg {
			t.Errorf("Validate() got = %v, %v, want %v, %v", errMsg, isInvalid, tt.errMsg, tt.wantErr)
		}
	}
}

func TestInputWordPreProcess(t *testing.T) {
	tests := []struct {
		word     InputWord
		expected InputWord
	}{
		{"  Hello  ", "hello"},
		{"HeLLo WOrld", "hello world"},
		{"   leading and trailing spaces   ", "leading and trailing spaces"},
	}

	for _, tt := range tests {
		tt.word.PreProcess()
		if tt.word != tt.expected {
			t.Errorf("PreProcess() got = %v, want %v", tt.word, tt.expected)
		}
	}
}

func TestInputWordString(t *testing.T) {
	word := InputWord("testword")
	expected := "testword"
	if word.String() != expected {
		t.Errorf("String() got = %v, want %v", word.String(), expected)
	}
}