package search

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