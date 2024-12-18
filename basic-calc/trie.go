package main

type TrieNode struct {
	chars map[int32]*TrieNode
	isEnd bool
}

type Trie = TrieNode

func NewTrie() *Trie {
	node := NewTrieNode()
	return node
}

func NewTrieNode() *TrieNode {
	t := TrieNode{}
	t.chars = map[int32]*TrieNode{}
	return &t
}

func (t *Trie) Insert(word string) {
	curr := t
	for _, ch := range word {
		_, okay := curr.chars[ch]
		if !okay {
			curr.chars[ch] = NewTrieNode()
		}
		curr = curr.chars[ch]
	}
	curr.isEnd = true
}

func (t *Trie) Has(word string) bool {
	curr := t
	for _, ch := range word {
		_, okay := curr.chars[ch]
		if !okay {
			return false
		}
		curr = curr.chars[ch]
	}
	return true
}
