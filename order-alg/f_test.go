package order_alg

import (
	"testing"
	"bbTool/n_log"
	"math/rand"
)

func TestBubbling(t *testing.T) {
	a := []int{
		1,3,8,2,6,10,3,2,
	}

	//a := []int{}
	//for i:=0;i <50;i++{
	//	a = append(a, i+1)
	//}
	//a = append(a,1)
	//a = append(a,5)
	//a = append(a,5)
	//a = append(a,3)
	//rand.Shuffle(len(a), func(i, j int) {
	//	a[i],a[j] =a[j],a[i]
	//})

	atmp := make([]int,len(a))
	MergeSort(a,atmp,0,len(a) -1)

	n_log.Info("a   %v",a)
}

func TestQuick(t *testing.T) {
	a := []int{
		//3,8,2,6,10,1,
		2, 8, 5, 1, 10, 3, 4, 6, 9, 7,
	}
	for i:=0; i < 10 ; i++{
		a = append(a,i+1)
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i],a[j] = a[j],a[i]
	})


	n_log.Info("a   %v",a)
	//
	return
	a = radix_sort(a)

	n_log.Info("a   %v",a)
}

func f1(a []int)  {
	a = append(a,[]int{
		1,2,3,4,5,6,7,8,9,10,
	}...)
	a[2] = 100
	n_log.Info("a  %v",a)
}

func Test111( t *testing.T)  {
	a := make([]int,100,200)
	a[0] = 1
	a[1] = 2
	a[2] = 3


	b := a[70:]
	b[0] = 100000

	c := a[:5]

	//f1 (a)
	n_log.Info("a  %v len  %v cap  %v",a,len(a),cap(a))
	n_log.Info("b  %v len %v cap %v",b,len(b),cap(b))

	n_log.Info("c  %v  len  %v cap  %v",c,len(c),cap(c))
}

func quick(a[]int)  {
	if len(a) <= 1{
		return
	}
	head,tail := 0,len(a) -1
	mid,i := a[0],1

	for head < tail{
		if a[i] > mid {
			a[i],a[tail] = a[tail],a[i]
			tail--
		}else {
			a[i],a[head] = a[head],a[i]
			head++
			i++
		}
	}
	a[head] = mid

	quick(a[:head])
	quick(a[head+1:])
}

func bubble(a []int)  {
	for i:= 0 ; i < len(a) ; i++{
		for j:= 0 ;j < len(a) - i -1;j++{
			if a[j] > a[j+1]{
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}

/*
	select 每次大小 ok？
 */

func select_order(a []int)  {
	for i:= 0; i < len(a) -1;i++{
		tmp := i
		for j := i;j < len(a);j++{
			if a[tmp] > a[j] {
				tmp = j
			}
		}

		if tmp != i {
			a[i],a[tmp] = a[tmp],a[i]
		}

	}
}

func bucket(a []int)  {
	max := 0
	for _,v := range a {
		if v < 0 {
			n_log.Panic("bucket  negative %v",a)
		}
		if v > max {
			max = v
		}
	}

	atmp := make([]int,max+1)
	for _,v := range a{
		atmp[v]++
	}

	pos := 0
	for k,v := range atmp {
		if v == 0 {
			continue
		}

		for j:=0; j < v ;j++{
			a[pos] = k
			pos++
		}
	}
}

func quick_h(a []int)  {
	if len(a) <= 1{
		return
	}
	head,tail := 0,len(a) - 1
	mid,i := a[0],1
	for head < tail {
		if a[i] > mid {
			a[i],a[tail] = a[tail],a[i]
			tail--
		}else {
			a[i],a[head] = a[head],a[i]
			head++
			i++
		}
	}
	a[head] = mid

	quick(a[:head])
	quick(a[head+1:])
}

func max_heap_h(a []int,start,end int)  {
	dad := start
	son := 2*dad+1

	for son <= end{
		if son <= end -1 && a[son] <a[son+1] {
			son++
		}
		if a[dad] > a[son] {
			return
		}
		a[son],a[dad] = a[dad],a[son]
		dad = son
		son = 2*dad+1
	}
}

func heap_h(a[]int)  {
	for i:=len(a)/2-1;i>=0 ;i--{
		max_heap_h(a, i,len(a) -1)
	}

	for i := len(a) -1;i >0;i--{
		a[i],a[0] = a[0],a[i]
		max_heap_h(a,0,i - 1)
	}
}

func TestMerge(t *testing.T) {
	r1 := m_merge([]int{
		1,7,10,19,
	},[]int{
		2,9,10,19,25,
	})
	n_log.Info("r1   %v",r1)
}

func TestR_h(t *testing.T)  {
	r_h([]int{

	},0)

}