package template

import (
	"testing"

	"os"
	"fmt"
	"text/template"
)

type User struct {
	Id   string
	Name string
}

func TestTemp(t *testing.T) {
	temp := template.Must(template.ParseFiles("test.txt"))

	user:=User{Id:"111",Name:"111"}

	file, err := os.OpenFile("out.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err =temp.Execute(file,user)
	if err!=nil{
		fmt.Println(err)
	}
}


