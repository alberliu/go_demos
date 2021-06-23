package json1

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"strconv"
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
