package sendemail

import (
	"SendEmailBatchKafka/domain/entity/email"
)

type SendEmail interface {
	Send()* email.Email
}
