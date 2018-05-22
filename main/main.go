package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("hello","world")
	f()
}

func f(){
	log.Println("hello","world")
}

