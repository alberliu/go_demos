package datetime

import (
	"testing"
	"fmt"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	time := time.Time{}

	if now.After(time){
		fmt.Println("yes")
	}
}
