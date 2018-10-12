package wx

import (
	"fmt"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"strings"
	"encoding/json"
	"errors"
	"testing"
	"time"
	"step/domain/lib"
)

var errorCode = map[string]int{
	"IllegalAesKey":     -41001,
	"IllegalIv":         -41002,
	"IllegalBuffer":     -41003,
	"DecodeBase64Error": -41004,
}

// WxBizDataCrypt represents an active WxBizDataCrypt object
type WxBizDataCrypt struct {
	AppID      string
	SessionKey string
}

type showError struct {
	errorCode int
	errorMsg  error
}

func (e showError) Error() string {
	return fmt.Sprintf("{code: %v, error: \"%v\"}", e.errorCode, e.errorMsg)
}

// Decrypt Weixin APP's AES Data
// If isJSON is true, Decrypt return JSON type.
// If isJSON is false, Decrypt return map type.
func (wxCrypt *WxBizDataCrypt) Decrypt(encryptedData string, iv string, isJSON bool) (interface{}, error) {
	if len(wxCrypt.SessionKey) != 24 {
		return nil, showError{errorCode["IllegalAesKey"], errors.New("sessionKey length is error")}
	}
	aesKey, err := base64.StdEncoding.DecodeString(wxCrypt.SessionKey)
	if err != nil {
		return nil, showError{errorCode["DecodeBase64Error"], err}
	}

	if len(iv) != 24 {
		return nil, showError{errorCode["IllegalIv"], errors.New("iv length is error")}
	}
	aesIV, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, showError{errorCode["DecodeBase64Error"], err}
	}

	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, showError{errorCode["DecodeBase64Error"], err}
	}
	aesPlantText := make([]byte, len(aesCipherText))

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, showError{errorCode["IllegalBuffer"], err}
	}

	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
	mode.CryptBlocks(aesPlantText, aesCipherText)
	aesPlantText = PKCS7UnPadding(aesPlantText)

	var decrypted map[string]interface{}
	aesPlantText = []byte(strings.Replace(string(aesPlantText), "\a", "", -1))
	err = json.Unmarshal([]byte(aesPlantText), &decrypted)
	if err != nil {
		return nil, showError{errorCode["IllegalBuffer"], err}
	}

	if decrypted["watermark"].(map[string]interface{})["appid"] != wxCrypt.AppID {
		return nil, showError{errorCode["IllegalBuffer"], errors.New("appId is not match")}
	}

	if isJSON == true {
		return string(aesPlantText), nil
	}

	return decrypted, nil
}

// PKCS7UnPadding return unpadding []Byte plantText
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unPadding := int(plantText[length-1])
	if unPadding < 1 || unPadding > 32 {
		unPadding = 0
	}
	return plantText[:(length - unPadding)]
}

func TestDecode(t *testing.T) {
	wx := WxBizDataCrypt{AppID: "wxe8633ed42df5ea99", SessionKey: "rWIG3jSAcFLGCXCHPmpkEA=="}
	str, err := wx.Decrypt(data, iv, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
}

var jsonstr = ` {
"stepInfoList": [
{
"timestamp": 1531584000,
"step": 8527
},
{
"timestamp": 1531670400,
"step": 9045
},
{
"timestamp": 1531756800,
"step": 9569
},
{
"timestamp": 1531843200,
"step": 7950
},
{
"timestamp": 1531929600,
"step": 7692
},
{
"timestamp": 1532016000,
"step": 8489
},
{
"timestamp": 1532102400,
"step": 7057
},
{
"timestamp": 1532188800,
"step": 4484
},
{
"timestamp": 1532275200,
"step": 8578
},
{
"timestamp": 1532361600,
"step": 8494
},
{
"timestamp": 1532448000,
"step": 7105
},
{
"timestamp": 1532534400,
"step": 9137
},
{
"timestamp": 1532620800,
"step": 8446
},
{
"timestamp": 1532707200,
"step": 8799
},
{
"timestamp": 1532793600,
"step": 12847
},
{
"timestamp": 1532880000,
"step": 9231
},
{
"timestamp": 1532966400,
"step": 8917
},
{
"timestamp": 1533052800,
"step": 10561
},
{
"timestamp": 1533139200,
"step": 10342
},
{
"timestamp": 1533225600,
"step": 8312
},
{
"timestamp": 1533312000,
"step": 12282
},
{
"timestamp": 1533398400,
"step": 1012
},
{
"timestamp": 1533484800,
"step": 8688
},
{
"timestamp": 1533571200,
"step": 10075
},
{
"timestamp": 1533657600,
"step": 9331
},
{
"timestamp": 1533744000,
"step": 10095
},
{
"timestamp": 1533830400,
"step": 11738
},
{
"timestamp": 1533916800,
"step": 14876
},
{
"timestamp": 1534003200,
"step": 1190
},
{
"timestamp": 1534089600,
"step": 8507
},
{
"timestamp": 1534176000,
"step": 3946
}
],
"watermark": {
"timestamp": 1534243383,
"appid": "wxe8633ed42df5ea99"
}
}`

type StepData struct {
	StepInfoList []DayStep `json:"stepInfoList"`
}

type DayStep struct {
	Timestamp int64 `json:"timestamp"`
	Step      int   `json:"step"`
}

func TestJson(t *testing.T) {
	var stepData StepData
	err := json.Unmarshal([]byte(jsonstr), &stepData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stepData)

	for _,v := range stepData.StepInfoList {
		fmt.Println(lib.FormatTime(time.Unix(v.Timestamp,0)),v.Step)
	}
}
