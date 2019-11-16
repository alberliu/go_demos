package nsq

import (
	"testing"

	"fmt"

	"time"

	"github.com/nsqio/go-nsq"
)

func TestProduct(t *testing.T) {
	producer, err := nsq.NewProducer("127.0.0.1:4150", nsq.NewConfig()) // 新建生产者
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := producer.Publish("test", []byte("hello")); err != nil { // 发布消息
		fmt.Println(err)
	}
}

func TestConsumer(t *testing.T) {
	NsqConsumer("test.1:2", "1", handle, 1)
	select {}
}

func handle(msg *nsq.Message) error {
	fmt.Println(string(msg.Body))
	return nil
}

// NsqConsumer 消费消息
func NsqConsumer(topic, channel string, handle func(message *nsq.Message) error, concurrency int) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 1 * time.Second

	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(nsq.HandlerFunc(handle), concurrency)
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic(err)
	}
	consumer.Stop()
}
