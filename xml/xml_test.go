package xml

import (
	"testing"
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName   string `xml:"xml"`
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

func TestXml(t *testing.T) {
	user := User{Id: 1, Name: "1"}
	buf, err := xml.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}
