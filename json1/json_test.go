package json1

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

// d-inter.caijiyouxi.com
func ReportBigLogEvent(eventId string, params interface{}) {
	request := httplib.Post("http://d-inter.caijiyouxi.com/t.png")

	request.Param("act", "server_customize_event")
	request.Param("app_id", "10008")
	request.Param("time", strconv.FormatInt(time.Now().Unix(), 10))
	request.Param("event_id", eventId)

	bytes, _ := json.Marshal(params)
	request.Param("params", string(bytes))

	resp, err := request.SetTimeout(time.Second, time.Second).Response()
	fmt.Println(err)
	fmt.Println(resp)
}

func TestJson(t *testing.T) {
	ReportBigLogEvent("alber_test1", map[string]string{
		"name": "hello,国豪",
		"desc": "这是一个测试数据",
	})
}

func a() []int {
	return nil
}

func b() []int {
	return []int{}
}

func TestNil(t *testing.T) {
	var userId int64 = 57827402
	userId = (int64(2) << 48) | userId
	fmt.Println(userId)

	fmt.Println(562951561014314 - (int64(2) << 48))
}

func Decode(key string) (userId int64, seq int64, err error) {
	var (
		idx int
		t   int64
	)
	if idx = strings.IndexByte(key, '_'); idx == -1 {
		return
	}
	if userId, err = strconv.ParseInt(key[:idx], 10, 64); err != nil {
		return
	}
	if t, err = strconv.ParseInt(key[idx+1:], 10, 64); err != nil {
		return
	}
	seq = t
	return
}

func TestUn(t *testing.T) {
	fmt.Println(562949954090114 | 1)

}

func TestRand(t *testing.T) {
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
	fmt.Println(rand.Int63n(5))
}
