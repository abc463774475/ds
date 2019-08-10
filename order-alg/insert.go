package order_alg

import "bbTool/n_log"

func insert(a []int)  {
	for i:=1; i < len(a);i++{
		// j 要找准位置  ok？？
		atmp := a[i]
		if a[i] > a[i-1] {
			//n_log.Info("a %v",a)
			continue
		}
		isHavePos := false
		for j :=i; j>0;j--{
			if atmp > a[j-1]{
				a[j] = atmp
				isHavePos = true
				break
			}else {
				a[j] = a[j-1]
			}
		}

		if !isHavePos {
			a[0] = atmp
		}
	}
}

func insert_gap(a []int,start,gap int)  {
	for i:=start; i < len(a);i+= gap{
		if i == start {
			continue
		}
		// j 要找准位置  ok？？
		atmp := a[i]
		if a[i] > a[i-gap] {
			//n_log.Info("a %v",a)
			continue
		}
		isHavePos := false
		for j :=i; j>start; j-= gap{
			if atmp > a[j-gap]{
				a[j] = atmp
				isHavePos = true
				break
			}else {
				a[j] = a[j-gap]
			}
		}

		if !isHavePos {
			a[start] = atmp
		}
	}
}

func insert_opt(a []int)  {
	for i:=0;i < len(a);i++{
		if i == 0 {
			continue
		}
		atmp := a[i]
		j := i-1

		for (j>=0) && (a[j] > atmp){
			a[j+1] = a[j]
			j--
		}

		if j != i-1 {
			a[j+1] = atmp
		}

		n_log.Info(" j %v a   %v",j,a)
	}
}
