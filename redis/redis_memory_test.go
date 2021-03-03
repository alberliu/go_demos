package redis

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

var keysSizeMap = map[string]int64{}

func TestMemoryRedis(t *testing.T) {
	urls := []string{
		"/Users/admin/Desktop/yuyin_redis/memory1.csv",
		"/Users/admin/Desktop/yuyin_redis/memory2.csv",
		"/Users/admin/Desktop/yuyin_redis/memory3.csv",
		"/Users/admin/Desktop/yuyin_redis/memory4.csv",
		"/Users/admin/Desktop/yuyin_redis/memory5.csv",
		"/Users/admin/Desktop/yuyin_redis/memory6.csv",
		"/Users/admin/Desktop/yuyin_redis/memory7.csv",
		"/Users/admin/Desktop/yuyin_redis/memory8.csv",
	}

	for i := range urls {
		size(urls[i])
	}

	printTop()
}

func size(url string) {
	file, err := os.Open(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	r := csv.NewReader(file)
	var i = 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			fmt.Println("eof")
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}

		i++
		if i == 1 {
			continue
		}

		key := strings.Split(record[2], ":")[0]
		size, _ := strconv.ParseInt(record[3], 10, 64)
		keysSizeMap[key] += size
	}
}

type keySize struct {
	key  string
	size int64
}

func printTop() {
	var keys = make([]struct {
		key  string
		size int64
	}, 0, len(keysSizeMap))

	for k, v := range keysSizeMap {
		keys = append(keys, keySize{
			key:  k,
			size: v,
		})
	}

	sort.Slice(keys, func(i, j int) bool {
		if keys[i].size >= keys[j].size {
			return true
		}
		return false
	})

	for i := -0; i < 100; i++ {
		fmt.Printf("%3d %-50s %20dM\n", i, keys[i].key, keys[i].size/1024/1024)
	}
}
