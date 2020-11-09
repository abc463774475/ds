package dac

import (
	"github.com/abc463774475/bbtool/n_log"
	"math/rand"
	"testing"
	"time"
)

func TestInfo(t *testing.T)  {
	rand.Seed(int64(time.Now().Nanosecond()))

	for z := 0 ; z < 10;z++{
		a := []int{10,-20,70,100,-30,60}
		rand.Shuffle(len(a), func(i, j int) {
			a[i],a[j] = a[j],a[i]
		})
		i,j,k := max_sub_array(a)
		n_log.Info("a  %v",a)
		n_log.Info ("111111   i %v   j  %v  k  %v",i,j,k)
	}
}

func TestMaxSubSum(t *testing.T)  {
	//a := []int{1,-2,3,5,-3,2}
	a := []int{0,-2,3,5,-1,2}
	m := sub_array_sum_max(a)
	n_log.Info("maxsubsum  is  %v",m)
}