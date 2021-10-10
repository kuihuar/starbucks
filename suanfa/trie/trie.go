package trie

import "fmt"

type Trie struct {
	children[26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}
func (this *Trie)Insert(word string)  {
	node := this
	for _, char := range word {
		char -= 'a'
		if node.children[char] == nil {
			node.children[char] = &Trie{}
		}
		node = node.children[char]
	}
	node.isEnd = true
}
func (this *Trie)Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}
func (this * Trie)searchPrefix(prefix string) *Trie {
	node := this
	for _, char := range prefix {
		char -= 'a'
		if node.children[char] == nil {
			return nil
		}
		node = node.children[char]
	}
	return node
}
func (this *Trie)StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func TestTrie()  {
	obj := Constructor()
	obj.Insert("word")
	res := obj.Search("word")
	res2 := obj.StartsWith("wor")
	fmt.Println(res,res2)
}
