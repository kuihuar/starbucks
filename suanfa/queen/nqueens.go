package queen

//def solveNQueens(self, n):
//	if n < 1:
//		return []
//  self.result = []
//  self.cols = set()
//  self.pie = set()
//  self.na = set()
//  self.DFS(n,0,[])
//  return self._generate_result(n)
//  def DFS(self, n, row, cur_state):
//  	//recursion terminator
//  	if row >= n:
//  		self.result.append(cur_state)
//      return
//    for col in range(n):
//    	if col in self.cols or row + col in self.pie or row - col in self.na:
//    		continue
//
//      //update the flags
//      self.cols.add(col)
//  		self.pie.add(row+col)
//      self.na.add(row-col)
//      self.dfs(n, row+1, cur_state + [col])
//      self.cols.remove(col)
//      self.pie.remove(row+col)
//      self.na.remove(row-col)

