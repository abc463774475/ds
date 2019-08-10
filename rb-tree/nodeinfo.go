package rb_tree

import "bbTool/n_log"

type NodeInfo struct {
	IsRed bool			// 不是红色就是黑色
	C []*NodeInfo		// 左右两个节点

	Key int				// 关键字
	P *NodeInfo

	//IsLeaf bool			// 是否叶节点
}

func (self *NodeInfo) Del ()  {
	self.P = nil
	self.C = []*NodeInfo{}
}

func CreateNodeInfo(k int) *NodeInfo {
	p := &NodeInfo{}
	p.IsRed = true
	//p.IsLeaf = false
	p.C = make([]*NodeInfo,2)
	p.Key = k

	return p
}

func (self *NodeInfo)Search(k int) *NodeInfo{
	if self.Key == k {
		return self
	}
	if self.Key < k {
		if self.C[1] == nil {
			return nil
		}else {
			return self.C[1].Search(k)
		}
	}else {
		if self.C[0] == nil {
			return nil
		}else {
			return self.C[0].Search(k)
		}
	}
}

// 搜索他的  父节点
func (self *NodeInfo) Search_father(k int)*NodeInfo {
	if self.Key == k {
		return self.P
	}

	// 木有子节点 那么就是他
	if self.C[0] == nil && self.C[1] == nil {
		return self
	}
	if self.Key < k {
		if self.C[1] == nil {
			return self
		}
		return self.C[1].Search_father(k)
	}
	if self.C[0] == nil {
		return self
	}
	return self.C[0].Search_father(k)
}

// 遍历
func (self *NodeInfo)Traverser()  {
	strColor := ""
	if self.IsRed {
		strColor = "red"
	}else {
		strColor = "black"
	}
	n_log.Info("%v  key  %v ",strColor,self.Key)
	for k,v := range self.C {
		if v != nil {
			if v.P != self {
				n_log.Panic("parent erro %v is red  %v   parent  %v  %v",v.Key,v.IsRed,v.P.IsRed,v.P.Key)
			}
			if k == 0 && v.Key > self.Key {
				n_log.Panic("eeeeerr  %v  %v",self.Key,v.Key)
			}
			if k == 1 && v.Key < self.Key {
				n_log.Panic("eeeeerr  %v  %v",self.Key,v.Key)
			}
			v.Traverser()
		}
	}
}


// 找到插入点 自下而上的 递归  ok？  结构貌似不是这么搞的
/*

 */
func (self *NodeInfo)Insert(k int)  {
	//
	parent := self.Search_father(k)
//	n_log.Info("parent  %v",parent)
	pNew := CreateNodeInfo(k)
	pNew.P = parent
	//isLeft := false
	if k < parent.Key {
		parent.C[0] = pNew
	//	isLeft = true
	}else {
		parent.C[1] = pNew
	}


	/*
		结构建立完成，，，然后开始做其他事情
		1 根据父节点做事情 ok？
	*/

	/*
		父节点是黑节点 则直接插入进行
	*/
	pNew.CheckNode()
}

func (self *NodeInfo) SwitchParent(other *NodeInfo)  {
	other.P = self.P
	if self == g_tree.Root {
		g_tree.Root = other
	}
	defer func() {
		// 交换完成   充值 自己的父节点
		self.P = other
	}()

	if self.P == nil {
		return
	}
	if self.P.C[0] == self {
		self.P.C[0] = other
	}else {
		self.P.C[1] = other
	}

}

func (self *NodeInfo) CheckNode ()  {
	if self == g_tree.Root {
		self.IsRed = false
		return
	}
	parent := self.P
	pNew := self
	isLeft := false
	if parent.C[0] == self {
		isLeft = true
	}

	if parent.IsRed == false {
		return
	}

	/*
	祖父节点不存在
	 */
	pGrand := parent.P
	if pGrand == nil {
		return
	}
	isParentLeft := false
	var pUncle *NodeInfo
	if pGrand.C[0] == parent {
		isParentLeft = true
		pUncle = pGrand.C[1]
	}else {
		pUncle = pGrand.C[0]
	}

	if pUncle == nil || pUncle.IsRed == false {
		if isLeft == true && isParentLeft == true {
			pGrand.IsRed = true
			parent.IsRed = false
			pGrand.C[0] = parent.C[1]
			if parent.C[1] != nil {
				parent.C[1].P = pGrand
			}
			parent.C[1] = pGrand

			pGrand.SwitchParent(parent)

			//pGrand.P = parent
		}else if isLeft == false && isParentLeft == true {
			pNew.IsRed = false
			pGrand.IsRed = true
			parent.P = pNew
			parent.C[1] = nil
			pNew.C[0] = parent
			pNew.C[1] = pGrand

			pGrand.C[0] = nil

			pGrand.SwitchParent(pNew)
		}else if isLeft == true && isParentLeft == false {
			pNew.IsRed = false
			pGrand.IsRed = true
			parent.P = pNew
			parent.C[0] = nil
			pNew.C[0] = pGrand
			pNew.C[1] = parent

			pGrand.C[1] = nil

			pGrand.SwitchParent(pNew)
		}else {
			/*
				1 都是右边
			 */
			parent.IsRed = false
			pGrand.IsRed = true

			pGrand.C[1] = parent.C[0]
			if parent.C[0] != nil{
				parent.C[0].P = pGrand
			}
			parent.C[0] = pGrand

			pGrand.SwitchParent(parent)
		}

	}else if pUncle.IsRed {
		/*
		recursion  咋个递归？？？
		 */
		pUncle.IsRed = false
		parent.IsRed = false
		pGrand.IsRed = true

		pGrand.CheckNode()
	}else {
		n_log.Panic("can not enter here")
	}
}

func (self *NodeInfo) Del_father()  {
	if self.P == nil {
		n_log.Erro("father not exist %v",self.Key)
		return
	}
	for k,v := range self.P.C {
		if v== self {
			self.P.C[k] = nil
			self.Del()
			return
		}
	}

	n_log.Panic("father not exist %v",self.Key)
}

// 获取兄弟节点
func (self *NodeInfo) GetBrother() *NodeInfo  {
	if self.P == nil {
		return nil
	}

	if self == self.P.C[0] {
		return self.P.C[1]
	}


	return self.P.C[0]
}

func (self *NodeInfo) Del_h(k int)  {
	if self == g_tree.Root {
		return
	}
	if self.IsRed {
		// 结束啦
		if self.Key == k {
			self.Del_father()
		}else {
			self.IsRed = false
			//pNode := self.Search(k)
			//pNode.Del_father()
		}

		return
	}

	selfIsLeft := false
	if self == self.P.C[0] {
		selfIsLeft = true
	}


	brother := self.GetBrother()
	if brother == nil {
		n_log.Panic("brother not exist  %v",self.Key)

		// 不是红黑树造成的 ，，，艹蛋疼。。。
	}
	p := self.P
	s := brother
	sl := s.C[0]
	sr := s.C[1]

	if selfIsLeft {
		if s.IsRed == false {
			// 先不管根节点  这种情况下 子树 可以实现自平衡
			if sr != nil && sr.IsRed {
				sr.IsRed = false
				s.IsRed = p.IsRed
				p.IsRed = false
				p.ChangeOnePosToOther(s)
				s.AddOnePos(0,p)
				p.AddOnePos(1,sl)

				return
			}else if sl != nil && sl.IsRed {
				sl.IsRed = p.IsRed
				p.ChangeOnePosToOther(sl)

				p.IsRed = false
				p.AddOnePos(1,sl.C[0])
				s.AddOnePos(0,sl.C[1])
				sl.AddOnePos(0,p)
				sl.AddOnePos(1,s)

			}else { // 全黑 或者一个不存在
				s.IsRed = true
				p.Del_h(k)

				return
			}
		}else { // 红节点

			p.ChangeOnePosToOther(s)

			s.AddOnePos(0,p)

			s.IsRed = false
			if sl != nil{
				sl.IsRed = true
			}
			p.AddOnePos(1,sl)
		}
	}else {
		if s.IsRed == true {
			p.ChangeOnePosToOther(s)
			s.IsRed = false
			s.AddOnePos(1,p)
			if sr != nil{
				sr.IsRed = true
			}
			p.AddOnePos(0,sr)
		}else {
			if sl != nil && sl.IsRed {
				sl.IsRed = false
				s.IsRed = p.IsRed
				p.IsRed = false
				p.ChangeOnePosToOther(s)

				s.AddOnePos(1,p)
				p.AddOnePos(0,sr)
			}else if sr != nil && sr.IsRed {
				sr.IsRed = p.IsRed
				p.ChangeOnePosToOther(sr)

				p.IsRed = false
				s.AddOnePos(1,sr.C[0])
				p.AddOnePos(0,sr.C[1])
				sr.AddOnePos(0,s)
				sr.AddOnePos(1,p)

			}else {
				s.IsRed = true
				p.Del_h(k)

				if k == 8 {
					n_log.Info("111111  s %v  %v self %v %v",s.IsRed,s.Key)
				}

				return
			}
		}
	}
}

/*
与层次无关
 */

var (
	totalChangeNode []*NodeInfo = []*NodeInfo{}
	changeNode *NodeInfo
)

func ClearChangeNode()  {
	totalChangeNode = []*NodeInfo{}
	changeNode = nil

}
func PrintTotalChangeNode()  {
	n_log.Info("total changeNode  %v",len(totalChangeNode))
	for _,v := range totalChangeNode {
		n_log.Info("v .. isRed  %v  key  %v",v.IsRed,v.Key)
	}
	n_log.Info("total changeNode end")
}

func FreshTotalChange()  {
	pChangeNode:= changeNode
	if len(totalChangeNode) == 0 {
		return
	}
	k1 := totalChangeNode[0].Key
	for i := len(totalChangeNode) -1 ; i >0 ; i--{
		totalChangeNode[i-1].Key = totalChangeNode[i].Key
	}

	totalChangeNode[len(totalChangeNode) -1].Key = pChangeNode.Key
	pChangeNode.Key = k1
}

func (self *NodeInfo) GetChangeNode_h(cengci int) *NodeInfo  {
	// 获取自己的节点
	if self.C[0] == nil && self.C[1] == nil {
		if self != changeNode{
			totalChangeNode = append(totalChangeNode, self)
		}
		return self
	} else if self.C[0] == nil {
		// 只有一个节点
		if self != changeNode {
			totalChangeNode = append(totalChangeNode, self)
		}

		p := self.C[1].GetChangeNode_h(1)
		return p
	}else if self.C[1] == nil {
		if self != changeNode {
			totalChangeNode = append(totalChangeNode, self)
		}
		p := self.C[0].GetChangeNode_h(1)
		return p
	}
	if cengci == 0 {
		return self.C[1].GetChangeNode_h(1)
	}

	return self.C[0].GetChangeNode_h(1)
}