package b_tree_self

import (
	"time"
	"bbTool/n_log"
)

var (
	pTree *BTree
)

type BTree struct {
	Root *BTreeNodeInfo
}

func CreateBTree() *BTree  {
	p := &BTree{}
	p.Root = CreateBTreeNodeInfo(true)

	return p
}

func (self *BTree) Traverse()  {
	if self.Root != nil {
		self.Root.Traverser()
	}
}

func (self *BTree) Search(k int) int  {
	start := time.Now()
	if self.Root == nil {
		return -1
	}

	pInfo := self.Root.Search_noPrecis(k)

	n_log.Info("cur  %v total time  %v",pInfo.Keys[:pInfo.N],time.Now().Sub(start))
	return 0
}

func (self *BTree) Search_precis(k int) *BTreeNodeInfo {
	if self.Root == nil {
		return nil
	}

	return self.Root.Search_precis(k)
}


func (self *BTree) Del_checkAdjacentKeys(k int)  {	// 删除节点 第一步置换节点信息
	pNode := self.Search_precis(k)
	if pNode == nil {
		n_log.Info("k not exist %v",k)
		return
	}
	if pNode.IsLeaf {
		return
	}

	_,pos:= pNode.FindKeysPos(k)
	pAdg_left := pNode.C[pos].FindAdjacent_left()

	pNode.Keys[pos],pAdg_left.Keys[pAdg_left.N-1] =pAdg_left.Keys[pAdg_left.N-1],pNode.Keys[pos]
}

func (self *BTree) Remove(k int)  {
	pNode := self.Search_precis(k)
	if pNode == nil {
		n_log.Erro("k not exist %v",k)
		return
	}
	pLeaf := pNode
//	var pos int
	if pNode.IsLeaf {
//		_,pos = pNode.FindKeysPos(k)
		pLeaf = pNode
	}else {
		_,pos:= pNode.FindKeysPos(k)
		pAdg_left := pNode.C[pos].FindAdjacent_left()
		pNode.Keys[pos],pAdg_left.Keys[pAdg_left.N-1] =pAdg_left.Keys[pAdg_left.N-1],pNode.Keys[pos]
		pNode = pAdg_left
//		pos = pNode.N - 1

		pLeaf = pNode
	}


	// 叶节点  位置都有啦
	pLeaf.Del(k)

	pLeaf.Remove_isIn()
}

func (self *BTree)Insert(k int)  {
	self.Root.Insert(k,false)
}

