package b_tree_self

import (
	"github.com/abc463774475/bbTool/n_log"
)

var (
	NodeKey_min = 2  // <=
	NodeKey_max = 4  // <=

	NodeNum = 0
)

type BTreeNodeInfo struct {
	N   int			// 关键字个数
	Keys []int		// 关键字
	C []*BTreeNodeInfo	// 孩子节点

	P  *BTreeNodeInfo	// 父节点  指针

	IsLeaf bool			// 是否是 叶节点
}

func CreateBTreeNodeInfo(_isLeaf bool) *BTreeNodeInfo {
	p := &BTreeNodeInfo{}
	p.IsLeaf =_isLeaf

	ln := NodeKey_max

	p.Keys = make([]int, ln+1)
	p.C = make([]*BTreeNodeInfo,ln+2)

	p.N = 0

	return p
}


func (self *BTreeNodeInfo)Traverser()  {
	NodeNum += self.N
	n_log.Info("this node info %v   %v  %p  isleaf  %v",self.N, self.Keys[:self.N],self,self.IsLeaf)
	for i:=0; i < self.N+1;i++{
		if self.IsLeaf == false {
			if self.C[i].P != self {
				self.Print()
				self.C[i].Print()

				self.C[i].P.Print()
				n_log.Panic("errrr  %p  ",self.C[i].P)
			}

			self.C[i].Traverser()		// 这儿木有问题哈 。。。。 这样遍历
		}
	}
}

func (self *BTreeNodeInfo)Print()  {
	n_log.Info("self isleaf %v  n %v  keys  %v",self.IsLeaf,self.N,self.Keys)
}

func (self *BTreeNodeInfo) FindKeysPos(k int) (isIn bool, pos int) {
	i := 0
	for ; i < self.N;i++{
		if self.Keys[i] == k {
			return true,i
		}
		if k < self.Keys[i] {
			return false,i
		}
	}

	return false,i
}

// 非精确查找 节点？？？
func (self *BTreeNodeInfo) Search_noPrecis(k int) *BTreeNodeInfo  {
	if self.IsLeaf {		// 到了这一层 说明 可以搜到啦
		return self
	}

	isIn,pos := self.FindKeysPos(k)

	if isIn {
		return self
	}
	return self.C[pos].Search_noPrecis(k)
}

func (self *BTreeNodeInfo) Search_precis(k int) (*BTreeNodeInfo)  {
	isIn,pos := self.FindKeysPos(k)

	if isIn {
		return self
	}
	if self.IsLeaf {
		return nil
	}

	return self.C[pos].Search_precis(k)
}


func (self *BTreeNodeInfo) Insert(k int,isSelfAdd bool)  {
	if self.N == 0 {		// 第一次插入
		self.Keys[0] = k
		self.N++
		return
	}

	var pNode *BTreeNodeInfo
	if !isSelfAdd{
		pNode = self.Search_noPrecis(k)
	}else {
		pNode = self
	}
	//n_log.Info("insert search  %p  keys %v",pNode,pNode.Keys[:pNode.N])
	isIn ,pos := pNode.FindKeysPos(k)

	//n_log.Info("is in  %v pos  %v",isIn,pos)
	if isIn {
		n_log.Info("have insert this node")
		return
	}

	for i := pNode.N; i >pos; i--{
		pNode.Keys[i] = pNode.Keys[i-1]
	}
	pNode.Keys[pos] = k
	pNode.N++

	pNode.Split()
}

func (self *BTreeNodeInfo) AddOneKey_noOrder(k int)  {
	self.Keys[self.N] = k
	self.N++
}

func (self *BTreeNodeInfo)FindAdjacent_left() *BTreeNodeInfo {
	if self.IsLeaf {
		return self
	}
	return self.C[self.N].FindAdjacent_left()
}

func (self *BTreeNodeInfo)FindAdjacent_right() *BTreeNodeInfo {
	if self.IsLeaf {
		return self
	}
	return self.C[0].FindAdjacent_right()
}

// 分裂
func (self *BTreeNodeInfo)Split()  {
	if self.N <= NodeKey_max {
		return
	}
	//n_log.Info("split  ok   %v  %p",self.Keys[:self.N],self.P)

	// 从中破开   ok？
	if self.P == nil {
		// 第一次分裂 创建两个节点
		pos_split := NodeKey_max/2

		c2 := CreateBTreeNodeInfo(self.IsLeaf)
		for i := pos_split + 1; i < self.N;i++{
			c2.AddOneKey_noOrder(self.Keys[i])
		}

		// 分裂之后  ，，要兼顾  childNode ok？？？  节点要全部移位
		if !self.IsLeaf {
			// self  c2 要触发 孩子节点分配
			//n_log.Info("cccccc  %v",self.N)
			pos := 0
			for i := pos_split+1; i <self.N+1;i++{
				c2.C[pos] = self.C[i]
				c2.C[pos].P = c2

				pos++
			}
		}


		self.N = pos_split

		parent := CreateBTreeNodeInfo(false)
		parent.Keys[0] = self.Keys[pos_split]
		parent.N++
		parent.C[0] = self
		parent.C[1] = c2

		self.P = parent
		c2.P = parent
		// 咋个改变跟节点？？？？
		pTree.Root = parent

		//pTree.Traverse()
		return
	}else {
		// 有父节点的情况
		pos_split := NodeKey_max/2

		c2 := CreateBTreeNodeInfo(self.IsLeaf)
		for i := pos_split + 1; i < self.N;i++{
			c2.AddOneKey_noOrder(self.Keys[i])
		}

		parent := self.P
		c2.P = parent

		if !self.IsLeaf {
			// self  c2 要触发 孩子节点分配
//			n_log.Info("cccccc  %v",self.N)
			pos := 0
			for i := pos_split+1; i <self.N+1;i++{
				c2.C[pos] = self.C[i]
				c2.C[pos].P = c2

				pos++
			}
		}

		self.N = pos_split

		pPos := 0
		for i,v := range parent.C {
			if v == self {
				pPos = i
				break
			}
		}

	//	n_log.Info("ppos  %v  parent  %v  %v  %v",pPos,parent.Keys,parent.N,parent.C)

		for i:= parent.N+1; i > pPos;i--{
			parent.C[i] = parent.C[i-1]
		}

		parent.C[pPos+1] = c2

		parent.Insert(self.Keys[pos_split],true)
	}
}

func (self *BTreeNodeInfo)Del(k int)  {
	_,pos := self.FindKeysPos(k)

	for i:= pos; i < self.N -1;i++{
		self.Keys[i] = self.Keys[i+1]
	}
	//n_log.Info("self  %v  ",self.Keys[:self.N])
	self.N--
}

func (self *BTreeNodeInfo)Remove_isIn()  {
	// 证明自己节点 满足最小条件 ，，所以无所谓
	//defer func() {
	//	pTree.Traverse()
	//}()

//	n_log.Info("remove isin  %v, %v",self.Keys,self.N)
	if self.N >= NodeKey_min {
		return
	}

	if self.P == nil {
		return
	}

	parent := self.P
	selfPos := 0
	for i:= 0; i <= parent.N ;i++{
		if self == parent.C[i] {
			selfPos = i
			break
		}
	}
	tmpArray := []*BTreeNodeInfo{}
	isLeft := false
	if selfPos != 0 {
		brother := parent.C[selfPos-1]
		if brother.N > NodeKey_min{
			tmpArray = append(tmpArray,brother)
			isLeft = true
		}
	}
	if selfPos != parent.N {
		brother := parent.C[selfPos+1]
		if brother.N > NodeKey_min{
			tmpArray = append(tmpArray,brother)
		}
	}



	if len(tmpArray) > 0 {
		pBorrow := tmpArray[0]
		n_log.Info("isleft  %v  keys %v  selfpos  %v",isLeft,pBorrow.Keys[:pBorrow.N],selfPos)

		ptmp := parent.Keys[selfPos]

		if isLeft {	// 向左借最大值
			for i := self.N; i >0 ;i--{
				self.Keys[i] = self.Keys[i-1]
			}

			ptmp := parent.Keys[selfPos -1]
			parent.Keys[selfPos-1],pBorrow.Keys[pBorrow.N-1] = pBorrow.Keys[pBorrow.N-1],parent.Keys[selfPos-1]
			pBC0 := pBorrow.C[pBorrow.N]
			pBorrow.N--
			self.Keys[0] = ptmp
			// 这是错误的  ok   还在节点没关
			if !self.IsLeaf {
				for i:= self.N+1; i >=1;i--{
					self.C[i] = self.C[i-1]
				}

				self.C[0] = pBC0
				pBC0.P = self
			}

			self.N++
			n_log.Info("borro  %v  ",pBorrow.Keys[:pBorrow.N])
		}else {	// 向右 最小值
			parent.Keys[selfPos] = pBorrow.Keys[0]
			for i:=0; i<pBorrow.N-1;i++{
				pBorrow.Keys[i] = pBorrow.Keys[i+1]
			}
			pBC0 := pBorrow.C[0]

			if !self.IsLeaf {
				// pborrwo 要移动
				for i:=0; i < pBorrow.N;i++{
					pBorrow.C[i] = pBorrow.C[i+1]
				}
			}

			pBorrow.N--
			self.Keys[self.N] = ptmp
			self.N++
			if !self.IsLeaf {
				self.C[self.N] = pBC0
				pBC0.P = self
			}
		}

		return
	}

	// 这儿向上合并  ？？？
	/*
	1 剩余keys 和 父节点关键字 一起合并 向左或者向右 递归处理 ok？
	 */
	self.CheckMerge()

}

func (self *BTreeNodeInfo) CheckMerge()  {
	if self.N >= NodeKey_min {
		return
	}

	parent := self.P
	if parent == nil {
		return
	}
	selfPos := 0
	for i:= 0; i <= parent.N ;i++{
		if self == parent.C[i] {
			selfPos = i
			break
		}
	}
	// 统一向左合并
	isLeft := true


	var pMerge *BTreeNodeInfo
	if selfPos == 0 {
		isLeft = false
	}else {
		selfPos--
	}
	pkey := parent.Keys[selfPos]

	n_log.Info("shazi hebing %v  %v  %v   %v",isLeft,selfPos,self.Keys,self.N)

	if isLeft {
		pMerge = parent.C[selfPos]
		// 父节点删除当前节点 删除当前key
		for i:= selfPos; i < parent.N-1 ;i++{
			parent.Keys[i] = parent.Keys[i+1]
		}
		for i:= selfPos + 1; i <= parent.N-1 ;i++{
			parent.C[i] = parent.C[i+1]
		}

		parent.N--
		if parent.N <= 0 {
			// 节点空了
			pTree.Root = pMerge
			pMerge.P = nil
			//pMerge.IsLeaf = true
		}

		pMerge.Keys[pMerge.N] = pkey
		for i:=0;i <self.N;i++{
			pMerge.Keys[pMerge.N+i+1] = self.Keys[i]
		}
		if !self.IsLeaf {
			for i:=0; i <= self.N;i++{
				self.C[i].P = pMerge
				pMerge.C[pMerge.N+i+1] = self.C[i]
			}
		}
		pMerge.N += self.N +1

	}else {
		pMerge = parent.C[selfPos+1]
		// 好好想下 这儿咋搞 ？？？

		//pTree.Traverse()

		parent.Print()
		n_log.Info("paren  %v  %v  perge  %v  %v  %p",parent.N,parent.Keys,pMerge.N,pMerge.Keys,parent)


		for i:= selfPos; i < parent.N-1 ;i++{
			parent.Keys[i] = parent.Keys[i+1]
		}
		for i:= selfPos; i <= parent.N-1 ;i++{
			parent.C[i] = parent.C[i+1]
		}

		parent.N--
		if parent.N <= 0 {
			// 节点空了
			pMerge.P = nil
			pTree.Root = pMerge

			//pMerge.IsLeaf = true
		}

		parent.Print()

		for i:= pMerge.N; i> 0;i--{
//			n_log.Info("i   %v",i)
			pMerge.Keys[i+self.N] = pMerge.Keys[i-1]
		}
		if !pMerge.IsLeaf{
			for i:= pMerge.N ; i>= 0;i--{
				pMerge.C[i+self.N+1] = pMerge.C[i]
			}
		}

		for i:=0; i < self.N;i++{
			pMerge.Keys[i] = self.Keys[i]
		}

		if !self.IsLeaf {
			for i:=0; i <= self.N;i++{
				pMerge.C[i] = self.C[i]
				self.C[i].P = pMerge
			}
		}

		pMerge.Keys[self.N] = pkey


		pMerge.N += self.N + 1

		// 貌似 好像对了 你觉得呢 ？ 晚上再仔细调节下

	}

	//pTree.Traverse()

	parent.Remove_isIn()
}