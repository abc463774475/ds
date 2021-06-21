package rb_tree

import "github.com/abc463774475/bbtool/n_log"

/*

 */
var (
	g_tree *TreeInfo = CreateTreeInfo()
)

/*
 */
type TreeInfo struct {
	Root *NodeInfo
}

func CreateTreeInfo()*TreeInfo  {
	p := &TreeInfo{}

	return p
}

func (self *TreeInfo)Search(k int)*NodeInfo  {
	if self.Root == nil {
		return nil
	}
	return self.Root.Search(k)
}

func (self *TreeInfo)Insert(k int)  {
	if pNode := self.Search(k); pNode!= nil{
		n_log.Info("have insert one k  %v",k)
		return
	}
	if self.Root == nil {
		pNode := CreateNodeInfo(k)
		pNode.IsRed = false
		self.Root = pNode
	}else {
		self.Root.Insert(k)
	}
}

func (self *TreeInfo)Traverser()  {
	if self.Root == nil {
		n_log.Info("this is empty")
		return
	}
	self.Root.Traverser()
}

func (self *TreeInfo) Remove(k int)  {
	if self.Root == nil {
		return
	}
	self.Root.Remove(k)
}
