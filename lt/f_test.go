package lt

import (
	"github.com/abc463774475/bbtool/n_log"
	"strings"
	"testing"
)

func TestF722(t *testing.T) {
	str := `
	///* select *from t1 wher id = "10"
	// insert into t1 values (100)*/
	
	update t1 = 1;
`
	str1 := F722(str)
	n_log.Info ("ret \n%v\n%v",str1, strings.Fields(str1))
}