package collection_test

import (
	"bufio"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"os"
	"sort"
	"strings"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	x := 11
	s := []int{3, 6, 8, 11, 45} //注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, "在s中的位置为：", pos)
	} else {
		fmt.Println("s不包含元素", x)
	}
}

func TestData(t *testing.T) {
	fmt.Println(GetFirstLetter("按"))
	fmt.Println(GetFirstLetter("a"))
	fmt.Println(GetFirstLetter("*"))
}

func Fallback(r rune, a pinyin.Args) []string {
	return []string{strings.ToLower(string(r))}
}

func GetFirstLetter(str string) string {
	a := pinyin.NewArgs()
	a.Style = pinyin.FIRST_LETTER
	a.Fallback = Fallback
	letters := pinyin.Pinyin(str, a)
	if len(letters) > 0 && len(letters[0]) > 0 {
		return letters[0][0]
	}
	return ""
}

type Log struct {
	Obj  string `json:"obj"`
	Func string `json:"func"`
}

type Param struct {
	Key   string
	Value string
}

var params []Param

func TestJson(t *testing.T) {
	file, err := os.Open("/Users/admin/project/go_demos/test.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	var (
		index      = -1
		key, value string
	)

	for {
		index++
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}

		if index%2 == 0 {
			key = string(line)
			continue
		}

		value = string(line)
		params = append(params, Param{
			Key:   key,
			Value: value,
		})
	}

	for _, v := range params {
		fmt.Printf("%s %s `gorm:\"column:%s\" json:\"%s\"` %s %s\n", v.Key, "string", strings.ToLower(v.Key), strings.ToLower(v.Key), "//", v.Value)
	}

	fmt.Println()

	var str string
	for _, v := range params {
		str = str + fmt.Sprintf("%s=__%s__&", strings.ToLower(v.Key), v.Key)
	}
	fmt.Println(str)
}

func TestData2(t *testing.T) {
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().AddDate(0, -3, 0).Unix())
	fmt.Println(time.Now().UnixNano() / 1000000)
}
