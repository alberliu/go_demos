package slice

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	sort(slice)
	fmt.Println(slice)
}

func search(slice []int, des int) int {
	var (
		low    = 0
		high   = len(slice) - 1
		middle = 0
	)
	//确保不会出现重复查找，越界
	for low <= high {
		//计算出中间索引值
		middle = (high + low) / 2 //防止溢出
		if des == slice[middle] {
			return middle
			//判断下限
		} else if des < slice[middle] {
			high = middle - 1
			//判断上限
		} else {
			low = middle + 1
		}
	}
	//若没有，则返回-1
	return -1
}

func sort(slice []int) {
	var (
		start = 0
		end   = len(slice)
		mid   = (start + end) / 2
	)

	for start < end {
		if slice[start] < slice[mid] {
			start++
		} else {
			slice[start], slice[mid] = slice[mid], slice[start]
			start++
		}
		if slice[mid] < slice[end] {
			end--
		} else {
			slice[mid], slice[end] = slice[end], slice[mid]
			start++
		}

	}
}
