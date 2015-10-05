package studyeyeguide

import (
	// "appengine/aetest"
	"testing"
)

func TestNewVocabWord(t *testing.T) {
	word := NewVocabWord("test1", "expedition", "An expedition is a journey made for a special purpose.")
	if word.Test != "test1" {
		t.Errorf("Did not get expected test name back, got %v", word.Test)
	}
	if word.Word != "expedition" {
		t.Errorf("Did not get expected word back, got %v", word.Word)
	}
	if word.Sentence != "An expedition is a journey made for a special purpose." {
		t.Errorf("Did not get expected sentence back, got %v", word.Sentence)
	}
}
