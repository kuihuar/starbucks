package trie

import "fmt"

// all words->Trie
// board DFS
func findWords(board [][]byte, words []string) []string {
	var res []string
	root := Constructor()
	for _, word := range words {
		root.Insert(word)
	}
	//fmt.Println(root.Search("oaa"))
	//fmt.Println(root.Search("oa"))
	//fmt.Println("root.Search(oa)")
	m := len(board)
	n := len(board[0])
	visited := make([][]bool,m)
	for i:=0; i<m;i++{
		visited[i] = make([]bool,n)
	}
	word := []byte{}
	for i := 0 ;i<m;i++{
		for j:= 0;j<n;j++{
			dfs(board,visited,i,j,&word,&res,&root)
		}
	}
	return  res
}
func dfs(board [][]byte,visited [][]bool, i, j int,word *[]byte, words *[]string, trie *Trie)  {
	if i<0 || j < 0 || i >=len(board) || j >= len(board[0]) || visited[i][j] {
		return
	}
	char := board[i][j]
	if trie.children[char-'a'] ==nil {
		return
	}
	visited[i][j] = true
	*word = append(*word,char)
	node  := trie.children[char-'a']
	if node.isEnd {//预防再次插入，防止结果重复
		*words = append(*words,string(*word))
		node.isEnd = false
	}

	//四个方向扩散
	directions := [][]int{{-1,0},{1,0},{0,-1},{0,1}}
	for _, direction := range directions{
		x, y := i+direction[0],j+direction[1]
		dfs(board,visited, x, y, word, words, node)
	}
	//回溯
	visited[i][j] = false
	*word = (*word)[:len(*word)-1]
}

func TestFindWords2()  {
	//words :=[]string {"oath","pea","eat","rain"}
	//board := [][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}}
	//words := []string{"oa","oaa"}
	//board := [][]byte{{'o','a','b','n'},{'o','t','a','e'},{'a','h','k','r'},{'a','f','l','v'}}
	words := []string{"aaa"}
	board := [][]byte{{'a'},{'a'}}
	res := findWords(board, words)
	fmt.Println("res=",res)
}