package bfs

import "fmt"

//BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordsSet := make(map[string]bool)
	for _,word:= range wordList{
		wordsSet[word]= true
	}
	var queue []string
	queue = append(queue,beginWord)
	visited := make(map[string]bool)
	visited[beginWord] = true
	step := 1
	for len(queue) > 0{
		currQueueSize := len(queue)
		//扩散一次
		for i:=0; i< currQueueSize;i++{
			word := queue[0]
			queue = queue[1:]
			for j:=0; j< len(word);j++{
				//originC := word[j]
				//枚举每一个字符
				for c:='a';c<='z';c++{
					//不在检查范围
					//if byte(c) == word[j] {
					//	continue
					//}
					newWord := word[:j]+ string(c) + word[j+1:]
					if !wordsSet[newWord] {
						continue
					}
					if newWord == endWord {
						return  step + 1
					}
					if !visited[newWord] {
						queue = append(queue,newWord)
						visited[newWord] = true
					}
				}
			}
		}
		//每一次扩散累加1
		step++
	}

	return 0
}
// 双向BFS
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	wordsSet := make(map[string]bool)
	for _,word := range wordList {
		wordsSet[word] = true
	}
	visited := make(map[string]bool)

	beginVisited := make(map[string]bool)
	beginVisited[beginWord]=true

	endVisited := make(map[string]bool)
	endVisited[endWord] = true
	step := 1
	if !wordsSet[endWord] {
		return 0
	}
	for len(beginVisited) > 0 {
		if len(beginVisited) > len(endVisited){
			beginVisited,endVisited = endVisited,beginVisited
		}
		nextLevelVisited := make(map[string]bool)
		for word,_ := range beginVisited {
			for j:=0; j< len(word);j++{
				for c:='a';c<='z';c++{
					newWord := word[:j]+ string(c) + word[j+1:]
					if endVisited[newWord] {
						return  step + 1
					}
					//什么样的word可以进入下一次循环
					if wordsSet[newWord]  && !visited[newWord]{
						nextLevelVisited[newWord] = true
						visited[newWord] = true
					}
				}
			}
		}
		//从begin这一侧向外扩散了一层
		beginVisited = nextLevelVisited
		step++
	}
	return 0
}
func TestLadderLength()  {
	var beginWord string
	var endWord string
	var wordList []string
	beginWord = "hit"
	endWord = "cog"
	//wordList = []string{"hot","dot","dog","lot","log","cog"}
	//wordList = []string{"hot","dot","dog","lot","log"}
	wordList = []string{"hot","cog","dot","dog","hit","lot","log"}
	res := ladderLength1(beginWord, endWord,wordList)
	fmt.Println(res)

}