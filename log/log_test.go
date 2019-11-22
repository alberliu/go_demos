package log

import (
	"bufio"
	"fmt"
	"github.com/json-iterator/go"
	"io"
	"os"
	"testing"
)

var Value = []int64{10, 50, 100, 200, 500, 1000, 5000, 10000, 20000, 30000, 500000}

type Statistics struct {
	m map[int64]int64
}

func NewStatistics() *Statistics {
	m := map[int64]int64{}
	for i := range Value {
		m[Value[i]] = 0
	}
	return &Statistics{m: m}
}

func (s *Statistics) Add(d int64) {
	for i := range Value {
		if d < Value[i] {
			s.m[Value[i]] = s.m[Value[i]] + 1
			return
		}
	}
}

func (s *Statistics) Print() {
	var sum int64 = 0
	for _, v := range s.m {
		sum += v
	}
	fmt.Println("sum:", sum)
	for i := range Value {
		if i == 0 {
			fmt.Printf("%10d-->%10d  %10d  %10.1f\n", 0, Value[i], s.m[Value[i]], float64(s.m[Value[i]]*100)/float64(sum))
		} else {
			fmt.Printf("%10d-->%10d  %10d  %10.1f\n", Value[i-1], Value[i], s.m[Value[i]], float64(s.m[Value[i]]*100)/float64(sum))
		}
	}
}

func TestLog(t *testing.T) {
	file, err := os.Open("/Users/abc/Workspace/yun_game/server/src/yunGame/arcadeconnect/test/bench_client/test.log")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer file.Close()

	statistics := NewStatistics()

	reader := bufio.NewReader(file)
	for {
		bytes, _, err := reader.ReadLine()
		if err == io.EOF {
			statistics.Print()
			return
		}

		duration := jsoniter.Get(bytes, "duration").ToInt64()
		statistics.Add(duration)

	}
}
