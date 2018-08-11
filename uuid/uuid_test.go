package uuid

import (
	"testing"
	"github.com/satori/go.uuid"
	"fmt"
	"math/rand"
	"time"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		u2, _ := uuid.NewV4()
		fmt.Println(len(u2.String()))
	}
}

func GenerateRandomBytes(length int) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func TestRand(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(RandString(10))
	}
}
