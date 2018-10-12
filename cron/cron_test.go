package cron

import (
	"testing"
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func TestCron(t *testing.T){
	c := cron.New()
	c.AddFunc("1 * * * * *", func() { fmt.Println("Every hour on the half hour",time.Now()) })
	c.Start()
	select {

	}
}


func TestCro(t *testing.T){
}



