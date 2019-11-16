package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"step-wx/lib"
	"testing"
)

var str = "jsapi_ticket=kgt8ON7yVITDhtdwci0qea0fTyIk0GytwiOPtFes6YqtaeUZgEg6R9ZqXD7MOiFaIqGA5skQu3Mt9NiSrDQZHQ&noncestr=OZDTSZWLEA&timestamp=1547191524&url=rtwert"
var str2 = "jsapi_ticket=kgt8ON7yVITDhtdwci0qea0fTyIk0GytwiOPtFes6YqtaeUZgEg6R9ZqXD7MOiFaIqGA5skQu3Mt9NiSrDQZHQ&nonceStr=OZDTSZWLEA&timestamp=1547191524&url=rtwert"

func TestSha1(t *testing.T) {
	h := sha1.New()
	h.Write([]byte(str))
	signature := hex.EncodeToString(h.Sum(nil))
	fmt.Println(signature)
}

func TestSha12(t *testing.T) {
	fmt.Println(lib.RandInt64(1, 2))

}
