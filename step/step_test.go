package step

import (
	"testing"
	"fmt"
	"encoding/base64"
	"crypto/aes"
	"crypto/cipher"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"giftone/gift-audit/lib"
	"errors"
	"step-wx/entity"
)

var data = `{"encrypted_data":"bU3uInpBAPaj5t6WwgRRiIiVf1E0yeTlxG7AIM9dHtz45CVZVCMKh33ztsrKbxnN9FtF4UmF6BDSS5YEfnkL8YPXS7sfd8jxDstXdZ8oMt7DabBomtbEt6oQnUIpIL2cdY5BBZsHGhDS1OJJ9fMj2jfb21fWSgFxiX1iE0F3t+5WIG9FD5sSZR24dnKKcs7W3jPC6tCOTcetLbAVbDvMypWnJnhPW1Bq/bewLVwROLn3g1ie6/YsLoopSFeIU3g/8Dc6Th835UWh1qs/mn/CqSRjMuWin9ebVaqs9wJWkp8hFzgh8Tu2yGzOYkSBvmhCH8h0GCPmpHIZTM/Sr8X+pyriS2byLzzpcx+iIMMVocMcNPKrYRY9rjC3GK1a5dmT9jK+u6hmwWarDAkTl/dgOfCzkvQizvygMp7bTT4shCihRBA7zwXZblnRcbQnrgdkI9IZO3MVrJUSXkNwxsa7mxBlLh1Kw8GTwXoDD2YuNgZ7QNAYktp2RYJaxOfXPbCpYnfG80waMnh06Pd+GUgqeUvttRw8E33sacSEQagOeWJohmj1s9CmyPfJ4BZJ0HwwAMSq1xV/cfn7VYxUmLqjtTEu1NU+2o1njAZw9zoYwfwy0A+gZSL9MQQxmDm3YJsw/zgX//Kduw/v08al2iNfvBWeRtEszRvaQi/oVTb1TlG7TU43u28BMpoBn/zKiVIIE3PPUbpT9UsZuYbVnVB80qVV2bj63MgWL47nzqSVi6mfsyF0zpuasHhHsrR0R9748qR99A/WrJucTFLHqjHj7AIS92hRvty+CcuHgswkM0d3BQPmO1gFU7+/pUdM+VWRdenLgaTOhc3kib6CrUduBTxulML8kdPn/UH9yId6VY8e8aquC+JNIedbSqFO7u5ocwqpPOxBE0pHP9uLF5aKsia4T4akwkuv1xB1UaVA+BcF5JeKc6XlnEtgtRIjoatN8DP0NhVwUt/KKAnvH+YfUX+Z4pxXG1/1rG35hMftOX5D8fgdF9ZHVfHPzZ9YM1bsZdFhl+8TrXJSvJTsqVUmH0zlfbzhKmHQFUNcBxCrpbH73wttYYiKhJhHteXYYMekZjBx86l9M/E6tsLn4Pl7i7jjyxcLyK0xTliObcQeBN7XFWsemJVlKIhSz3M38hnzYd7c3ps4Ufi34RuK37VZgPVDu5LWb2q6GymNse+XzvUvN2MlFkunN/3taxijx8oOch9SoJm0tvX3ygaM3hJa9UUUDWMM+7eNMK7K44xz7S9flJC6LhrP5b/DLlt85dpaatGQODS9tOAlrHc9gtGlWaOhpE173QdI7amEYgc+k3/0nsQDn/gp9atw1phhtcNzse5pJBKHNZ0sBDeqNC5nWrY+JzlM4sdbWpUZlcQTa0TZBUvOPGNWmHNJSBtvfcs3qscxUWEyAKJ7kTiw6SotbZFA5SnSsy8NUrDkVwI800/cL1xKxOap9GR7H+/OaZRJSxbqZhiUS2nndR6eDK1KBn0fSydzpWsDUFgqBLmZApdY7/kxio3TgiSzvTHDwq/RHc8E5bMLUJ/uMAwOTXDZzLU4O/6MnsP0PxD4tLUDREMCyPRKmxT9pzkleFgLUbmEc6uRMFb+rP+8sYB/za9tEhV+B9bS48jnRdDwAFa29XVPzbqDn49eR5Tu6XZhkGzGOoAAI26sU4anSarfSeKzmw==","iv":"+SFWLkATo1Vgf0OnLC7R1Q=="}`
func TestStepNum(t *testing.T) {
	jsonStr := struct {
		EncryptedData string `json:"encrypted_data"`
		Iv            string `json:"iv"`
	}{}
	err:=json.Unmarshal([]byte(data),&jsonStr)
	if err!=nil{
		fmt.Println(err)
		return
	}
	c := NewWXBizDataCrypt("SEm1AZyNhep6iXjniYFFNw==" )
	buf, err := c.Decrypt(jsonStr.EncryptedData, jsonStr.Iv)
	fmt.Println(err)
	fmt.Println(string(buf))
}

func TestToken(t *testing.T){
	fmt.Println(lib.Md5("26GJGKFHKGFKJJHKHOUTHGKHYTERREWUXNBLKJOI"))
}

// 微信数据解密
var (
	ErrAppIDNotMatch       = errors.New("app id not match")
	ErrInvalidBlockSize    = errors.New("invalid block size")
	ErrInvalidPKCS7Data    = errors.New("invalid PKCS7 data")
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

type WXBizDataCrypt struct {
	appID, sessionKey string
}

const AppID = "wxe8633ed42df5ea99"

func NewWXBizDataCrypt(sessionKey string) *WXBizDataCrypt {
	return &WXBizDataCrypt{
		appID:      AppID,
		sessionKey: sessionKey,
	}
}

func (w *WXBizDataCrypt) Decrypt(encryptedData, iv string) ([]byte, error) {
	aesKey, err := base64.StdEncoding.DecodeString(w.sessionKey)
	if err != nil {
		return nil, err
	}
	cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, ivBytes)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, err = pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// pkcs7Unpad returns slice of the original data without padding
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

// DecryptStepNum 解密微信运动步数,获取微信当天的运动步数
func DecryptStepNum(sessionKey, encryptedData, iv string) (int, error) {
	bytes, err := NewWXBizDataCrypt(sessionKey).Decrypt(encryptedData, iv)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	var stepData entity.StepData
	err = json.Unmarshal(bytes, &stepData)
	if err != nil {
		logs.Error(err)
		return 0, err
	}
	lenth := len(stepData.StepInfoList)
	return stepData.StepInfoList[lenth-1].Step, nil

}
