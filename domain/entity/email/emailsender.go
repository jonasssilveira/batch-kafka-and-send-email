package email

type EmailSender struct {
	From       string   `json:"from";gorm:"type:text"`
	Copy       []string `json:"copy";gorm:"type:text"`
	OccultCopy []string `json:"occult-copy";gorm:"type:text"`
	To         []string `json:"to";gorm:"type:text"`
	Assunto    string   `json:"assunto";gorm:"type:text"`
	Body       string   `json:"body";gorm:"type:text"`
}
