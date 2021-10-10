package disjointset

import "fmt"

//并查集

type UnionFind struct {
	//连通数
	count int
	parent []int
	//树高度
	rank []int
}
func newUnionFind(n int) *UnionFind{
	uf := &UnionFind{
		count:  n,
		parent: make([]int,n),
		rank: make([]int,n),
	}
	//自己指向自己
	for i:= 0; i<n;i++{
		uf.parent[i] = i
	}
	return uf
}
func (uf UnionFind)find(p int) int {
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]] //爷爷节点
		p = uf.parent[p]
	}
	return p
}
func (uf *UnionFind)union(p, q int)  {
	rootP := uf.find(p)
	rootQ := uf.find(q)
	fmt.Printf("pNode=%d,qNode=%d\n",rootP,rootQ)
	if rootP == rootQ {
		return
	}
	if uf.rank[rootQ] > uf.rank[rootP]{
		uf.parent[rootP] = rootQ
	}else{
		uf.parent[rootQ] = rootP
		if uf.rank[rootP] == uf.rank[rootQ]{
			uf.rank[rootP]++
		}
	}
	uf.count--
}
func (uf *UnionFind) Count() int {
	return uf.count
}
func (uf *UnionFind)fineParent(i int) int {
	root := i
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	return root
}