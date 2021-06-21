/*
  c++  去注释算法
/*
 */

package lt

const (
	comment1_start = "//"
	comment1_end   = "\n"

	comment2_start = "/*"
	comment2_end = "*/"
)

func F722(str string) string {
	retData := make([]byte,0,len(str))

	chars_startPos := -1
	comment_type := -1
	// comment_startPos := -1

	for i :=0 ; i < len(str) ; {
		if str[i] == '"' {
			// 注释中
			if comment_type != -1 {
				//retData = append(retData, str[i])
				i++
				continue
			}

			if chars_startPos == -1 {
				// 字符串开始
				chars_startPos = i
			}else {
				// 字符串结束
				chars_startPos = -1
			}

			retData = append(retData, str[i])
			i++
			continue
		}

		// 字符串中
		if chars_startPos != -1 {
			retData = append(retData, str[i])
			i++
			continue
		}

		if i == len(str) - 1 {
			retData = append(retData, str[i])
			break
		}

		if comment_type != -1 {
			if comment_type == 0 {
				if str[i:i+1] == comment1_end {
					i += 1
					comment_type = -1

					// comment_startPos = -1
				}else {
					// retData = append(retData, str[i])
					i++
				}
			} else {
				if str[i:i+2] == comment2_end {
					i += 2
					comment_type = -1

					// comment_startPos = -1
				}else {
					// retData = append(retData, str[i])
					i++
				}
			}

			continue
		}

		strTmp := str[i:i+2]
		if strTmp == comment1_start {
			// comment_startPos = i
			comment_type = 0
			i += 2
		}else if strTmp == comment2_start {
			// comment_startPos = i
			comment_type = 1
			i+=2
		}else {
			retData = append(retData, str[i])
			i++
		}
	}

	return string(retData)
}


func removeComments(strs []string) []string {
	retData := make([]string,0)
	comment_type := -1

	for i := range strs {
		if comment_type == 1 {
			for j := 0; j < len(strs);  {

			}
		}else {

		}
	}


	return retData
}



