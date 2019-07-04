package reg

import (
	"fmt"
	"testing"
)

func TestReg(t *testing.T) {
	fmt.Println(IsExcludeId(111111))
	fmt.Println(IsExcludeId(222222))
	fmt.Println(IsExcludeId(123456))
	fmt.Println(IsExcludeId(654321))
	fmt.Println(IsExcludeId(111222))
	fmt.Println(IsExcludeId(333222))
	fmt.Println(IsExcludeId(112233))
	fmt.Println(IsExcludeId(121212))
	fmt.Println(IsExcludeId(787878))
	fmt.Println(IsExcludeId(123123))
	fmt.Println(IsExcludeId(678678))
	fmt.Println(IsExcludeId(12458963))

}

// IsExcludeId 检查是否是排除掉的id
func IsExcludeId(id int64) bool {
	str := fmt.Sprint(id)
	if isAllSame(str) {
		return true
	}
	if isIncreaseOneByOne(str) {
		return true
	}
	if isReduceOneByOne(str) {
		return true
	}
	if IsDoubleSame(str) {
		return true
	}
	if IsCircle2Same(str) {
		return true
	}
	if IsCircle3Same(str) {
		return true
	}
	return false
}

// isAllSame 是否字符一致
func isAllSame(str string) bool {
	for i := 1; i < len(str); i++ {
		if str[0] != str[i] {
			return false
		}
	}
	return true
}

// isIncreaseOneByOne 是否是连续递增的
func isIncreaseOneByOne(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1]-1 {
			return false
		}
	}
	return true
}

// isReduceOneByOne 是否是连续递减的
func isReduceOneByOne(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		if str[i] != str[i+1]+1 {
			return false
		}
	}
	return true
}

// IsDoubleSame 是否是连续两个以上重复
func IsDoubleSame(str string) bool {
	var (
		i         = 0
		sameTimes = 0
	)
	for i < len(str)-1 {
		if str[i] == str[i+1] {
			sameTimes++
			i++
			continue
		} else {
			if sameTimes == 0 {
				return false
			} else {
				sameTimes = 0
				i++
			}
		}
	}
	return true
}

// IsCircle2Same 是否循环重复
func IsCircle2Same(str string) bool {
	var i = 0
	for len(str)-i >= 4 {
		if str[i:i+2] == str[i+2:i+4] {
			i += 2
		} else {
			return false
		}
	}
	return true
}

// IsCircle3Same 是否循环重复
func IsCircle3Same(str string) bool {
	var i = 0
	for len(str)-i >= 6 {
		if str[i:i+3] == str[i+3:i+6] {
			i += 3
		} else {
			return false
		}
	}
	return true
}
