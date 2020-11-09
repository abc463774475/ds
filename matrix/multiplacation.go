package matrix

// O(m*n*z    a m*n   b n*z)
func Multiplication(a [][]int,b [][]int) [][]int  {
	//rowLine := len(a)
	columnLine := len(a[0])
	c := make([][]int,len(a))
	for i := 0 ; i < len(c); i++{
		c[i] = make([]int, len(b[0]))
	}

	// n_log.Info("c  is  %v",c)
	// c 构造完成
	for i:=0; i < len(c);i++{
		curRow := c[i]
		for j:=0 ;j < len(curRow); j++{
			column := 0

			for k := 0 ; k < columnLine;k++ {
				column += a[i][k]*b[k][j]
			}

			curRow[j] = column
		}
	}

	return c
}