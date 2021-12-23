package repository

import (
	"SendEmailBatchKafka/adapter/migrations"
	"SendEmailBatchKafka/domain/entity/email"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

type TransactionalRepositoryDB struct {
	db *gorm.DB
}

func NewTransactionalRepositoryDB()*TransactionalRepositoryDB{
	db, err := gorm.Open(sqlite.Open("email.db"), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
		return &TransactionalRepositoryDB{}
	}
	config, _ := db.DB()
	config.SetMaxOpenConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigration(db)

	return &TransactionalRepositoryDB{
		db: db,
	}
}

func (repo *TransactionalRepositoryDB) Insert(sender *email.Email) error{
	create := repo.db.Create(sender)
	if create == nil{
		log.Fatal("Erro ao salvar")
		return create.Error
	}
	return nil
}