package beego

import (
	"time"
	"fmt"
	"github.com/astaxie/beego/orm"
	"encoding/json"
)

type Time struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON（尚未实现）
func (t *Time) UnmarshalJSON(data []byte) error {
	var err error
	t.Time,err=time.Parse("2006-01-02 15:04:05", string(data)[1:20])
	return err
}

func (t *Time) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

func (t *Time) FieldType() int {
	return orm.TypeDateTimeField

}

func (t *Time) SetRaw(value interface{}) error {
	fmt.Println("setraw")
	switch value.(type) {
	case time.Time:
		t.Time = value.(time.Time)

	}
	return nil
}

func (t *Time) RawValue() interface{} {
	fmt.Println("rawvalue")
	str:=t.Format("2006-01-02 15:04:05")
	if str=="0001-01-01 00:00:00"{
		return nil
	}
	return str
}


type Money struct {
	Value int64
}