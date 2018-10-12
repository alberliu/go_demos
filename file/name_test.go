package file

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"strings"
	"math/rand"
	"time"
	"step/domain/lib"
	"encoding/json"
)

var names = `蔡军,秦磊,马丽,赖刚,卢勇,胡秀英,程秀兰,杜洋,阎静,唐丽,姜娟,石娟,郭丽,阎勇,魏勇,董强,王刚,金勇,邵霞,姜丽,杜秀兰,方涛,崔桂英,段超,易丽,陆超,傅芳,胡敏,史超,石秀兰,龚杰,曹娜,董秀兰,赵静,卢杰,韩刚,廖秀兰,魏军,萧艳,彭静,吴桂英,潘娜,易秀英,贾超,吴强,吕敏,宋秀英,蔡秀兰,贺娜,石明,范洋,梁秀兰,许伟,魏伟,戴刚,侯洋,金明,周敏,宋芳,程磊,万超,苏静,吕丽,阎敏,袁艳,康磊,张刚,谭伟,陆芳,卢明,黎敏,吕秀英,段勇,程杰,周芳,彭勇,段刚,蒋刚,丁勇,郭涛,范敏,罗勇,夏杰,蒋秀兰,薛霞,胡军,许磊,阎丽,康磊,韩勇,唐娜,姚平,胡军,曾洋,韩超,邵磊,罗平,石明,武强,韩娜,毛涛,江超,熊洋,阎刚,萧艳,薛娟,姜强,陆洋,谭艳,梁强,江芳,毛军,任娜,傅平,谭娜,崔军,罗伟,汤芳,韩娜,廖平,于涛,苏艳,万娟,杨娟,郭涛,史霞,冯涛,顾娟,汪静,宋娜,马磊,崔静,康刚,文静,曹涛,梁明,胡娟,宋秀英,石丽,曹娜,毛娟,高磊,何洋,沈桂英,叶杰,徐平,吴涛,田秀英,史丽,侯涛,赖明,万洋,周静,贾霞,常娜,郑娟,傅杰,康军,董艳,宋丽,姜秀兰,戴艳,韩静,阎磊,万强,李娟,叶娟,罗磊,曹秀英,姚娜,胡娟,孔敏,吕明,余超,唐杰,卢军,顾娟,夏敏,戴秀兰,邵杰,黄丽,秦超,谭军,张静,乔平,杨娟,王艳,方丽,丁桂英,萧超,徐刚,苏艳,袁静,梁秀英,康娜,郝杰,蔡敏,徐超,张静,杜平,袁刚,何艳,胡杰,丁静,秦霞,黎娜,锺秀英,张丽,蔡超,郭敏,邹军,邓丽,邱刚,邵勇,白霞,孙娜,锺磊,徐磊,叶磊,蔡杰,姚平,何军,邵桂英,蒋芳,梁丽,薛娜,赵娜,朱桂英,赵平,金娟,康勇,武娟,邱敏,张娟,史秀英,沈杰,于勇,周娜,锺强,何娟,何霞,梁芳,卢娜,段娜,常军,姚艳,戴军,赖艳,韩伟,袁娜,贺涛,张静,金娜,戴娜,孙涛,贺超,魏芳,杨杰,何勇,潘敏,尹刚,黄桂英,文霞,沈明,段芳,龚娜,韩刚,阎芳,顾娜,秦洋,陈杰,龚霞,戴超,顾超,徐芳,黄杰,罗伟,姜霞,卢勇,薛秀英,曾艳,吴杰,卢娜,冯秀英,邵磊,方敏,梁敏,贾娜,尹明,秦艳,叶涛,李丽,`

// GoodsOrder 商品订单
type GoodsOrder struct {
	Id            int64  `json:"id"`              // 自增主键
	UserNickName  string `json:"user_nick_name"`  // 用户昵称
	UserAvatarUrl string `json:"user_avatar_url"` // 用户头像
	CreateTime    string `json:"create_time"`     // 创建时间
}

type GoodsOrderData struct {
	Id          int      `json:"id"`
	GoodsOrders []GoodsOrder `json:"goods_orders"`
}

func TestName(t *testing.T) {
	file, err := os.Open("/Users/alberliu/Desktop/详情页兑换记录用户头像url.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	var avatarUrls []string

	for {
		avatarUrl, err := reader.ReadString('')
		if err != nil {
			break
		}
		avatarUrls = append(avatarUrls, avatarUrl)
	}

	nickNames := strings.Split(names, ",")

	index := 0

	var goodsOrderDatas []GoodsOrderData

	numExchanged := []int{35, 80, 26, 0, 25, 13, 10, 8, 5, 2}

	for i, v := range numExchanged {
		var orders []GoodsOrder
		for i := 0; i < v; i++ {
			s := rand.Intn(345600) - 172800
			addTime := time.Now().Add((time.Duration(s) * time.Second))

			order := GoodsOrder{
				Id:            rand.Int63n(1000),
				UserNickName:  nickNames[index],
				UserAvatarUrl: avatarUrls[index],
				CreateTime:    lib.FormatTime(addTime),
			}
			orders = append(orders, order)
			index++
		}
		goodsOrderDatas=append(goodsOrderDatas,GoodsOrderData{
			Id:i+1,
			GoodsOrders:orders,
		})
	}

	bytes, err := json.Marshal(goodsOrderDatas)
	fmt.Println(string(bytes))
}
