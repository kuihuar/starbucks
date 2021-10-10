package suanfa


//def bfs(graph, start, end):
//	queue = []
//	quene.append([start])//先入先出
//	visited.add(start)
//	while queue:
//		node = queue.popleft()
//		visited.add(node)
//		process(node)
//		nodes = generate_related_nodes(node)
//		queue.push(nodes)s

// bfs search template
//visited = set()
//def dfs(node,visited):
//	visited.add(node)
//	#process current node here.
//	for next_node in node.children()
//     if not next_node in visited:
//     	dfs(next_node, visited) // 不等当前层走完，就到下一层


// 递归写法

//visited = set()
//
//def dfs(node, visited):
//	if node in visited: //terminator
//		//already visited
//		return
//  visited.add(node)
//  // process current node here
//  for next_node in node.children():
//  	if not next_node in visited:
//  		dfs(next_node, visited)

// 非递归写法

//def dfs (self, tree):
//	if tree.root is None:
//		 return
//  visited, stack = [], [tree.root]
//  while stack:
//  	node = stack.pop()
//  	visited.add(node)
//  	process(node)
//  	nodes = renerate_related_node(node)
//  	stack.push(nodes)


// 动态规划 dynamic programming（动态递推）
//将一个复杂的问题分解成简单的子问题（用一种递归的方式）
//divide & conquer + Optional substructure
//分治 + 最优结构
//// 一般动态规划求一个最优解，最大值，最少的方式，不需要每一步都保存中间结果（所有的状态）
//// 只需要保存最优的状态。最后的放要推导出全局的最优的值
//// 状态的存储，
//动态规划和递归或者分冶没有根本区别, 关键看有无最优的子结构
//共性: 找到重复子问题
//差异: 最优子结构, 中途可以淘汰次优解



//如果没有最优子结构，那说明所有的 子问题都需要计算一遍，同时把最后结果合并在一起,
//所以传统意义上称之为分冶，每次的最优解就是当前解。它没有每次比较和淘汰的一个过程,
//动态规划有最优子结构，中途可以淘汰次优解，也必须淘汰，否则复杂。
// 复杂度从指数级降到了多项式

//dp 动态规划，一般是走接至底向上递推
//二维路径
//opt[i,j] = opt[i+1,j] + opt[j,j+1] //动态转移方程(有时候累加，有时候取最大值或最小值)
//	if a[i,j] == '空地':
//		opt[i,j] = opt[i+1,j] + opt[j,j+1]
//	else:
//		opt[i,j]=0
//分冶的话，要放在函数参数里。
//动态的话，要新建一个数组


//
//def recursion(level,param1,param2):
//	//recursion terminator
//	if level > MAX_LEVEL
//		process_result
//		return
//	//process logic in current level
//	process(level,data...)
//
//	// drill down
//	self.recursion(level+1,p1,...)
//	// reverse the current lvevel status if needed
//
//
//def devide_conquer(problem, param1,param2,...):
//	//recursion terminator
//	if problem is None:
//		print_result
//		return
//
//	//prepare data
//	data = prepare_data(problem)
//	subprolems = split_problem(problem, data)
//
//	//conquer subproblems
//	//每个子问题递归调用求解得了所谓子结果
//	subresult1 = self.divide_conquer(subproblems[0],p1,...)
//	subresult2 = self.divide_conquer(subproblems[1],p1,...)
//	subresult3 = self.divide_conquer(subproblems[3],p1,...)
//
//	//最后结果再合在一起所谓的本层的汇总结果
//	//process and generate the final result
//	result  = process_result(subresult1,subresult2,subresult3)
//
//	// revert rhe current level states
//
//
//visited = set()
//def dfs (node, visited):
//	// terminator
//	if node int visited:
//		//already visited
//		return
//  visited.add(node)
//
//  // process current node here
//  ...
//  for next_node in node.children():
//  	if next_node not in visited:
//  		 dfs(next_node, visited)
//
//def DFS(self, tree):
//
//	if tree.root is None:
//		return []
//  visited, stack = [],[tree.root]
//
//
//  while stack:
//  	node = stack.pop()
//  	visited.add(node)
//
//  	process(node)
//  	nodes = generate_related_nodes(node)
//  	stack.push(nodes)
//
//  // other processing work
//
//def BFS(graph, start, end):
//	visited = set()
//	queue = []
//	queue.append([start])
//
//	while queue:
//		node = queue.pop()
//		visited.add(node)
//
//		process(node)
//		nodes = generate_related_nodes(node)
//		queue.push(nodes)
//
//	// other processing work
//
//class Trie(object):
//	def __init__(self):
//		self.root= {}
//		self.end_of_word = '#'
//
//	def insert(self,word):
//		node = self.root
//		for char in word:
//			node = node.setdefault(char,{})
//		node[self.end_of_word] = self.end_of_word
//
//	def search(self,word):
