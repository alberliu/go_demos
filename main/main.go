package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	ticm "github.com/TencentCloud/tencentcloud-sdk-go/tencentcloud/ticm/v20181127"
)

func main() {
	credential := common.NewCredential(
		"AKIDQejZIloHH3ZrvmG6zc1A5lBPNHxQicMp",
		"QbdQ0LTLyPbdS3PeeU8auysZ5SdkcFck",
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	client, err := ticm.NewClient(credential, regions.Guangzhou, cpf)
	if err != nil {
		fmt.Println(err)
		return
	}

	request := ticm.NewImageModerationRequest()
	url := "https://dpic1.tiankong.com/8p/xz/QJ6382090849.jpg?x-oss-process=style/240h"
	request.ImageUrl = &url
	scenes := "PORN"
	request.Scenes = []*string{&scenes}

	response, err := client.ImageModeration(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	if response.Response.PornResult.Suggestion == "" {

	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
