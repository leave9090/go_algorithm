package sort

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")
	trie.Insert("application")
	trie.Insert("banana")

	fmt.Println(trie.Search("apple"))       // true
	fmt.Println(trie.Search("app"))         // true
	fmt.Println(trie.Search("application")) // true
	fmt.Println(trie.Search("ban"))         // false
	fmt.Println(trie.StartsWith("app"))     // true
	fmt.Println(trie.StartsWith("ban"))     // true
}
