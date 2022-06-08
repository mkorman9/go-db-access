package main

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/mkorman9/go-commons/logging"
	"github.com/mkorman9/go-commons/postgres"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

type Client struct {
	ID          string
	Gender      string
	FirstName   string
	LastName    string
	Address     string
	PhoneNumber string
	Email       string
	BirthDate   time.Time
	CreditCards []string
	IsDeleted   bool
}

func (c *Client) ToEntity() *ClientEntity {
	return &ClientEntity{
		ID:                c.ID,
		Gender:            c.Gender,
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		Address:           c.Address,
		PhoneNumber:       c.PhoneNumber,
		Email:             c.Email,
		BirthDate:         c.BirthDate,
		CreditCardsString: strings.Join(c.CreditCards, ";"),
		IsDeleted:         c.IsDeleted,
	}
}

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

func (e *ClientEntity) ToClient() *Client {
	return &Client{
		ID:          e.ID,
		Gender:      e.Gender,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Address:     e.Address,
		PhoneNumber: e.PhoneNumber,
		Email:       e.Email,
		BirthDate:   e.BirthDate,
		CreditCards: strings.Split(e.CreditCardsString, ";"),
		IsDeleted:   e.IsDeleted,
	}
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

	id := uuid.NewV4().String()
	client := Client{
		ID:          id,
		Gender:      "M",
		FirstName:   "AAA",
		LastName:    "BBB",
		Address:     "AAA 123/456",
		PhoneNumber: "123-456-789",
		Email:       "aaa@example.com",
		BirthDate:   time.Now().UTC(),
		CreditCards: []string{"1111 2222 3333 4444"},
		IsDeleted:   false,
	}
	if result := db.Create(client.ToEntity()); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to insert record")
	}

	var entity ClientEntity
	if result := db.First(&entity, "id = ?", id); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to select record")
	}

	log.Info().Msgf("%v", entity.ToClient())
}
