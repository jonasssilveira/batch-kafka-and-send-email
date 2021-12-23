package migrations

import (
	"SendEmailBatchKafka/domain/entity/email"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB){
	db.AutoMigrate(&email.Email{})
}
