package order_alg

import (
	"github.com/abc463774475/bbtool/n_log"
	"math"
)

/*
 1-50 先来 ok？？？
 */


func r_h(a []int, nayiwei int) []int {
	rIndex := make([][]int,10)
	n_log.Info("rindex  %v",rIndex)

	for _,v := range a{
		n := int(math.Pow(float64(10),float64(nayiwei)))
		pv := v/n%10
		rIndex[pv] = append(rIndex[pv],v)
	}

	rback := make([]int,0,len(a))
	for _,v := range rIndex {
		for _,v1 := range v {
			rback = append(rback,v1)
		}
	}

	return rback
}

func radix_sort(a []int)  []int{
	max := 0
	for _,v := range a {
		if v > max {
			max = v
		}
	}

	// 第一次  个位排序
	maxPos := int(math.Log10(float64(max)))


	for i:=0; i < maxPos+1;i++{
		a = r_h(a,i)
	}
	return a
}
