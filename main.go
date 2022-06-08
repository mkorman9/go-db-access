package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/mkorman9/go-commons/logging"
	"github.com/mkorman9/go-commons/postgres"
	"github.com/rs/zerolog/log"
	"time"
)

type ClientEntity struct {
	ID                string    `gorm:"column:id; type:uuid; primaryKey"`
	Gender            string    `gorm:"column:gender; type:char(1)"`
	FirstName         string    `gorm:"column:first_name; type:varchar(255)"`
	LastName          string    `gorm:"column:last_name; type:varchar(255)"`
	Address           string    `gorm:"column:home_address; type:varchar(1024)"`
	PhoneNumber       string    `gorm:"column:phone_number; type:varchar(64)"`
	Email             string    `gorm:"column:email; type:varchar(64)"`
	BirthDate         time.Time `gorm:"column:birth_date; type:timestamp"`
	CreditCardsString string    `gorm:"column:credit_cards; type:varchar(255)"`
	IsDeleted         bool      `gorm:"column:deleted; type:boolean"`
}

func (ClientEntity) TableName() string {
	return "clients"
}

func main() {
	config.AddDriver(yaml.Driver)
	_ = config.LoadFiles("config.yml")

	logging.Setup()

	db, closeDB, err := postgres.DialPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot establish connection to Postgres, exiting!")
	}
	defer closeDB()

	err = db.AutoMigrate(&ClientEntity{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to auto-migrate schema")
	}

	log.Info().Msg("Done :)")
}
