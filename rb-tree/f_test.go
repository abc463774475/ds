package rb_tree

import (
	"testing"
	"bbTool/n_log"
	"math/rand"
	"time"
)

// 速度好慢 艹
func TestTreeInfo_Insert(t *testing.T) {
	a := []int{}
	for i:= 0 ; i < 500 ;i++{
		a = append(a,i+1)
	}

	rand.Seed(int64(time.Now().Nanosecond()))

	rand.Shuffle(len(a), func(i, j int) {
		//a[i],a[j] = a[j],a[i]
	})

	//a =[]int{
	//	6 ,7 ,9 ,10 ,3 ,4 ,5 ,2 ,8 ,1,
	//}

	n_log.Info("a   %v",a)


	start := time.Now()
	for _,v := range a{
//		n_log.Info("i   %v",v)
		g_tree.Insert(v)

	}

	g_tree.Traverser()

	n_log.Info("total time %v",time.Now().Sub(start))

	rand.Shuffle(len(a), func(i, j int) {
		a[i],a[j] = a[j],a[i]
	})
	//a = []int{8, 9, 4, 1, 7, 5, 10, 2, 6, 3}

	n_log.Info("a   %#v",a)
	for _,v := range a {
		if true || v == 8{
			n_log.Info("remove  i  %v",v)
			g_tree.Remove(v)
			//g_tree.Traverser()
		}
	}
	//g_tree.Remove(4)

	g_tree.Traverser()
}

func TestNodeInfo_GetChangeNode_h(t *testing.T) {
	p1 := CreateNodeInfo(1)
	p1.IsRed = false
	p5 := CreateNodeInfo(5)
	p5.IsRed = false

	p7 := CreateNodeInfo(7)
	p7.IsRed = false
	p5.AddOnePos(1,p7)

	p10 := CreateNodeInfo(10)
	p10.IsRed = false

	p20 := CreateNodeInfo(20)
	p20.IsRed = false

	p15 := CreateNodeInfo(15)
	p15.IsRed = false

	p30 := CreateNodeInfo(30)
	p30.IsRed = false
	p5.AddOnePos(0,p1)
	p20.AddOnePos(0,p15)
	p20.AddOnePos(1,p30)

	p10.AddOnePos(0,p5)
	p10.AddOnePos(1,p20)

	p10.Traverser()

	n_log.Info("111111111111111")
	//ptmp := p10
	//changeNode = ptmp
	//pret := ptmp.GetChangeNode_h(0)
	//
	//PrintTotalChangeNode()
	//
	//FreshTotalChange()
	g_tree.Root = p10

	p10.Remove(10)
	p10.Traverser()

	//n_log.Info("pret  is red  %v  key  %v",pret.IsRed,pret.Key)
}

func TestNodeInfo_GetChangeNode_h11(t *testing.T) {
	p1 := CreateNodeInfo(1)
	p1.IsRed = true
	p3 := CreateNodeInfo(3)
	p3.IsRed = false

	p13 := CreateNodeInfo(13)
	p13.IsRed = false

	p20 := CreateNodeInfo(20)
	p20.IsRed = true

	p17 := CreateNodeInfo(17)
	p17.IsRed = false

	p3.AddOnePos(0,p1)
	p17.AddOnePos(1,p20)
	{
		p14 := CreateNodeInfo(14)
		p14.IsRed = true
		p17.AddOnePos(0,p14)
	}

	p13.AddOnePos(0,p3)
	p13.AddOnePos(1,p17)


	g_tree.Root = p13

	g_tree.Traverser()
	n_log.Info("111111111111111")

	a := []int{1,3,13,14,17,20}
	rand.Seed(int64(time.Now().Nanosecond()))
	rand.Shuffle(len(a), func(i, j int) {
		a[i],a[j] = a[j],a[i]
	})

	for _,v := range a{
		//p1 := g_tree.Search(17).getExchangeNode()
		n_log.Info("1111111111  %v  %v",v)
		g_tree.Remove(v)
		g_tree.Traverser()

	}


}