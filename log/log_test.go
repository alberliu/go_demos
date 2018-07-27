package main

import (
	"testing"
	"log"
	"fmt"
)

func TestLog(t *testing.T){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("hello")

	fmt.Println(len("6c814682-4a63-4215-8352-b34600543536"))
}
