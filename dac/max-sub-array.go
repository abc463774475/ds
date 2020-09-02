package dac

import "github.com/abc463774475/bbtool/n_log"

/*
  获取 前后最大差值
  最小值  最大值  最大差  分治法   666
 */

func max_sub_array(a []int) (int,int,int) {
	if len(a) == 0 {
		n_log.Panic("cannot ",)
	}
	if len(a) == 1 {
		return a[0],a[0],0
	}
	if len(a) == 2 {
		if a[0] > a[1] {
			return a[1],a[0],a[1]-a[0]
		}else {
			return a[0],a[1],a[1]-a[0]
		}
	}

	l := len(a)/2
	i1,j1,k1 := max_sub_array(a[:l])
	i2,j2,k2 := max_sub_array(a[l:])
	k3 := j2 - i1

	// 数值差
	max_cha := k1
	min_i := i1
	max_j := j1
	if max_cha < k2 {
		max_cha = k2
	}
	if max_cha < k3 {
		max_cha = k3
	}
	if min_i > i2 {
		min_i = i2
	}
	if max_j < j2 {
		max_j = j2
	}

	return min_i,max_j,max_cha
}

/*
	获取最大子数组   分治法
 */

func sub_array_sum_max(a []int)  {


}

