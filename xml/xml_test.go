package xml

import (
	"encoding/xml"
	"fmt"
	"testing"
)

type User struct {
	XMLName string `xml:"xml"`
	Id      int    `xml:"id"`
	Name    string `xml:"name"`
}

func TestXml(t *testing.T) {
	buf, err := xml.Marshal(User{Id: 1, Name: "2"})
	fmt.Println(string(buf), err)
	m := map[interface{}]interface{}{}
	fmt.Println(xml.Unmarshal(buf, m))
}
