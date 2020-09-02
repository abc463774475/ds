package order_alg

import "github.com/abc463774475/bbTool/n_log"

/*
	recursive function

 */

func Quick(a []int)  {
	if len(a) <= 1 {
		return
	}
	// 划分两个区域  大的 在右边 小的在左边
	mid,i := a[0],1
	head,tail := 0,len(a) - 1
	for head < tail {
		if a[i] > mid {
			a[i],a[tail] = a[tail],a[i]
			tail--
		}else {
			a[i],a[head]= a[head],a[i]
			head++
			i++
		}
	}
	a[head] = mid

	Quick(a[:head])
	Quick(a[head+1:])
}


func Quick_h(aold []int)  {
	a := append([]int{},aold...)
	if len(a) <= 1 {
		return
	}
	cur := a[0]
	aMax := []int{}
	aMin := []int{}
	for i:=1; i <len(a);i++{
		if a[i] >= cur {
			aMax = append(aMax,a[i])
		}else {
			aMin = append(aMin,a[i])
		}
	}

	a = append(aMin,cur)
	a = append(a,aMax...)
	for k,v := range a {
		aold[k] = v
	}

	n_log.Info("a old  %v",a)

	Quick_h((aold[:len(aMin)]))
	Quick_h((aold[len(aMin)+1:]))
}