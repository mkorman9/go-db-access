package main

import (
	"encoding/json"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/mkorman9/go-commons/logging"
	"github.com/mkorman9/go-commons/postgres"
	"github.com/mkorman9/go-commons/utils"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

func main() {
	config.AddDriver(yaml.Driver)
	_ = config.LoadFiles("config.yml")

	logging.Setup()

	db, closeDB, err := postgres.DialPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot establish connection to Postgres, exiting!")
	}
	defer closeDB()

	err = db.AutoMigrate(
		&Account{},
		&Credentials{},
		&Session{},
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to auto-migrate schema")
	}

	accountID := uuid.NewV4().String()
	account := createAccount(accountID, db)

	sessionID := uuid.NewV4().String()
	createSession(sessionID, account, db)
}

func createAccount(accountID string, db *gorm.DB) *Account {
	account := Account{
		ID:          accountID,
		Username:    RandStringRunes(10),
		Roles:       []string{"PERMISSIONS_ADMIN"},
		IsDeleted:   false,
		BannedUntil: nil,
		Credentials: &Credentials{
			Email:          fmt.Sprintf("%s@example.com", RandStringRunes(8)),
			PasswordBcrypt: RandStringRunes(20),
		},
	}
	if result := db.Create(&account); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to insert account")
	}

	var accountFound Account
	if result := db.
		Preload("Credentials").
		First(&accountFound, "id = ?", accountID); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to select record")
	}

	b, _ := json.Marshal(accountFound)
	log.Info().Msg(string(b))

	return &account
}

func createSession(sessionID string, account *Account, db *gorm.DB) {
	session := Session{
		ID:        sessionID,
		AccountID: account.ID,
		Token:     RandStringRunes(48),
		Roles:     account.Roles,
		ExpiresAt: utils.ToPtr(time.Now().UTC().Add(4 * time.Hour)),
		Account:   account,
	}
	if result := db.Create(&session); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to insert session")
	}

	var sessionFound Session
	if result := db.
		Preload("Account").
		Preload("Account.Credentials").
		First(&sessionFound, "id = ?", sessionID); result.Error != nil {
		log.Fatal().Err(result.Error).Msg("Failed to select record")
	}

	b, _ := json.Marshal(sessionFound)
	log.Info().Msg(string(b))
}
