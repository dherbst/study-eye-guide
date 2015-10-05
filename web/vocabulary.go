package studyeyeguide

import (
	"appengine"
	// "appengine/datastore"
	"appengine/user"
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/api/1/vocab/save", VocabSaveHandler)
	http.HandleFunc("/api/1/vocab/test", VocabTestHandler)
}

type VocabWord struct {
	Test     string `json:"test"`
	Word     string `json:"word"`
	Sentence string `json:"sentence"`
}

func NewVocabWord(test, word, sentence string) *VocabWord {
	value := VocabWord{Test: test, Word: word, Sentence: sentence}
	return &value
}

func VocabSaveHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	u := user.Current(c)
	if u == nil {
		http.Error(w, "Authorization required", http.StatusUnauthorized)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Must be post", http.StatusBadRequest)
		return
	}

	word := NewVocabWord("", "", "")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(word)
	if err != nil {
		c.Errorf("Error decoding word %v", err)
		http.Error(w, "Error decoding word", http.StatusBadRequest)
		return
	}

}

func VocabTestHandler(w http.ResponseWriter, r *http.Request) {

}
