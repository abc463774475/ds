package order_alg

func Select(a []int)  {
	for i:=0; i< len(a) ;i++{

		tmp := i
		for j:= i+1;j < len(a);j++{
			if a[j] < a[tmp]{
				tmp = j
			}
		}

		if tmp != i {
			a[tmp],a[i] = a[i],a[tmp]
		}
	}

}
