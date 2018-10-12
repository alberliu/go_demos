package gin

import (
	"testing"
	"github.com/gin-gonic/gin"
	"fmt"
)

func F1(c *gin.Context){
	fmt.Println("hello")
	defer recoverPanic()
	c.Next()
	fmt.Println("world")

}

func recoverPanic(){
	p:=recover()
	fmt.Println(p)
}

func TestGin(t *testing.T){
	r := gin.New()

	r.Use(F1)
	r.GET("/hello1", func(c *gin.Context) {
		fmt.Println("do1")
		panic("panic:hello1")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/hello2", func(c *gin.Context) {
		fmt.Println("do2")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func TestGinJsonIn(t *testing.T){

}

