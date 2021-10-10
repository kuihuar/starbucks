package datastructures

import "fmt"

const (
	ALPHABET_SIZE = 26
)

type trieNode struct {
	children [ALPHABET_SIZE]*trieNode
	isWordEnd bool
}
type trie struct {
	root *trieNode
}
func initTrie() *trie {
	return &trie{
		root:&trieNode{},
	}
}

func (t *trie) insert(word string) {
	wordLength := len(word)
	current := t.root
	for i :=0; i< wordLength; i++ {
		index := word[i] - 'a'
		//fmt.Printf("word[i] = %c\n", word[i])
		//fmt.Printf("index = %d\n", index)
		//如果对应的index上有孩子节点，证明当前字符存在
		if current.children[index] == nil {
			current.children[index] = &trieNode{}
		}
		current = current.children[index]
	}
	current.isWordEnd = true

}
func (t *trie)find(word string) bool{
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}
	if current.isWordEnd {
		return true
	}
	return false
}
func TestTrie(){
	trie := initTrie()
	words := []string{"sam", "john", "trim", "jose", "rose", "cat","dog", "dogs", "roses"}
	for i:=0; i< len(words); i++ {
		trie.insert(words[i])
	}
	//wordsToFind := []string{"sam", "john", "trim", "Jose","san","ans", "cat","dog", "dogs", "roses"}
	wordsToFind := []string{"sam", "john", "trim", "jose","san","ans", "cat","dog", "dogs", "roses"}
	for i:=0; i<len(wordsToFind); i++{
		found := trie.find(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		}else{
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}
