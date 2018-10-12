package httpclient

import (
	"testing"
	"fmt"
	"io"
	"os"
	"bytes"
	"mime/multipart"
	"net/http"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
)

func TestUploadFile(t *testing.T){
	//打开文件句柄操作
	file, err := os.Open("/Users/alberliu/Workspace/my.jpg")
	if err != nil {
		fmt.Println("error opening file")
		return
	}
	defer file.Close()

	//创建一个模拟的form中的一个选项,这个form项现在是空的
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作, 设置文件的上传参数叫uploadfile, 文件名是filename,
	//相当于现在还没选择文件, form项里选择文件的选项
	fileWriter, err := bodyWriter.CreateFormFile("file","filename" )
	if err != nil {
		fmt.Println("error writing to buffer")
		return
	}

	//iocopy 这里相当于选择了文件,将文件放到form中
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	//获取上传文件的类型,multipart/form-data; boundary=...
	contentType := bodyWriter.FormDataContentType()

	//这个很关键,必须这样写关闭,不能使用defer关闭,不然会导致错误
	bodyWriter.Close()


	//发送post请求到服务端
	// http://www.runoob.com/mongodb/mongodb-query.html
	resp, err := http.Post("http://localhost:8080/upload/photo", contentType, bodyBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	logs.Info(resp.Status)
	logs.Info(string(resp_body))
}

func TestGet(t *testing.T){
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
