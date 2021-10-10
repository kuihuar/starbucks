package search

import (
	"container/heap"
	"fmt"
)

//1. BFS
//2. A*(启发式搜索)

func shortestPathBinaryMatrix(grid [][]int) int {
	var queue [][3]int
	queue = append(queue,[3]int{0,0,2})//起点，终点，步数
	n := len(grid)

	if grid[0][0] ==1 || grid[n-1][n-1] ==1{
		return -1
	}
	if n<=2 {
		return n
	}
	for len(queue)>0{
		node := queue[0]
		queue = queue[1:]
		i,j,d:= node[0],node[1],node[2]
		directions:=[][]int{{i-1,j-1},{j-1,j},{i-1,j+1},{i,j-1},{i,j+1},{i+1,j-1},{i+1,j},{j+1,j+1}}
		for _,direction := range directions{
			x,y:=direction[0],direction[1]
			if x>=0 && x < n && y>=0 && y<n && grid[x][y] == 0 {
				if x == n-1 && y== n-1 {
					return d +1
				}
				queue=append(queue,[3]int{x,y,d+1})
			}
		}
	}
	return -1
}
type cell [2]int
func shortestPathBinaryMatrixWithAstar(grid [][]int) int {
	// 与当前单元格相邻的单元格，八连通相连的node(i,j附近的8个点)
	var successorFun func(cell) []cell
	successorFun = func(c cell) (cells []cell){
		i,j := c[0],c[1]
		directions:=[][]int{{i-1,j-1},{i-1,j},{i-1,j+1},{i,j-1},{i,j+1},{i+1,j-1},{i+1,j},{i+1,j+1}}
		for _,direction := range directions{
			x,y:=direction[0],direction[1]
			if x>=0 && x < len(grid) && y>=0 && y<len(grid[0]) && grid[x][y] == 0 {
				cells=append(cells,cell{x,y})

			}
		}
		return cells
	}
	//启发式函数
	var heuristic func(c1 cell, c2 cell) int
	heuristic = func(c1 cell,c2 cell) int{
		disX := float64(c2[0]) - float64(c1[0])
		disY := float64(c2[1]) - float64(c1[1])
		if disX > disY {
			return int(disX)
		}else{
			return int(disY)
		}
	}
	var reconstructPath func(map[cell]cell,  cell, cell) []cell
	reconstructPath = func(cameFrom map[cell]cell, start cell, end cell) []cell {
		fmt.Printf("cameFrom=%v, start=%v,end=%v\n",cameFrom,start,end)
		var revertpath = []cell{}
		revertpath = append(revertpath,end)
		fmt.Printf("start=%v\n", start)
		fmt.Printf("end=%v\n", end)
		for end != start {
			end = cameFrom[end]
			fmt.Printf("end&&&&&=%v\n", end)
			revertpath = append(revertpath, end)
		}
		fmt.Printf("revertpath=%v\n",revertpath)
		for i,j:=0,len(revertpath)-1;i<j;i,j=i+1,j-1{
			fmt.Printf("%d<%d\n", i,j)
			revertpath[i],revertpath[j] = revertpath[j],revertpath[i]
		}
		fmt.Printf("revertpath=%v\n",revertpath)
		return revertpath
	}
	var aStarGraphSearch func (grid [][]int) []cell
	aStarGraphSearch = func(grid [][]int) []cell {
		visited:= make(map[cell]bool)
		cameFrom := make(map[cell]cell)
		distance := make(map[cell]int)
		startCell := cell{0,0}
		distance[startCell] = 0
		var pq PriorityQueue
		heap.Init(&pq)
		heap.Push(&pq, &Item{
			value:    startCell,
			priority: 0,
		})

		goalCell:= cell{len(grid)-1,len(grid[0])-1}
		for len(pq) > 0{
			//cell := pq.Pop().(*Item).value
			cell := pq[0].value
			//fmt.Printf("%+v\n",pq[0])
			heap.Pop(&pq)
			if visited[cell] {
				continue
			}
			fmt.Printf("%v == %v\n",goalCell,cell)
			if goalCell == cell {
				return  reconstructPath(cameFrom, startCell, cell)
			}
			visited[cell] = true
			for _, successorCell := range successorFun(cell){
				heap.Push(&pq, &Item{
					value:    successorCell,
					priority: distance[cell]+1 + heuristic(successorCell,goalCell),
				})
				if _,ok := distance[successorCell];!ok{
					if distance[cell] +1 > distance[successorCell]{
						distance[successorCell] = distance[cell] +1
						cameFrom[successorCell] = cell
					}
					fmt.Printf("distance=%v\n",distance)
				}
			}
		}

		return nil
	}
	shortedPath := aStarGraphSearch(grid)
	fmt.Printf("shortedPath:%v\n",shortedPath)
	if shortedPath == nil  || grid[0][0]==1{
		return -1
	}
	return len(shortedPath)
}

type Item struct {
	value cell
	priority int
	index int
}
type PriorityQueue []*Item

func (pq PriorityQueue) Len()int {
	return len(pq)
}
func (pq PriorityQueue) Less(i,j int) bool{
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue)Swap(i, j int)  {
	pq[i],pq[j] = pq[j], pq[i]
	pq[i].index =i
	pq[j].index =j
}
func (pq *PriorityQueue) Push(x interface{})  {
	n := len(*pq)
	item := x.(*Item)
	item.index=n
	*pq = append(*pq, item)

}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n:= len(old)
	item := old[n-1]
	item.index=-1
	*pq = old[:n-1]
	return item
}
func (pq *PriorityQueue)Update(item *Item, value cell, priority int)  {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func TestSPBM()  {
	//grid := [][]int{{0,0,0},{1,1,0},{1,1,0}}
	grid := [][]int{{0,0,0},{1,1,0},{1,1,0}}
	//res := shortestPathBinaryMatrix(grid)
	res := shortestPathBinaryMatrixWithAstar(grid)
	fmt.Println(res)
}









type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}