package sego

import (
	"testing"
	"github.com/huichen/sego"
	"fmt"
)

func TestSego(t *testing.T){
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("/Users/alberliu/Workspace/gopath/src/github.com/huichen/sego/data/dictionary.txt")

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)

	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToString(segments, true))
}
