package order_alg


/*
	这玩意儿很容易理解
	以前 做 c++的时候 用到过 绝对的  所以啊，，，我自己缺少很多基础知识
 */

 // 要求正整数
func Bucket(a []int)  {
	atmp := []int{}
	max := 0
	for _,v := range a{
		if v >= max {
			max = v
		}
	}
	atmp = make([]int,max+1)
	for _,v := range a {
		atmp[v]++
	}

	j := 0
	for k,v := range atmp{
		if v > 0{
			for z:=0;z < v;z++{
				a[j] = k
				j++
			}
		}
	}
}
