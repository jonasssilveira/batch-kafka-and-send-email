package main

import (
	"SendEmailBatchKafka/adapter/kafka"
	"SendEmailBatchKafka/adapter/repository"
	"SendEmailBatchKafka/adapter/sender"
	"SendEmailBatchKafka/domain/entity/email"
	"encoding/json"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"

	"SendEmailBatchKafka/usecase"
)

func main() {

	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp",
		"group.id":          "email-batch-processor",
	}

	var msgChan = make(chan *ckafka.Message)

	topics := []string{"msg-email-batch"}
	consumer := kafka.NewConsumer(configMapProducer, topics)

	db := repository.NewTransactionalRepositoryDB()
	go consumer.Consume(msgChan)

	for msg := range msgChan {
		var emailSender email.EmailSender
		fmt.Println(msg.Value)
		json.Unmarshal(msg.Value, &emailSender)
		log.Println(emailSender)
		newEmail := sender.NewEmail(&emailSender, email.ConnectToEmail())
		email := usecase.NewEmail(db, newEmail)
		email.Send()

	}

}
