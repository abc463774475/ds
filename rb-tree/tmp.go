package rb_tree

import "github.com/abc463774475/bbTool/n_log"

/*
交换父节点  以自己为中心
*/
func (self *NodeInfo)getRightMinNode(count int) *NodeInfo  {
	if self.C[0] == nil && self.C[1] == nil {
		return self
	}
	if count == 0 {
		if self.C[1] != nil {
			return self.C[1].getRightMinNode(1)
		}
		return self.C[0]
	}
	if self.C[0] == nil {
		return self
	}



	ret  :=  self.C[0].getRightMinNode(1)
	return ret
}


// 这里交换的值不对
func (self *NodeInfo) getExchangeNode() *NodeInfo  {
	if self.C[1] == nil && self.C[0] == nil{
		return self
	}

	if self.C[1] == nil {
		self.Key,self.C[0].Key = self.C[0].Key,self.Key
		return self.C[0]
	}

	p1 := self
	var p2 *NodeInfo

	for {
		count := 0
		p2 = p1.getRightMinNode(count)
		count++
//		n_log.Info("p1  %v  %v  p2  %v  %v",p1.IsRed,p1.Key,p2.IsRed,p2.Key)

		if p2 == p1 {
			break
		}
		p2.Key,p1.Key = p1.Key,p2.Key
		p1 = p2
	}

	return p1
}


func (self *NodeInfo) Remove_check(k int) *NodeInfo {
	pNode := self.Search(k)
	if pNode == nil {
		n_log.Panic("node not exist %v",k)
		return nil
	}
	//ClearChangeNode()
	//changeNode = pNode
	pChangeNode := pNode.getExchangeNode()
	//FreshTotalChange()
	{
//		n_log.Info("555555555")
		//g_tree.Traverser()

//		n_log.Info("66666666666")
	}
	/*
	 交换两者的值
	 */
	//pChangeNode.Key, pNode.Key = pNode.Key, pChangeNode.Key


	return pChangeNode
}

func (self *NodeInfo)Remove(k int)  {
	pNode := self.Remove_check(k)
	if pNode == nil {
		n_log.Info("node not exist  %v",k)
		return
	}
	if pNode == g_tree.Root {
		g_tree.Root = nil
		return
	}

	//g_tree.Traverser()

	pNode.Del_h(k)

	pNode.Del_father()
}

// 把自己的位置换给 其他节点
func (self *NodeInfo) ChangeOnePosToOther(pNode *NodeInfo)  {
	pNode.P = self.P
	if self == g_tree.Root {
		g_tree.Root = pNode
	}
	if self.P == nil {
		return
	}

	if self == self.P.C[0] {
		self.P.C[0] = pNode
	}else {
		self.P.C[1] = pNode
	}
}
// 增加一个位置的节点
func (self *NodeInfo) AddOnePos(pos int,pNode *NodeInfo)  {
	self.C[pos] = pNode
	if pNode != nil {
		pNode.P = self
	}
}