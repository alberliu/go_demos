package basic_date

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap1(t *testing.T) {
	a := []interface{}{1}
	b := []interface{}{1}
	a = append(a, b...)
	fmt.Println(a)
}

func TestMap2(t *testing.T) {
	jsonStr := `{"a":1}`
	var m map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &m)
	fmt.Println(m)
}
