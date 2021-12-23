package factory

import (
	"SendEmailBatchKafka/domain/entity/email"
	"SendEmailBatchKafka/domain/repository"
)

type EmailSenderFactory interface {
	Send() repository.TransactionalRepository
}

type EmailTransactionalFactory interface {
	Save(sender email.EmailSender)
}
