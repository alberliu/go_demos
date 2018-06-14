package file_test

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"time"
	"strconv"
	"io"
)

func TestWrite(t *testing.T) {
		write("data")
}

func write(fileName string) {
	outputFile, err := os.OpenFile("data", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!"

	for i := 0; i <= 30; i++ {
		outputWriter.WriteString(outputString + strconv.Itoa(i) + "\n")
		time.Sleep(1 * time.Second)
		outputWriter.Flush()
	}
	outputFile.Close()
}

func TestRead(t *testing.T) {
	file, err := os.Open("data")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			time.Sleep(50*time.Millisecond)
			continue
		}
		fmt.Print(str)
	}
}
