package xml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"testing"
)

type User struct {
	XMLName   string `xml:"xml"`
	Id        int    `xml:"id"`
	Name      string `xml:"name"`
	Interface interface{}
}

func TestXml(t *testing.T) {
	buf, err := xml.Marshal(User{Id: 1, Name: "2"})
	fmt.Println(string(buf), err)
	m := map[interface{}]interface{}{}
	fmt.Println(xml.Unmarshal(buf, m))
}

func TestJson(t *testing.T) {
	var user User
	json.Unmarshal([]byte(`{
  "Id":1,
  "Interface": [
      {
        "From_Account": "144115197276518801",
        "IsPlaceMsg": 0,
        "MsgBody": [
          {
            "MsgContent": {
              "Data": "\b\u0001\u0010\u0006\u001A\u0006猫瞳",
              "Desc": "MIF",
              "Ext": ""
            },
            "MsgType": "TIMCustomElem"
          },
          {
            "MsgContent": {
              "Data": "",
              "Index": 15
            },
            "MsgType": "TIMFaceElem"
          }
        ],
        "MsgRandom": 51083293,
        "MsgSeq": 7803321,
        "MsgTimeStamp": 1458721802
      }
    ]
}`), &user)
	fmt.Println(user)
	b, _ := json.Marshal(user.Interface)
	fmt.Println(string(b))

}
