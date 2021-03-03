package file_test

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestWrite(t *testing.T) {
	write("data")
}

func write(fileName string) {
	file, err := os.OpenFile("data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputWriter := bufio.NewWriter(file)
	outputString := "hello world!"

	for i := 0; i <= 30; i++ {
		outputWriter.WriteString(outputString + strconv.Itoa(i) + "")
		time.Sleep(1 * time.Second)
		outputWriter.Flush()
	}
	file.Close()
}

type Log struct {
	Obj  string `json:"obj"`
	Func string `json:"func"`
}

func TestRead(t *testing.T) {
	var result = map[string]string{}

	file, err := os.Open("/Users/admin/Desktop/log1.textClipping")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		var log Log
		err = json.Unmarshal(str, &log)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result[log.Obj+"."+log.Func] = ""
	}

	file, err = os.Open("/Users/admin/Desktop/log2.textClipping")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer file.Close()
	reader = bufio.NewReader(file)
	for {
		str, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		var log Log
		err = json.Unmarshal(str, &log)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result[log.Obj+"."+log.Func] = ""
	}

	for k, _ := range result {
		fmt.Println(k)
	}
}
