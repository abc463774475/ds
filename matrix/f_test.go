package matrix

import (
	"github.com/abc463774475/bbtool/n_log"
	"testing"
)

func TestMultiplication(t *testing.T) {
	a := [][]int{
		{1,2,3},
		{2,3,4},
	}

	b := [][]int{
		{1,3},
		{2,4},
		{5,5},
	}

	c := Multiplication(a,b)

	n_log.Info("c is\n%v",c)
}


