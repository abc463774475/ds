package order_alg

import "github.com/abc463774475/bbtool/n_log"

//func merge_h(left, right []int) (relust []int) {
//	l, r := 0, 0
//	for l < len(left) && r < len(right) {
//		if left[l] < right[r] {
//			relust = append(relust, left[l])
//			l++
//		} else {
//			relust = append(relust, right[r])
//			r++
//		}
//	}
//
//	relust = append(relust, left[l:]...)
//	relust = append(relust, right[r:]...)
//	return
//}
//
//func merge(r []int) []int {
//	if len(r) <= 1 {
//		return r
//	}
//
//	num := len(r) / 2
//	left := merge(r[:num])
//	right := merge(r[num:])
//
//	return merge_h(left, right)
//}

func Merge(sourceArr, tempArr []int, startIndex, midIndex, endIndex int) {
	i := startIndex
	j := midIndex + 1
	k := startIndex
	for i != midIndex+1 && j != endIndex+1 {
		if sourceArr[i] > sourceArr[j] {
			tempArr[k] = sourceArr[j]
			k++
			j++
		} else {
			tempArr[k] = sourceArr[i]
			k++
			i++
		}
	}
	for i != midIndex+1 {
		tempArr[k] = sourceArr[i]
		k++
		i++
	}

	for j != endIndex+1 {
		tempArr[k] = sourceArr[j]
		k++
		j++
	}

	for i = startIndex; i <= endIndex; i++ {
		sourceArr[i] = tempArr[i]
	}

}

//内部使用递归
func MergeSort(sourceArr, tempArr []int, startIndex, endIndex int) {
	var midIndex int
	if startIndex < endIndex {
		midIndex = startIndex + (endIndex-startIndex)/2 //避免溢出int
		MergeSort(sourceArr, tempArr, startIndex, midIndex)
		MergeSort(sourceArr, tempArr, midIndex+1, endIndex)
		Merge(sourceArr, tempArr, startIndex, midIndex, endIndex)
	}
}


// 合并
func m_merge(left []int,right []int) []int {
	allLen := len(left) + len(right)
	rback := make([]int,allLen)
	rp := 0
	lp := 0
	i :=0
	for i< allLen {
		if lp >= len(left) {
			rback[i] = right[rp]
			rp++
			i++
			continue
		}

		if rp >= len(right) {
			rback[i] = left[lp]
			lp++
			i++
			continue
		}

		n_log.Info("i  %v",i)
		if left[lp] < right[rp] {
			rback[i] = left[lp]
			lp++
		}else {
			rback[i] = right[rp]
			rp++
		}
		i++
	}

	return rback
}

// 递归合并
func m_sort(a []int) []int {
	r := make([]int,len(a))
	copy(r,a)
	if len(r) <= 1 {
		return r
	}
	num := len(r)/2
	left := m_sort(r[:num])
	right := m_sort(r[num:])

	return m_merge(left,right)
}
