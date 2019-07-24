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


	n_log.Info("a   %v",a)
}

func TestQuick(t *testing.T) {
	a := []int{
		//3,8,2,6,10,1,
	}
	for i:=0; i < 50 ; i++{
		a = append(a,i)
	}

	rand.Shuffle(len(a), func(i, j int) {
		a[i],a[j] = a[j],a[i]
	})

	Quick_h(a)

	n_log.Info("a   %v",a)
}

func f1(a []int)  {
	a = []int{
	//	11,23,
	}



	n_log.Info("a  %v",a)
}

func Test111( t *testing.T)  {
	a := []int{
		1,3,5,
	}
	n_log.Info("a  %v",a)
	f1 (a)

	n_log.Info("a  %v",a)
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

func max_heap(a []int,start,end int)  {
	//建立父节点指标和子节点指标
	dad := start
	son := dad*2 + 1

	for son <= end{
		//先比较两个子节点大小，选择最大的
		if son +1 <= end && a[son] < a[son+1] {
			son++
		}
		if a[dad] > a[son] { //如果父节点大於子节点代表调整完毕，直接跳出函数
			return
		}else { //否则交换父子内容再继续子节点和孙节点比较
			a[dad],a[son] = a[son],a[dad]
			dad = son
			son = dad*2+1
		}
	}
}

func heap(a []int)  {
	//初始化，i从最後一个父节点开始调整
	for i := len(a)/2-1; i >=0;i--{
		max_heap(a,i,len(a) -1)
	}

	//先将第一个元素和已排好元素前一位做交换，再重新调整，直到排序完毕
	for i:= len(a) - 1;i >0;i--{
		a[0],a[i] = a[i],a[0]
		max_heap(a,0,i-1)
	}
}



