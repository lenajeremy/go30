package main

import (
	"fmt"
	"testing"
)

func TestTrie_Has(t *testing.T) {
	trie := NewTrie()

	hasWord := trie.Has("jeremiah")
	fmt.Println(hasWord)
	if hasWord {
		t.Errorf("")
	}
}

func TestTrie_Insert(t *testing.T) {
	trie := NewTrie()
	words := map[string]bool{
		"jeremiah": true, "jeremiah lena": true, "jerry": true,
	}

	for w, _ := range words {
		trie.Insert(w)
	}

	searchWords := map[string]bool{
		"jeremiah": true, "jeremiah lena": true, "jerry": true,
		"johnson": false, "jerry gana": false, "jeremiah lena osas": false,
	}

	for w, inTrie := range searchWords {
		exists := trie.Has(w)
		if exists != inTrie {
			t.Errorf("Error: %s exists: %v, result: %v", w, inTrie, exists)
		}
	}
}
