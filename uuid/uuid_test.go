package uuid

import (
	"testing"
	"github.com/satori/go.uuid"
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		u2, _ := uuid.NewV4()
		str:=strings.Replace(u2.String(),"-","",-1)
		fmt.Println(str)
	}
}

func TestUUID2(t *testing.T){
	// Creating UUID Version 4
	// panic on error
	u1 := uuid.Must(uuid.NewV4())
	fmt.Printf("UUIDv4: %s\n", u1)

	// or error handling
	u2, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("UUIDv4: %s\n", u2)

	// Parsing UUID from string input
	u2, err = uuid.FromString("6ba7b8109dad11d180b400c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
	}
	fmt.Printf("Successfully parsed: %s", u2)
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
