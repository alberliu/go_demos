package log

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	logrus.Info("hello, world.")
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}

func BenchmarkLorus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// logrus.Info("hello, world.")
		fmt.Println("hello world")
	}
}
