package order_alg

import "github.com/abc463774475/bbtool/n_log"

/*
 这个排序需要优化插入排序之后才能做出
 */
func shell(a []int)  {
	for gap := len(a)/2; gap > 0; gap/=2 {
		for i:=gap; i < len(a);i++{
			for j :=i-gap; j>=0 &&a[j+gap] < a[j];j -=gap{
				a[j],a[j+gap] = a[j+gap],a[j]
			}
		}
	}
}


/*
完全不用担心 shell 排序 因为  gap 永远 会 变成 1  为1的时候就成为 insert排序了，，只是他的效率会更高
 */
func shell_h(a []int)  {
	for gap:=len(a)/2 ; gap > 0;gap /= 2{
		// 插入排序  有多少组  ok？
		for i:=0; i < gap;i++{
			// 对这一组  进行插入排序

			/*
			0 - gap-1 组 每一组
			[0，gap,gap*2,gap*3]
			[1，gap+1,gap*2+1,gap*3+1]
			[2，gap+2,gap*2+2,gap*3+2]
			[i，gap+i,gap*2+i,gap*3+i]
			.
			.
			.
			[gap-1]
			分别进行插入排序
			 */
			for j:= i ; j < len(a);j+= gap {
				n_log.Info("i  %v  gap  %v",i,gap)
				insert_gap(a,i,gap)
			}
		}
	}
}
