package trie_test

import (
	. "ds/trie"
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()
	trie.Add("abcdeq", 100)
	trie.Add("eqcdiw", 200)
	trie.Add("wtcdwc", 300)
	trie.Add("aqcdut", "Hello")
	trie.Add("ebcdwz", true)
	trie.Add("abkdyq", 1.2)
	trie.Add("obcnzp", 9)
	trie.Traverse(func(e interface{}) {
		fmt.Println(e)
	})

	fmt.Println(trie.Get("abkdyq"))
	fmt.Println(trie.Get("abkdxy"))
	if v, _ := trie.Get("abkdyq"); v != 1.2 {
		t.Error("get error")
	}
	if v, _ := trie.Get("eqcdiw"); v != 200 {
		t.Error("get error")
	}

}
