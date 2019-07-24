package b_tree

import (
	"testing"
	"bbTool/n_log"
	//"math/rand"
	"sort"
)

func TestBTree_Insert(t0 *testing.T) {
	t := CreateBTree(5)
	//array :=[]int{
	//	10,20,5,6,12,30,7,17,3,9,8,50,41,
	//}

	array :=[]int{
		//10,20,5,6,12,30,7,17,3,9,8,50,41,
		1,2,6,7,11,4,8,13,10,5,17,9,16,20,3,12,14,18,19,15,
	}
	for _,v := range array {
		t.Insert(v)
		//t.Traverse()
	}

	//t.Traverse()

	//rand.Shuffle(len(array), func(i, j int) {
	//	array[i],array[j] = array[j],array[i]
	//})
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
	n_log.Info("array  %#v",array)

	t.Traverse()

	//for _,v := range array {
	//	n_log.Erro("sssssss  %v",v)
	//	t.Remove(v)
	//
	//	t.Traverse()
	//	n_log.Erro("eeeeeeeee  %v",v)
	//}
}
