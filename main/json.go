package main

import (
	"github.com/json-iterator/go"
	"goweb"
	"log"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func jsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func jsonMarshal(v interface{}) ([]byte, error){
	return json.Marshal(v)
}

/*func main() {
	goweb:=goweb.NewGoWeb();
	goweb.Unmarshal=jsonUnmarshal
	goweb.Marshal=jsonMarshal

	goweb.ListenAndServe(":8000")

}*/

func Hello(){
	log.Println("hello")
}