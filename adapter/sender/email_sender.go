package sender

import (
	"SendEmailBatchKafka/domain/entity/email"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
)

type Email struct {
	sender  *email.EmailSender
	connect *email.EmailConnect
}

func NewEmail(sender *email.EmailSender, connect *email.EmailConnect) *Email {
	return &Email{
		sender:  sender,
		connect: connect,
	}
}

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Parabens você ganhou</title>
</head>
<body>
   <p>Um pau no cu</p>
</body>
`
var server = mail.NewSMTPClient()
var sender = mail.NewMSG()
func (e *Email) Send() *email.Email {

	//prepare send 
	smtpClient, err := configureServer(e)
	succes := email.SEND
	erro := ""
	if err != nil {
		log.Println("Error ", err)
		erro = err.Error()
		succes = email.ERROR
	}
	//create email
	err = sendEmail(e,smtpClient)
	if err != nil {
		log.Println("Error ", err)
		erro = err.Error()
		succes = email.ERROR
	}

	email := email.NewEmailEntity(*e.sender, succes, erro)

	return email

}

func sendEmail(e *Email, smtpClient *mail.SMTPClient) error {
	sender.SetFrom(e.sender.From)
	for _, to := range e.sender.To {
		sender.AddTo(to)
	}
	for _, cc := range e.sender.Copy {
		sender.AddCc(cc)
	}
	sender.SetSubject("Email do Jonas")

	sender.SetBody(mail.TextHTML, htmlBody)
	return sender.Send(smtpClient)
}

func configureServer(e *Email) (*mail.SMTPClient, error) {
	server.Host = e.connect.ServerSmt
	server.Port = e.connect.Port
	server.Username = e.connect.Login
	server.Password = e.connect.Password
	server.Encryption = mail.EncryptionTLS
	smtpClient, err := server.Connect()
	return smtpClient, err
}
