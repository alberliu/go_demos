package httpclient

import (
	"testing"
	"github.com/astaxie/beego/httplib"
)

func TestBeegoGet(t *testing.T){
	req:=httplib.Get("http://beego.me/")
	req.ToJSON()

	s:=`13_D6tr6nBn49MJGCyPfYteQkCpM0n5m71jxWlMfZSJKdiAHBo9gDb2s3GGqaWKV2gWhgtfpzpAVEpSEu58rd_e1FYubyzdorS4UGmu9O34_ogNgchFbnJmbeFHkGIRZDhAEAWEI`
	a:=`13_CPWGh1K0L4z5THJlAgH8rb8XcnhUs01Cq8BjNPOdnX9nrzN9_xBgWwIPlxu5SK4wkF-WeFkvSKCqjJ5IYHuDSl50Nu6Ckzj9WThC8YJ3frNkqdK9pZAe_n3d9XlasWtkFsSyWy1Tp7rHqpn-YJHhADATEC`
}

func TestBeegoPost(t *testing.T){

}


