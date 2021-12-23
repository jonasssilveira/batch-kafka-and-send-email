package usecase

import (
	"SendEmailBatchKafka/domain/entity/email"
	"SendEmailBatchKafka/domain/repository"
	"SendEmailBatchKafka/domain/sendemail"
	"log"
)

type Email struct {
	Repository  repository.TransactionalRepository
	EmailSender sendemail.SendEmail
	Email       *email.Email
}

func NewEmail(repository repository.TransactionalRepository, sendemail sendemail.SendEmail) *Email {
	return &Email{
		Repository:  repository,
		EmailSender: sendemail,
	}
}

func (emailS *Email) Send() {
	send := emailS.EmailSender.Send()

	err := emailS.Repository.Insert(send)
	emailS.Email = send
	if err != nil {
		log.Fatal("Error to save", err)
		return
	}
}
