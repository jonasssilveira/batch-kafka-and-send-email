package email

import (
	"github.com/satori/go.uuid"
	"time"
)

const (
	SEND  = "SUCCESS"
	ERROR = "ERROR"
)

type Email struct {
	ID         uuid.UUID `gorm:"id";json:"id"` //TODO pegar o id nao precisa ser agora
	From       string    `gorm:"type:text";json:"from"`
	Copy       []string  `gorm:"type:text";json:"copy"`
	OccultCopy []string  `gorm:"type:text";json:"occultcopy"`
	To         []string  `gorm:"type:text";json:"to"`
	Assunto    string    `gorm:"type:text";json:"assunto"`
	Body       string    `gorm:"type:text";json:"body"`
	Error      string    `gorm:"type:text";json:"erro"`
	Status     string    `gorm:"type:text";json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
}

func NewEmailEntity(sender EmailSender, status, error string) *Email {
	return &Email{
		ID:         uuid.NewV4(),
		From:       sender.From,
		Copy:       sender.Copy,
		OccultCopy: sender.OccultCopy,
		To:         sender.To,
		Assunto:    sender.Assunto,
		Body:       sender.Body,
		UpdateAt:   time.Now(),
		CreatedAt:  time.Now(),
		Status:     status,
		Error:      error,
	}
}
