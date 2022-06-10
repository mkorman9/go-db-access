package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/mkorman9/go-commons/logging"
	"github.com/mkorman9/go-commons/postgres"
	"github.com/mkorman9/go-commons/server"
	"github.com/mkorman9/go-commons/utils"
	"github.com/mkorman9/go-commons/web"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
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

	err = db.AutoMigrate(&ClientEntity{})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to auto-migrate schema")
	}

	id := uuid.NewV4().String()
	client := Client{
		ID:          id,
		Gender:      GenderMale,
		FirstName:   "AAA",
		LastName:    "BBB",
		Address:     "AAA 123/456",
		PhoneNumber: "123-456-789",
		Email:       "aaa@example.com",
		BirthDate:   utils.TimePtr(time.Now().UTC()),
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

	b, _ := json.Marshal(entity.ToClient())
	log.Info().Msg(string(b))

	// server code
	s := server.NewServer()
	errorChannel := make(chan error)

	s.Engine.LoadHTMLGlob("templates/*.html")

	s.Engine.GET("/", func(c *gin.Context) {
		var entities ClientEntities
		if result := db.Find(&entities); result.Error != nil {
			web.InternalError(c, err, "Error while retrieving clients")
			return
		}

		c.HTML(200, "index.html", gin.H{
			"clients": entities.ToClients(),
		})
	})

	s.Start(errorChannel)
	server.BlockThread(errorChannel)
}
