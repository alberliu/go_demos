package tsw

import (
	"testing"
	"time"
	"fmt"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("hello")
	}

}

type timeSlide struct {
	time     time.Duration
	minTime  time.Duration
	allCount int
	counts   []int
}

func NewTimeSlide(time time.Duration, minTime time.Duration) timeSlide {
	numCount := time / minTime
	counts := make([]int, numCount)
	return timeSlide{time, minTime, 0, counts}
}

