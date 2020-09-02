package b_tree

import "github.com/abc463774475/bbTool/n_log"

type BTree struct {
	Root *BTreeNode
	T int		// degree
}

func CreateBTree(_t int) *BTree {
	p := &BTree{}
	p.Root = nil
	p.T = _t
	return p
}

func (self *BTree) Traverse()  {
	if self.Root != nil {
		self.Root.Traverse()
	}
}



func (self *BTree)Search(k int) *BTreeNode  {
	if self.Root != nil {
		return self.Root.Search(k)
	}
	return nil
}

func (self *BTree) Insert(k int)  {
	if self.Root == nil {
		self.Root = CreateBTreeNode(self.T,true)
		self.Root.Keys[0] = k
		self.Root.N++
		return
	}

	if self.Root.N == 2*self.T -1 {
		s := CreateBTreeNode(self.T,false)
		s.C[0] = self.Root
		s.spiltChild(0,self.Root)

		i := 0
		if s.Keys[0] <k {
			i++
		}
		s.C[i].InsertNonFull(k)
		self.Root = s
	}else {
		self.Root.InsertNonFull(k)
	}

}

func (self *BTree) Remove(k int)  {
	if self.Root == nil {
		n_log.Info("cur is empty")
		return
	}
	self.Root.Remove(k)
	if self.Root.N == 0 {
		if self.Root.IsLeaf {
			self.Root = nil
		}else {
			self.Root = self.Root.C[0]
		}
	}
	return
}
