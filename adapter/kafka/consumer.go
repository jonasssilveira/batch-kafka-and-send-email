package kafka

import (
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics    []string
}

func NewConsumer(config *ckafka.ConfigMap, Topics []string) *Consumer {
	return &Consumer{
		ConfigMap: config,
		Topics:    Topics,
	}
}

func (c *Consumer) Consume(msgChan chan *ckafka.Message){

	consumer, err:=ckafka.NewConsumer(c.ConfigMap)

	if err != nil {
		log.Fatal("Erro", err)

	}
	err = consumer.SubscribeTopics(c.Topics,nil)

	if err != nil{
		log.Fatal("Erro", err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		fmt.Println(msg)
		if err == nil {
			msgChan <- msg
		}
	}
}
