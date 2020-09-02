package b_tree_self

import (
	"github.com/abc463774475/bbTool/n_log"
	"math/rand"
	"testing"
	"time"
)

func TestBTree_Insert(t *testing.T) {
	n_log.Info("111111111")

	pTree = CreateBTree()
	array := []int{
		//5,10,15,8,9,
	}

	for i := 0 ; i < 10000;i++{
		array = append(array,i+1)
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	rand.Shuffle(len(array), func(i, j int) {
		array[i],array[j] = array[j],array[i]
	})


	//n_log.Info("array  %v",array)

	for _,v := range array{
		pTree.Insert(v)
	}

	NodeNum = 0
	//pTree.Traverse()

	//pTree.Remove(29)
	//s2 := time.Now()

	//n_log.Info("node num  %v  %v",NodeNum)

	//pTree.Search(100)

	//pTree.Traverse()

	if true {
		rand.Shuffle(len(array), func(i, j int) {
			array[i],array[j] = array[j],array[i]
		})
		for _,v := range array{

			n_log.Info("rrrrrrrrr  %v",v)
			pTree.Remove(v)


		}
	}

	pTree.Traverse()
	pTree.Remove(10002)


	time.Sleep(1*time.Second)
}
