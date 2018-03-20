package basic_date

import (
	"testing"
	"fmt"
)

//比较两个url，返回0，s1==s2;返回1，s1>s2;返回-1，s1<s2;
func equals(s1, s2 string) int {
	ls1 := len(s1)
	ls2 := len(s2)
	for i, j := 0, 0; ; {
		if i == ls1 && j == ls2 {
			return 0
		}
		if i == ls1 {
			return -1
		}
		if j == ls2 {
			return 1
		}

		if (s1[i] == '{') {
			i = index(s1[i:], '}') + i + 1
			j = index(s2[j:], '/') + j
			continue
		}
		if (s2[j] == '{') {
			i = index(s1[i:], '/') + i
			j = index(s2[j:], '}') + j + 1
			continue
		}

		if s1[i] == s2[j] {
			i++
			j++
			continue
		}
		if s1[i] > s2[j] {
			return 1
		}
		if s1[i] < s2[j] {
			return -1
		}

	}
}

//返回字符c在字符串str中第一次出现的位置，如果没有找到，就返回字符串的长度
func index(str string, c byte) int {
	lstr := len(str)
	for i := 0; i < lstr; i++ {
		fmt.Println(str[i], "  ", c)
		if str[i] == c {
			return i
		}
	}
	return lstr
}

func TestTest(t *testing.T) {
	var s string
	s="'aaaa"
	fmt.Println(s=="")
}
