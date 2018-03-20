package file

import (
	"testing"
	"os"
	"fmt"
	"bufio"
	"time"
	"strconv"
)

func TestWrite(t *testing.T) {
	for i := 0; i < 10; i++ {
		write(strconv.Itoa(i))
	}
}

func write(fileName string) {
	outputFile, outputError := os.OpenFile("data", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!"

	for i := 0; i <= 10; i++ {
		outputWriter.WriteString(outputString + strconv.Itoa(i) + "\n")
		time.Sleep(1 * time.Second)
		outputWriter.Flush()
	}
	outputFile.Close()
	err := os.Rename("data", fileName)
	if err != nil {
		fmt.Println(err)
	}
}

func TestRead(t *testing.T) {
	file, err := os.Open("E:\\java_projects\\java_demos\\log.log")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			fmt.Print(str)
		} else {
			fmt.Println("read:" + err.Error())
			time.Sleep(300 * time.Millisecond)
			fileNext, inputError := os.Open("E:\\java_projects\\java_demos\\log.log")
			if inputError != nil {
				fmt.Println("open:" + inputError.Error())
				return // exit the function on error
			}
			if file.Fd() != fileNext.Fd() {
				file = fileNext
			}
			fileNext.Close()
		}
	}
}
