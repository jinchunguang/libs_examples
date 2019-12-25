package main

import (
	"github.com/nsqio/go-nsq"
	"log"
)

var err error

func main() {

	config := nsq.NewConfig()

	producer, _ := nsq.NewProducer("127.0.0.1:4150", config)
	err = producer.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer producer.Stop()

	topicName := "test-topic"
	for i := 1; i < 5; i++ {
		err = producer.Publish(topicName, []byte("hello world!"))
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("ok")
}
