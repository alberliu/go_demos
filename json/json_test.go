package json_test

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

var jsonStr = `{
  "receiver_type":3,
  "receiver_id":2,
  "to_user_ids":[],
  "message_id":"123456",
  "send_time":123455,
  "message_body":{
  	"message_type":1,
  	"message_content":{
  		"text":"large group message"
  	}
  }
}`

//json-iterator测试
func TestJsoniter(t *testing.T) {
	var str interface{}
	jsoniter.Get([]byte(jsonStr), "message_body", "message_content").ToVal(&str)
	body, _ := jsoniter.Marshal(str)

	fmt.Println(str)
	fmt.Println(string(body))
}
