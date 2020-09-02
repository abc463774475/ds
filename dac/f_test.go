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
