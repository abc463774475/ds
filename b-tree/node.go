package b_tree

import "bbTool/n_log"

type BTreeNode struct {
	Keys []int		// 数组
	T int			// degree

	C []*BTreeNode
	N int

	IsLeaf bool	// 是否是叶节点 ？
}

func CreateBTreeNode(_t1 int,_isLeaf bool) *BTreeNode  {
	p := &BTreeNode{}
	p.T = _t1
	p.IsLeaf = _isLeaf
	p.Keys = make([]int,2*_t1-1)
	p.C = make([]*BTreeNode,2*_t1)
	p.N = 0

	return p
}

func (self *BTreeNode) Traverse()  {
	n_log.Info("n  %v  keys %#v",self.N,self.Keys[:self.N])
	i := 0
	for ;i < self.N;i++ {
		if self.IsLeaf == false{
			//n_log.Info("keys %#v",self.Keys)
			self.C[i].Traverse()
		}
	}

	if self.IsLeaf == false {
		self.C[i].Traverse()
	}
}

func (self *BTreeNode)FindKey(k int) int {
	idx := 0
	for {
		if idx < self.N && self.Keys[idx] < k {
			idx++
		}else {
			break
		}
	}
	return idx
}

func (self *BTreeNode) Remove(k int)  {
	// todo
	idx := self.FindKey(k)
	if idx < self.N && self.Keys[idx] == k {
		if self.IsLeaf {
			self.RemoveFromLeaf(idx)
		}else {
			self.RemoveFromNonLeaf(idx)
		}
	}else {
		if self.IsLeaf {
			n_log.Info("keys %v  not find in node",k)
			return
		}
		flag := false
		if idx == self.N {
			flag = true
		}
		if self.C[idx].N < self.T {
			self.Fill(idx)
		}

		if flag && idx > self.N {
			self.C[idx-1].Remove(k)
		}else {
			self.C[idx].Remove(k)
		}

	}

	return
}

func (self *BTreeNode) RemoveFromLeaf(idx int)  {
	for i:=idx+1;i < self.N;i++{
		self.Keys[i-1] = self.Keys[i]
	}

	self.N--
	return
}

func (self *BTreeNode) RemoveFromNonLeaf(idx int)  {
	// todo
	k := self.Keys[idx]
	if self.C[idx].N >= self.T {
		pred := self.GetPred(idx)
		self.Keys[idx] = pred
		self.C[idx].Remove(pred)
	}else if self.C[idx+1].N >= self.T {
		succ := self.GetSucc(idx)
		self.Keys[idx] = succ
		self.C[idx+1].Remove(succ)
	}else {
		self.Merge(idx)
		self.C[idx].Remove(k)
	}
	return
}

func (self *BTreeNode)GetPred(idx int) int {
	cur := self.C[idx]
	for {
		if !cur.IsLeaf {
			cur = cur.C[cur.N]
		}else {
			break
		}
	}

	return cur.Keys[cur.N-1]
}

func (self *BTreeNode) GetSucc(idx int) int {
	cur := self.C[idx+1]

	for {
		if !cur.IsLeaf {
			cur = cur.C[0]
		}else {
			break
		}
	}

	return cur.Keys[0]
}

func (self *BTreeNode) Fill(idx int)  {
	// todo
	if idx != 0 && self.C[idx-1].N >= self.T {
		self.BorrowFromPrev(idx)
	}else if idx!=self.N && self.C[idx+1].N >= self.T {
		self.BorrowFromNext(idx)
	}else {
		if idx != self.N {
			self.Merge(idx)
		}else {
			self.Merge(idx-1)
		}
	}
}

func (self *BTreeNode)BorrowFromPrev(idx int)  {
	child := self.C[idx]
	sibling := self.C[idx -1]
	for i:= child.N -1; i >=0;i--{
		child.Keys[i+1] = child.Keys[i]
	}

	if !child.IsLeaf {
		for i:= child.N ;i >=0;i--{
			child.C[i+1] = child.C[i]
		}
	}

	child.Keys[0] = self.Keys[idx - 1]
	if !child.IsLeaf {
		child.C[0] = sibling.C[sibling.N]
	}
	self.Keys[idx-1] = sibling.Keys[sibling.N-1]
	child.N++
	sibling.N--
	return
}

func (self *BTreeNode)BorrowFromNext(idx int)  {
	child := self.C[idx]
	sibling := self.C[idx+1]

	child.Keys[child.N] = self.Keys[idx]
	if !child.IsLeaf	{
		child.C[child.N +1] = sibling.C[0]
	}
	self.Keys[idx] = sibling.Keys[0]

	for i:=1; i <sibling.N;i++{
		sibling.Keys[i-1] = sibling.Keys[i]
	}

	if !self.IsLeaf {
		for i:=1; i <sibling.N;i++{
			sibling.C[i-1] = sibling.C[i]
		}
	}

	child.N++
	sibling.N--
	return
}

func (self *BTreeNode)Merge(idx int)  {
	child := self.C[idx]
	sibling := self.C[idx+1]

	child.Keys[self.T-1] = self.Keys[idx]
	for i:=0; i< sibling.N;i++{
		child.Keys[i+self.T] = sibling.Keys[i]
	}

	if !child.IsLeaf {
		for i:=0; i<sibling.N;i++{
			child.C[i+self.T] = sibling.C[i]
		}
	}

	for i:=idx+1; i <self.N;i++{
		self.Keys[i-1]= self.Keys[i]
	}
	for i:=idx+2; i <self.N;i++{
		self.C[i-1] = self.C[i]
	}

	child.N += sibling.N+1
	self.N--
}

func (self *BTreeNode) Search(k int)  *BTreeNode{
	i := 0
	for {
		if i < self.N && k > self.Keys[i] {
			i++;
		}else {
			break
		}
	}
	if self.Keys[i] == k {
		return self
	}
	if self.IsLeaf == true {
		return nil
	}
	return self.C[i].Search(k)
}

func (self *BTreeNode) InsertNonFull(k int)  {
	i := self.N-1
	if self.IsLeaf {
		for {
			if i >= 0 && self.Keys[i] > k {
				self.Keys[i+1] = self.Keys[i]
				i--
			}else {
				break
			}
		}
		self.Keys[i+1] = k
		self.N++
	}else {
		for {
			if i >= 0 && self.Keys[i] > k {
				self.Keys[i+1] = self.Keys[i]
				i--
			}else {
				break
			}
		}

		if self.C[i+1].N == 2 *self.T -1 {
			//self.sp
			self.spiltChild(i+1, self.C[i+1])

			if self.Keys[i+1] < k {
				i++
			}
		}

		self.C[i+1].InsertNonFull(k)
	}
}

func (self *BTreeNode) spiltChild(i int, y *BTreeNode)  {
	z := CreateBTreeNode(y.T,y.IsLeaf)
	z.N = self.T - 1

	for j:=0; j < self.T-1; j++{
		z.Keys[j] = y.Keys[j+self.T]
	}

	if y.IsLeaf == false {
		for j:=0; j<self.T;j++{
			z.C[j] = y.C[j+self.T]
		}
	}

	y.N = self.T -1
	for j:= self.N; j >=i+1;j--{
		self.C[j+1] = self.C[j]
	}
	self.C[i+1] = z

	for j:= self.N -1;j>= i;j--{
		self.Keys[j+1] = self.Keys[j]
	}
	self.Keys[i] = y.Keys[self.T-1]

	self.N++
}


