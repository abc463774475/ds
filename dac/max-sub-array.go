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

// dont think about slice len is zero
func max(vs ...int) int  {
	m := vs[0]
	for i := 1 ; i < len(vs); i++{
		if vs[i] > m {
			m = vs[i]
		}
	}
	return m
}

// 求最大和值 _type :0  from [0:n-1]  :-1 from[n-1:0]
func maxSum(a []int, _type int) int {
	if _type == 0 {
		m := 0
		sum := 0
		for _,v := range a{
			sum += v
			if sum > m {
				m = sum
			}
		}
		return m
	}else {
		m := 0
		sum := 0
		for i:= len(a) -1 ;i >=0 ;i--{
			sum += a[i]
			if sum > m {
				m = sum
			}
		}

		return m
	}
}

/*
	获取最大子数组   分治法
 */
func sub_array_sum_max(a []int) int {
	if len(a) == 0 {
		n_log.Panic("a  zero")
	}
	if len(a) == 1 {
		return a[0]
	}
	mid := len(a)/2
	maxLeft := sub_array_sum_max(a[:mid])
	maxRight := sub_array_sum_max(a[mid:])
	maxMidle := 0
	// 中间往左   中间往右的处理
	m_l := maxSum (a[:mid],-1)
	m_r := maxSum(a[mid:],0)


	maxMidle = m_l + m_r
	return max(maxLeft,maxRight,maxMidle)
}

