package conf

import (
	"testing"
	"github.com/astaxie/beego/config"
	"fmt"
	"strconv"
)

func TestIni(t *testing.T){
	iniconf, err := config.NewConfig("ini", "conf.ini")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(iniconf.String("a"))
	iniconf.Set("a",strconv.Itoa(4)  )
	fmt.Println(iniconf.String("a"))
	iniconf.SaveConfigFile("conf.ini")
}
