package gin

import (
	"fmt"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGin(t *testing.T) {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("do1")
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
