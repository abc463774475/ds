package order_alg

func  max_heapify(a []int,start,end int)  {
	dad := start
	son := dad*2+1

	for son <= end{
		if son+1 <= end && a[son] < a[son+1]{
			son++
		}
		if a[dad] > a[son] {
			return
		}else {
			a[dad],a[son] = a[son],a[dad]
			dad = son

			son = dad*2+1
		}
	}
}

func Heap(a []int)  {
	l := len(a)

	for i:= l/2-1;i>=0 ;i--{
		max_heapify(a,i,l -1)
	}

	for i:= l -1; i >0;i--{
		a[0],a[i] = a[i],a[0]

		max_heapify(a,0,i-1)
	}
}
