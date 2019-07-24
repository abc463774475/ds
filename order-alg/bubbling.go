package order_alg

func Bubbling(a []int)  {
	//for i := 0 ; i < len(a) ;i++{
	//	for j := i+1;j < len(a);j++{
	//		if a[i] >= a[j] {
	//			a[i],a[j] = a[j],a[i]
	//		}
	//	}  // select order   ok?
	//}

	// 个人感觉两周 方式都可以  ？？？？木有问题
	for i:=0;i < len(a) ;i++{
		for j:=0; j < len(a) -i-1;j++{
			if a[j] > a[j+1] {
				a[j],a[j+1] = a[j+1],a[j]
			}
		}
	}
}