package repository

import (
	"SendEmailBatchKafka/domain/entity/email"
)

type TransactionalRepository interface {
	Insert(sender *email.Email) error
}
